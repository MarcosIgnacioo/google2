package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/Marciglez/google2/utils"
)

var debug bool

const N = 10

type Function2 struct {
	FileName     string
	Row          int
	Weight       int
	FunctionName string // put all of this in the original struct
	Generics     *ast.FieldList
	Receiver     *ast.FieldList
	Params       *ast.FieldList
	Results      *ast.FieldList
}

func (f Function2) String() string {
	var builder strings.Builder
	if f.FileName != "" {
		builder.WriteString(f.FileName + " ")
	}
	if f.Receiver != nil {
		builder.WriteString("(" + utils.StringifyFieldList(f.Receiver) + ") ")
	}
	builder.WriteString(f.FunctionName)
	if f.Generics != nil {
		builder.WriteString("[" + utils.StringifyFieldList(f.Generics) + "]")

	}
	builder.WriteString("(" + utils.StringifyFieldList(f.Params) + ")")
	builder.WriteString(": (" + utils.StringifyFieldList(f.Results) + ") ")
	return builder.String()
}

func (f Function2) SignatureString() string {
	var builder strings.Builder
	if f.Receiver != nil {
		builder.WriteString("(" + utils.StringifyFieldListNoNames(f.Receiver) + ") ")
	}
	if f.Generics != nil {
		builder.WriteString("[" + utils.StringifyFieldListNoNames(f.Generics) + "]")
	}
	builder.WriteString("(" + utils.StringifyFieldListNoNames(f.Params) + ")")
	builder.WriteString(": (" + utils.StringifyFieldListNoNames(f.Results) + ") ")
	return builder.String()
}

// this is dirt i dont like it
func parse_ast_func_decl(fset *token.FileSet, func_decl *ast.FuncDecl) Function2 {
	var fun Function2
	if fset == nil {
		fun = Function2{
			FunctionName: func_decl.Name.Name,
			Generics:     (func_decl.Type.TypeParams),
			Receiver:     (func_decl.Recv),
			Params:       (func_decl.Type.Params),
			Results:      (func_decl.Type.Results),
		}
	} else {
		file := fset.Position(func_decl.Type.Func)
		fun = Function2{
			FileName:     file.String(),
			Row:          file.Line,
			FunctionName: func_decl.Name.Name,
			Generics:     (func_decl.Type.TypeParams),
			Receiver:     (func_decl.Recv),
			Params:       (func_decl.Type.Params),
			Results:      (func_decl.Type.Results),
		}
	}
	return fun
}

func stringify_array[T interface{ String() string }](array []T) string {
	var builder strings.Builder
	for _, item := range array {
		str := item.String()
		builder.WriteString(str)
		builder.WriteString("\n")
	}
	return builder.String()
}

func print_array[T interface{ String() string }](array []T) {
	fmt.Println(stringify_array(array))
}

func lev_distance_impl(a, b string, a_len, b_len int, cache [][]int) int {
	// we need to add all the characters from 'b' to turn 'a' into 'b'
	if cache[a_len][b_len] != -1 {
		return cache[a_len][b_len]
	}

	if a_len == 0 {
		cache[a_len][b_len] = b_len
		return b_len
	}

	// we need to remove all the characters from 'a' to turn 'b' into 'b'
	if b_len == 0 {
		cache[a_len][b_len] = a_len
		return cache[a_len][b_len]
	}

	// a_tail := a[:a_len-1]
	// b_tail := b[:b_len-1]

	if a[a_len-1] == b[b_len-1] { // ignore, the character matches
		match := lev_distance_impl(a, b, a_len-1, b_len-1, cache)
		cache[a_len][b_len] = match
		return cache[a_len][b_len]
	}

	// the + 1 is for the character action we are performing in each of these
	// situations., we ass one of this specificic ation to the others that the recursion
	// might have found
	deletion := lev_distance_impl(a, b, a_len-1, b_len, cache) + 1
	insertion := lev_distance_impl(a, b, a_len, b_len-1, cache) + 1
	replacement := lev_distance_impl(a, b, a_len-1, b_len-1, cache) + 1

	minimal_plus_one := min(
		deletion,
		insertion,
		replacement,
	)

	cache[a_len][b_len] = minimal_plus_one

	return cache[a_len][b_len]
}

// sadly i dont get what all of this mean asdfhasdhf
func lev_distance(a, b string) int {
	var cache [][]int = make([][]int, len(a)+1)
	for row := 0; row < len(a)+1; row += 1 {
		cache[row] = make([]int, len(b)+1)
		for col := 0; col < len(b)+1; col += 1 {
			cache[row][col] = -1
		}
	}
	return lev_distance_impl(a, b, len(a), len(b), cache)
}

func parse_raw_string(source string) (f *ast.File, err error) {
	var fset token.FileSet
	var flags parser.Mode
	var file_path string

	flags |= parser.SkipObjectResolution
	// yeah its scrapping the parser time
	// because we do this silly thing we dont actually the possibility to actually
	// pass generics or method declaration which is not great !!!
	// sooo we might aswell write (steal) our own function declaration go parser
	source = fmt.Sprintf("package main\n func x%s", source)
	return parser.ParseFile(&fset, file_path, source, flags)
}

func parse_user_function_query(query string) (*ast.FuncDecl, error) {
	query_ast, err := parse_raw_string(query)

	if err != nil {
		return nil, err
	}

	func_declarations := query_ast.Decls

	if len(func_declarations) < 1 {
		err = errors.New("not a single declaration in query")
		return nil, err
	}

	if func_declaration, ok := func_declarations[0].(*ast.FuncDecl); ok {
		return func_declaration, nil
	} else {
		err = errors.New("query input isnt a function declaration")
		return nil, err
	}

}

func get_lev_distance_from_field_list(target_field_list *ast.FieldList, left_field_list *ast.FieldList, right_field_list *ast.FieldList) (left_changes, right_changes int) {
	if target_field_list != nil && target_field_list.List != nil {
		str_list_target := utils.StringifyFieldListNoNames(target_field_list)
		str_list_left := utils.StringifyFieldListNoNames(left_field_list)
		str_list_right := utils.StringifyFieldListNoNames(right_field_list)
		left_changes += lev_distance(str_list_left, str_list_target)
		right_changes += lev_distance(str_list_right, str_list_target)
	}
	return
}

type List struct {
	left   *ast.FieldList
	right  *ast.FieldList
	target *ast.FieldList
}

func google2(user_query, file_name string) []Function2 {
	var fset token.FileSet
	var flags parser.Mode
	var user_function_search Function2
	var functions []Function2
	var user_function_normalized string

	query_ast, err := parse_user_function_query(user_query)

	if err != nil {
		panic(err)
	}
	// defer delete(query_ast)

	user_function_search = parse_ast_func_decl(nil, query_ast)
	// defer delete(user_function_search)
	user_function_normalized = user_function_search.SignatureString()
	// defer delete(user_function_normalized)

	flags |= parser.SkipObjectResolution
	file_ast, err := parser.ParseFile(&fset, file_name, nil, flags)
	if err != nil {
		panic(err)
	}
	// defer delete(file_ast)
	func_declarations := file_ast.Decls
	functions = make([]Function2, 0, 512)

	for _, declaration := range func_declarations {
		switch v := declaration.(type) {
		case *ast.FuncDecl:
			{
				function := parse_ast_func_decl(&fset, v)
				functions = append(functions, function)
			}
		}
	}

	log.Println("[FUZZY SEARCHING]:", user_function_normalized)

	sort.Slice(functions, func(i, j int) bool {
		left := functions[i]
		right := functions[j]
		// here i have access to the types and params and return things in the ast
		// data structure so i can kinda do some weird stuff for ordering stuff
		var left_changes_needed_for_match int
		var right_changes_needed_for_match int

		var field_lists []List

		field_lists = []List{
			{
				left.Generics,
				right.Generics,
				user_function_search.Generics,
			},
			{
				left.Receiver,
				right.Receiver,
				user_function_search.Receiver,
			},
			{
				left.Params,
				right.Params,
				user_function_search.Params,
			},
			{
				left.Results,
				right.Results,
				user_function_search.Results,
			},
		}

		for _, list := range field_lists {
			left_changes, right_changes := get_lev_distance_from_field_list(
				list.target,
				list.left,
				list.right,
			)
			left_changes_needed_for_match += left_changes
			right_changes_needed_for_match += right_changes
		}

		return left_changes_needed_for_match < right_changes_needed_for_match
	})

	return functions[:N]
}

func uwumain() {
	user_query := "(float32) Vector2"
	fndecl, _ := parse_user_function_query(user_query)
	fmt.Println(fndecl.Recv)
	fmt.Println(fndecl.Type.TypeParams)
	fmt.Println(fndecl.Type.Params)
	fmt.Println(fndecl.Type.Results)
}

// TODO: add a flag to make it case insensesitive

func main() {
	// parser2.Hello()
	debug = false
	fmt.Println("hello world")
	var functions []Function2

	if len(os.Args) < 3 {
		file_path := "/home/marcig/personal/fun/google2/examples/raymath.go"
		user_query := "(float32) Vector2"
		functions = google2(user_query, file_path)
	} else {
		functions = google2(os.Args[1], os.Args[2])
	}

	for _, fn := range functions {
		fn_str := fn.String()
		fmt.Println(fn_str)
	}
}
