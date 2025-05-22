package position

import (
	"cmp"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
)

func (f *File) Offset(p Pos) int {
	return f.fixOffset(int(p) - f.base)
}

// Line returns the line number for the given file position p;
// p must be a [Pos] value in that file or [NoPos].
func (f *File) Line(p Pos) int {
	return f.Position(p).Line
}

func searchLineInfos(a []lineInfo, x int) int {
	i, found := slices.BinarySearchFunc(a, x, func(a lineInfo, x int) int {
		return cmp.Compare(a.Offset, x)
	})
	if !found {
		// We want the lineInfo containing x, but if we didn't
		// find x then i is the next one.
		i--
	}
	return i
}

// unpack returns the filename and line and column number for a file offset.
// If adjusted is set, unpack will return the filename and line information
// possibly adjusted by //line comments; otherwise those comments are ignored.
func (f *File) unpack(offset int, adjusted bool) (filename string, line, column int) {
	f.mutex.Lock()
	filename = f.name
	if i := searchInts(f.lines, offset); i >= 0 {
		line, column = i+1, offset-f.lines[i]+1
	}
	if adjusted && len(f.infos) > 0 {
		// few files have extra line infos
		if i := searchLineInfos(f.infos, offset); i >= 0 {
			alt := &f.infos[i]
			filename = alt.Filename
			if i := searchInts(f.lines, alt.Offset); i >= 0 {
				// i+1 is the line at which the alternative position was recorded
				d := line - (i + 1) // line distance from alternative position base
				line = alt.Line + d
				if alt.Column == 0 {
					// alternative column is unknown => relative column is unknown
					// (the current specification for line directives requires
					// this to apply until the next PosBase/line directive,
					// not just until the new newline)
					column = 0
				} else if d == 0 {
					// the alternative position base is on the current line
					// => column is relative to alternative column
					column = alt.Column + (offset - alt.Offset)
				}
			}
		}
	}
	// TODO(mvdan): move Unlock back under Lock with a defer statement once
	// https://go.dev/issue/38471 is fixed to remove the performance penalty.
	f.mutex.Unlock()
	return
}

func (f *File) position(p Pos, adjusted bool) (pos Position) {
	offset := f.fixOffset(int(p) - f.base)
	pos.Offset = offset
	pos.Filename, pos.Line, pos.Column = f.unpack(offset, adjusted)
	return
}

// PositionFor returns the Position value for the given file position p.
// If p is out of bounds, it is adjusted to match the File.Offset behavior.
// If adjusted is set, the position may be adjusted by position-altering
// //line comments; otherwise those comments are ignored.
// p must be a Pos value in f or NoPos.
func (f *File) PositionFor(p Pos, adjusted bool) (pos Position) {
	if p != NoPos {
		pos = f.position(p, adjusted)
	}
	return
}

// Position returns the Position value for the given file position p.
// If p is out of bounds, it is adjusted to match the File.Offset behavior.
// Calling f.Position(p) is equivalent to calling f.PositionFor(p, true).
func (f *File) Position(p Pos) (pos Position) {
	return f.PositionFor(p, true)
}

// -----------------------------------------------------------------------------
// FileSet

// A FileSet represents a set of source files.
// Methods of file sets are synchronized; multiple goroutines
// may invoke them concurrently.
//
// The byte offsets for each file in a file set are mapped into
// distinct (integer) intervals, one interval [base, base+size]
// per file. [FileSet.Base] represents the first byte in the file, and size
// is the corresponding file size. A [Pos] value is a value in such
// an interval. By determining the interval a [Pos] value belongs
// to, the file, its file base, and thus the byte offset (position)
// the [Pos] value is representing can be computed.
//
// When adding a new file, a file base must be provided. That can
// be any integer value that is past the end of any interval of any
// file already in the file set. For convenience, [FileSet.Base] provides
// such a value, which is simply the end of the Pos interval of the most
// recently added file, plus one. Unless there is a need to extend an
// interval later, using the [FileSet.Base] should be used as argument
// for [FileSet.AddFile].
//
// A [File] may be removed from a FileSet when it is no longer needed.
// This may reduce memory usage in a long-running application.
type FileSet struct {
	mutex sync.RWMutex         // protects the file set
	base  int                  // base offset for the next file
	files []*File              // list of files in the order added to the set
	last  atomic.Pointer[File] // cache of last file looked up
}

// NewFileSet creates a new file set.
func NewFileSet() *FileSet {
	return &FileSet{
		base: 1, // 0 == NoPos
	}
}

// Base returns the minimum base offset that must be provided to
// [FileSet.AddFile] when adding the next file.
func (s *FileSet) Base() int {
	s.mutex.RLock()
	b := s.base
	s.mutex.RUnlock()
	return b
}

func (s *FileSet) AddFile(filename string, base, size int) *File {
	// Allocate f outside the critical section.
	f := &File{name: filename, size: size, lines: []int{0}}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if base < 0 {
		base = s.base
	}
	if base < s.base {
		panic(fmt.Sprintf("invalid base %d (should be >= %d)", base, s.base))
	}
	f.base = base
	if size < 0 {
		panic(fmt.Sprintf("invalid size %d (should be >= 0)", size))
	}
	// base >= s.base && size >= 0
	base += size + 1 // +1 because EOF also has a position
	if base < 0 {
		panic("token.Pos offset overflow (> 2G of source code in file set)")
	}
	// add the file to the file set
	s.base = base
	s.files = append(s.files, f)
	s.last.Store(f)
	return f
}
