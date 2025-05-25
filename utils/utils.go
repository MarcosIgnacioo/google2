package utils

import (
	"fmt"
	"go/ast"
	"log"
	"strings"
)

func StringifyExpression(expression ast.Expr) string {
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
			expr := StringifyExpression(v.X)
			return fmt.Sprintf("%s", expr)
		}
	case *ast.BinaryExpr:
		{
			left := StringifyExpression(v.X)
			right := StringifyExpression(v.Y)
			return fmt.Sprintf("%s %s %s", left, v.Op.String(), right)
		}
	case *ast.FuncLit:
		{
			return StringifyExpression(v.Type)
		}
	case *ast.InterfaceType:
		{
			var builder strings.Builder
			builder.WriteString("interface{")
			if v.Methods != nil {
				builder.WriteString(StringifyFieldList(v.Methods))
			}
			builder.WriteString("}")
			return builder.String()
		}
	case *ast.BadExpr:
		log.Fatalf("Expression not supported jijiji %v", v)
		return "[NOT SUPPORTED]"
	case *ast.IndexExpr:
		{
			exp := StringifyExpression(v.X)
			index := StringifyExpression(v.Index)
			return fmt.Sprintf("%s[%s]", exp, index)
		}
	case *ast.Ident:
		{
			return v.Name
		}
	case *ast.Ellipsis:
		{
			elipsis_type := StringifyExpression(v.Elt)
			return fmt.Sprintf("...%s", elipsis_type)
		}
	case *ast.SelectorExpr:
		{
			selector_expr := StringifyExpression(v.X)
			return fmt.Sprintf("%s.%s", selector_expr, v.Sel.Name)
		}
	case *ast.StarExpr:
		{
			star_expr := StringifyExpression(v.X)
			return fmt.Sprintf("*%s", star_expr)
		}
	case *ast.ArrayType:
		{
			var count string
			if v.Len != nil {
				count = StringifyExpression(v.Len)
			}
			array_type := StringifyExpression(v.Elt)
			return fmt.Sprintf("%s[%s]", array_type, count)
		}
	case *ast.FuncType:
		{
			var params strings.Builder
			var results strings.Builder
			AppendFieldList(&params, v.Params)
			AppendFieldList(&results, v.Results)
			// this looks goofy in an interface i think
			return fmt.Sprintf("(%s) %s", params.String(), results.String())
			// return fmt.Sprintf("func (%s) %s", params.String(), results.String())
		}
	case *ast.MapType:
		{
			key := StringifyExpression(v.Key)
			value := StringifyExpression(v.Value)
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

func AppendFieldList(builder *strings.Builder, field_list_container *ast.FieldList) {
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
		param_type := StringifyExpression(param.Type)
		_, err := builder.WriteString(fmt.Sprintf("%s", param_type))
		if err != nil {
			fmt.Println("[ERROR]:an error has happened unu")
		}
	}
}

// todo make just two functions one for stringify without names and one with them, because
// the if is making me itchy
func StringifyFieldList(field_list_container *ast.FieldList) string {
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
		param_type := StringifyExpression(param.Type)
		builder.WriteString(fmt.Sprintf("%s", param_type))
	}

	return builder.String()
}

func StringifyFieldListNoNames(field_list_container *ast.FieldList) string {
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
		param_type := StringifyExpression(param.Type)
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
