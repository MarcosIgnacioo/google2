package parser

import (
	"go/ast"
	"go/token"
	"strings"
	"testing"

	"github.com/Marciglez/google2/utils"
)

func stringify_func_declaration(fn *ast.FuncDecl) string {
	if fn == nil {
		return ""
	}
	var builder strings.Builder
	if fn.Recv != nil {
		builder.WriteString("(")
		builder.WriteString(utils.StringifyFieldListNoNames(fn.Recv))
		builder.WriteString(") ")
	}
	var signature *ast.FuncType
	signature = fn.Type

	if fn.Name != nil {
		builder.WriteString("func ")
		builder.WriteString(fn.Name.Name)
	}

	if signature.TypeParams != nil {
		builder.WriteString("[")
		builder.WriteString(utils.StringifyFieldListNoNames(signature.TypeParams))
		builder.WriteString("]")
	}

	builder.WriteString("(")
	builder.WriteString(utils.StringifyFieldListNoNames(signature.Params))
	builder.WriteString(") ")
	builder.WriteString(utils.StringifyFieldListNoNames(signature.Results))

	return builder.String()
}

func TestParseEasyQuery(t *testing.T) {
	t.Log("hello world")
	type result struct {
		input    string
		expected string
	}
	queries := []result{
		{
			"[interface{ String() string }]()",
			"",
		},
		{
			"([]int)",
			"",
		},
		{
			"(int, []foo , bar) string",
			"(int, []foo, bar) string",
		},
		{
			"(p *parser) (int, []foo , bar) string",
			"[T interface{ String() string }](array []T)",
		},
		{
			"[T string] (array []T) string",
			"",
		},
	}
	for _, query := range queries {
		var fset token.FileSet
		var flags Mode
		flags |= SkipObjectResolution
		fn, err := ParseQuery(&fset, query.input, flags)
		if err != nil {
			t.Log(err.Error())
		} else {
			t.Log(stringify_func_declaration(fn))
		}
	}
}
