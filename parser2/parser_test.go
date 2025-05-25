package parser2

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/Marciglez/google2/utils"
)

func parse_raw_string(source string) (f *ast.File, err error) {
	var fset token.FileSet
	var flags parser.Mode
	var file_path string

	flags |= parser.SkipObjectResolution
	// yeah its scrapping the parser time
	// because we do this silly thing we dont actually the possibility to actually
	// pass generics or method declaration which is not great !!!
	// sooo we might aswell write (steal) our own function declaration go parser
	source = fmt.Sprintf("package main\n %s", source)
	return parser.ParseFile(&fset, file_path, source, flags)
}

func AreFieldEquals(this, that *ast.Field) bool {
	equals := true

	a_len := len(this.Names)
	b_len := len(that.Names)

	if a_len != b_len {
		return false
	}

	a_t := utils.StringifyExpression(this.Type)
	b_t := utils.StringifyExpression(that.Type)

	if a_t != b_t {
		return false
	}

	for i := range this.Names {
		a := this.Names[i]
		b := that.Names[i]
		if !AreIdentEquals(a, b) {
			equals = false
			break
		}
	}

	return equals
}

func AreaFieldListEquals(this, that *ast.FieldList) bool {
	equals := true
	if this == nil && that == nil {
		return true
	}
	if this.List == nil || that.List == nil {
		return false
	}
	if len(this.List) != len(that.List) {
		return false
	}

	for i := range this.List {
		a := this.List[i]
		b := that.List[i]
		if !AreFieldEquals(a, b) {
			equals = false
			break
		}
	}

	return equals
}

func AreIdentEquals(this, that *ast.Ident) bool {
	// the only wway this will be true is if both are nil
	if this == that {
		return true
	}

	// if just one is nil we just return false
	if this == nil || that == nil {
		return false
	}

	return this.Name == that.Name
}

func AreFuncDeclEquals(this, that *ast.FuncDecl) bool {
	if this == that {
		return true
	}
	if this == nil || that == nil {
		return false
	}
	if !AreIdentEquals(this.Name, that.Name) {
		return false
	}

	method_eq := AreaFieldListEquals(this.Recv, that.Recv)
	generics_eq := AreaFieldListEquals(this.Type.TypeParams, that.Type.TypeParams)
	params_eq := AreaFieldListEquals(this.Type.Params, that.Type.Params)
	result_eq := AreaFieldListEquals(this.Type.Results, that.Type.Results)

	return method_eq && generics_eq && params_eq && result_eq
}

func stringify_func_declaration(fn *ast.FuncDecl) string {
	if fn == nil {
		return ""
	}
	var builder strings.Builder
	if fn.Recv != nil {
		builder.WriteString("(")
		builder.WriteString(utils.StringifyFieldList(fn.Recv))
		builder.WriteString(") ")
	}
	var signature *ast.FuncType
	signature = fn.Type

	builder.WriteString("func ")
	builder.WriteString(fn.Name.Name)

	if signature.TypeParams != nil {
		builder.WriteString("[")
		builder.WriteString(utils.StringifyFieldList(signature.TypeParams))
		builder.WriteString("]")
	}

	builder.WriteString("(")
	builder.WriteString(utils.StringifyFieldList(signature.Params))
	builder.WriteString(") ")
	builder.WriteString(utils.StringifyFieldList(signature.Results))

	return builder.String()
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

func TestParseRegularQuery(t *testing.T) {
	fn, err := parse_user_function_query("func foo(foo int, int, float32) string")
	t.Log(err.Error())
	// p := NewParser2()
	fmt.Println(stringify_func_declaration(fn))
}

func TestParseComplexQueries(t *testing.T) {
	// comparing := "func foo[T interface{ String() string }](array []T)"

	queries := []string{
		"func print_array[T interface{ String() string }](array []T)",
	}
	for _, query := range queries {
		fn_decl, err := parse_user_function_query(query)
		// fn_decl2, err := parse_user_function_query(comparing)
		if err != nil {
			t.Log(err.Error())
		}
		t.Log(stringify_func_declaration(fn_decl))
		// t.Log(stringify_func_declaration(fn_decl2))
	}
}

func TestComparing(t *testing.T) {
	type Comparing struct {
		title    string
		this     string
		that     string
		expected bool
	}
	comparing := []Comparing{
		{
			"print_array == print_array",
			"func print_array[T interface{ String() string }](array []T)",
			"func print_array[T interface{ String() string }](array []T)",
			true,
		},
		{
			"print_array != print_array (implement different interface in generics)",
			"func print_array[T interface{ String() string }](array []T)",
			"func print_array[T interface{ Number() int }](array []T)",
			false,
		},
		{
			"print_array != print_array (the other one returns int and error)",
			"func print_array[T interface{ String() string }](array []T)",
			"func print_array[T interface{ String() (int, error) }](array []T)",
			false,
		},
		{
			"print_array != print_array (the other one returns has another generic type)",
			"func print_array[T interface{ String() string }](array []T)",
			"func print_array[T interface{ String() string }, V string](array []T, foo V)",
			false,
		},
		{
			"print_array != print_array (the other one has another method in the genric T)",
			"func print_array[T interface{ String() string }](array []T)",
			"func print_array[T interface{ String() string, Number() int }](array []T, foo V)",
			false,
		},
		{
			"print_array != Clamp",
			"func (p *Parser2) expect(tok token.Token) token.Pos",
			"func Clamp(value, min, max float32) float32",
			false,
		},
		{
			"Lerp == Lerp",
			"func Lerp(start, end, amount float32) float32",
			"func Lerp(start, end, amount float32) float32",
			true,
		},
		{
			"Lerp != Lerp (one has one less arg)",
			"func Lerp(start, end, amount float32) float32",
			"func Lerp(start, end float32) float32",
			false,
		},
		{
			"Lerp == Lerp (one returns float64 and the other float32)",
			"func Lerp(start, end, amount float32) float64",
			"func Lerp(start, end, amount float32) float32",
			false,
		},
		{
			"Remap != Remap",
			"func (p *Parser) Remap(value, inputStart, inputEnd, outputStart, outputEnd float32) float32",
			"func (a *ast.Ast) Remap(value, inputStart, inputEnd, outputStart, outputEnd float32) float32",
			false,
		},
		{
			"Lerp == Lerp",
			"func Lerp(start, end, amount float32) float32",
			"func Lerp(start, end, amount float32) float32",
			true,
		},
	}
	for _, compa := range comparing {
		this, err := parse_user_function_query(compa.this)
		that, err := parse_user_function_query(compa.that)
		if err != nil {
			t.Log(err.Error())
		}
		result := AreFuncDeclEquals(this, that)
		if result == compa.expected {
			t.Log(compa.title)
			// t.Logf("cmp(%s, %s)", stringify_func_declaration(this), stringify_func_declaration(that))
		} else {
			// maybe implement a lesbian distance here so i can
			// show what are the diffs
			// okay now implenting giff would be pretty cool
			t.Log("----------------------------------------")
			t.Logf("expected: %v got:%v", compa.expected, result)
			t.Errorf("\n\tcomp(\n\t\t%s,\n\t\t%s,\n\t)", stringify_func_declaration(this), stringify_func_declaration(that))
		}
	}
}

func TestParseFields(t *testing.T) {
	p := NewParser2("(interface {} int boo, float popo, uwu nya caca)")
	p.next()
	fl, _ := p.parse_field_list(token.RPAREN)
	str := utils.StringifyFieldList(fl)
	t.Log(str)
}
