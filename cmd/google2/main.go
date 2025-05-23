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
)

var debug bool

const N = 10

func stringify_expression(expression ast.Expr) string {
	switch v := expression.(type) {
	// case *ast.ChanType:
	// 	fallthrough
	// case *ast.InterfaceType:
	// 	fallthrough
	// case *ast.StructType:
	// 	fallthrough
	// case *ast.KeyValueExpr:
	// 	fallthrough
	// case *ast.CompositeLit:
	// 	fallthrough
	// case *ast.IndexListExpr:
	// 	fallthrough
	// case *ast.SliceExpr:
	// 	fallthrough
	// case *ast.TypeAssertExpr:
	// 	fallthrough
	// case *ast.CallExpr:
	// 	fallthrough
	// case *ast.UnaryExpr:
	// 	fallthrough

	case *ast.ParenExpr:
		{
			expr := stringify_expression(v.X)
			return fmt.Sprintf("%s", expr)
		}
	case *ast.BinaryExpr:
		{
			left := stringify_expression(v.X)
			right := stringify_expression(v.Y)
			return fmt.Sprintf("%s %s %s", left, v.Op.String(), right)
		}
	case *ast.FuncLit:
		{
			return stringify_expression(v.Type)
		}
	case *ast.InterfaceType:
		{
			return "interface{}"
		}
	case *ast.BadExpr:
		log.Fatalf("Expression not supported jijiji %v", v)
		return "[NOT SUPPORTED]"
	case *ast.IndexExpr:
		{
			exp := stringify_expression(v.X)
			index := stringify_expression(v.Index)
			return fmt.Sprintf("%s[%s]", exp, index)
		}
	case *ast.Ident:
		{
			return v.Name
		}
	case *ast.Ellipsis:
		{
			elipsis_type := stringify_expression(v.Elt)
			return fmt.Sprintf("...%s", elipsis_type)
		}
	case *ast.SelectorExpr:
		{
			selector_expr := stringify_expression(v.X)
			return fmt.Sprintf("%s.%s", selector_expr, v.Sel.Name)
		}
	case *ast.StarExpr:
		{
			star_expr := stringify_expression(v.X)
			return fmt.Sprintf("*%s", star_expr)
		}
	case *ast.ArrayType:
		{
			var count string
			if v.Len != nil {
				count = stringify_expression(v.Len)
			}
			array_type := stringify_expression(v.Elt)
			return fmt.Sprintf("%s[%s]", array_type, count)
		}
	case *ast.FuncType:
		{
			var params strings.Builder
			var results strings.Builder
			append_field_list(&params, v.Params)
			append_field_list(&results, v.Results)
			return fmt.Sprintf("func (%s) %s", params.String(), results.String())
		}
	case *ast.MapType:
		{
			key := stringify_expression(v.Key)
			value := stringify_expression(v.Value)
			return fmt.Sprintf("map[%s]%s", key, value)
		}
	case *ast.BasicLit:
		{
			kind := v.Kind.String()
			value := v.Value
			return fmt.Sprintf("%s %s", value, kind)
		}
	default:
		{
			log.Fatal("[ERROR]:", v)
			panic("nimodo")
		}
	}
}

func append_field_list(builder *strings.Builder, field_list_container *ast.FieldList) {
	if field_list_container == nil {
		return
	}
	field_list := field_list_container.List
	for i, param := range field_list {
		if i > 0 {
			builder.WriteString(",")
		}
		for k, param_name := range param.Names {
			if k > 0 {
				builder.WriteString(",")
			}
			_, err := builder.WriteString(fmt.Sprintf("%s ", param_name.Name))
			if err != nil {
				fmt.Println("an error has happened unu")
			}
		}
		param_type := stringify_expression(param.Type)
		_, err := builder.WriteString(fmt.Sprintf("%s", param_type))
		if err != nil {
			fmt.Println("[ERROR]:an error has happened unu")
		}
	}
}

// todo make just two functions one for stringify without names and one with them, because
// the if is making me itchy
func stringify_field_list(field_list_container *ast.FieldList) string {
	var builder strings.Builder

	if field_list_container == nil {
		return ""
	}
	field_list := field_list_container.List
	for i, param := range field_list {
		if i > 0 {
			builder.WriteString(", ")
		}
		for k, param_name := range param.Names {
			if k > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(fmt.Sprintf("%s", param_name.Name))
		}
		if len(param.Names) > 0 {
			builder.WriteString(" ")
		}
		param_type := stringify_expression(param.Type)
		builder.WriteString(fmt.Sprintf("%s", param_type))
	}

	return builder.String()
}

func stringify_field_list_without_names(field_list_container *ast.FieldList) string {
	var builder strings.Builder

	if field_list_container == nil {
		return ""
	}
	field_list := field_list_container.List
	for i, param := range field_list {
		if i > 0 {
			builder.WriteString(", ")
		}
		// builder.WriteString(fmt.Sprintf("%s", param_type))
		param_type := stringify_expression(param.Type)
		if param.Names != nil {
			for k := range param.Names {
				if k > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(fmt.Sprintf("%s", param_type))
			}
		} else {
			builder.WriteString(fmt.Sprintf("%s", param_type))
		}

	}

	return builder.String()
}

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
		builder.WriteString("(" + stringify_field_list(f.Receiver) + ") ")
	}
	builder.WriteString(f.FunctionName)
	if f.Generics != nil {
		builder.WriteString("[" + stringify_field_list(f.Generics) + "]")

	}
	builder.WriteString("(" + stringify_field_list(f.Params) + ")")
	builder.WriteString(": (" + stringify_field_list(f.Results) + ") ")
	return builder.String()
}

func (f Function2) SignatureString() string {
	var builder strings.Builder
	if f.Receiver != nil {
		builder.WriteString("(" + stringify_field_list_without_names(f.Receiver) + ") ")
	}
	if f.Generics != nil {
		builder.WriteString("[" + stringify_field_list_without_names(f.Generics) + "]")
	}
	builder.WriteString("(" + stringify_field_list_without_names(f.Params) + ")")
	builder.WriteString(": (" + stringify_field_list_without_names(f.Results) + ") ")
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
		str_list_target := stringify_field_list_without_names(target_field_list)
		str_list_left := stringify_field_list_without_names(left_field_list)
		str_list_right := stringify_field_list_without_names(right_field_list)
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

// TODO: add a flag to make it case insensesitive

func main() {
	// parser2.Hello()
	debug = false
	var functions []Function2
	if len(os.Args) < 3 {
		file_path := "/home/marcig/personal/fun/google2/examples/raymath.go"
		user_query := "(Vector3, float32) Vector2"
		functions = google2(user_query, file_path)
	} else {
		functions = google2(os.Args[1], os.Args[2])
	}
	for _, fn := range functions {
		fn_str := fn.String()
		if fn_str != "(): () " {
			fmt.Println(fn_str)
		}
	}
}
