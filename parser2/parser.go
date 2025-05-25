package parser2

import (
	"errors"
	"go/ast"
	"go/scanner"
	"go/token"
	"log"
)

// most complex input possible
// [T interface{ String() string }]([]T, *T) (string, error)
// most simple
// () int
// middle ground
// (*File) (string) int

type Parser2 struct {
	errors  scanner.ErrorList
	source  []byte
	pos     token.Pos   // token position
	tok     token.Token // one token look-ahead
	lit     string      // token literal
	scanner scanner.Scanner
}

func NewParser2(src string) *Parser2 {
	p := &Parser2{}
	p.init([]byte(src))
	return p
}

func (p *Parser2) init(src []byte) {
	eh := func(pos token.Position, msg string) { p.errors.Add(pos, msg) }
	src_length := len(src)
	file := token.NewFileSet().AddFile("cool", src_length, src_length)
	p.scanner.Init(file, src, eh, scanner.ScanComments)
	p.next()
}

func (p *Parser2) next0() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}

func (p *Parser2) next() {
	p.next0()
}

func (p *Parser2) expect(tok token.Token) token.Pos {
	pos := p.pos
	if p.tok != tok {
		log.Fatal("This is not valid golang function syntax")
	}
	p.next() // make progress
	return pos
}

func (p *Parser2) parse_field_list(closing token.Token) (*ast.FieldList, error) {
	field_list := &ast.FieldList{
		List: make([]*ast.Field, 0, 12),
	}
	for p.tok != closing {
		field := &ast.Field{Type: ast.NewIdent(p.lit)}
		field_list.List = append(field_list.List, field)
		p.next()
		if p.tok == token.COMMA {
			p.next()
		}
	}
	p.expect(closing)
	return nil, nil
}

// regular functions
// ()
// (): int
// (int, foo): int
// (int, foo)
// (int, ...): string
// (int, ...)
// (int): (int, error)
// (*Foo): (int, error)
// (*Foo) ()

// methods
// (*Foo) (): int
// (*Foo) (int): int
// (*Foo) () : (int, error)
// for today we will just implement simple type parsing, aka, wont parse interface{} things
// or functions types which are boring
// we leave generics also for another day
func (p *Parser2) ParseQuery() (function *ast.FuncDecl, err error) {
	function = &ast.FuncDecl{}
	func_type := &ast.FuncType{}
	var receiver, params, results *ast.FieldList

	if p.tok != token.LPAREN {
		err = errors.New("There is no left parenthesis unu")
		return
	}

	p.next()
	fields, err := p.parse_field_list(token.RPAREN)

	if err != nil {
		return
	}

	if p.tok == token.LPAREN {
		p.next()
		receiver = fields
		params, err = p.parse_field_list(token.RPAREN)
		if err != nil {
			return
		}
	} else {
		params = fields
	}

	if p.tok == token.COLON {
		p.next()
		results, err = p.parse_field_list(token.RPAREN)
		if err != nil {
			return
		}
	}

	function.Recv = receiver
	function.Type = func_type
	// memory allocation
	function.Type.Params = params
	function.Type.Results = results

	return
}
