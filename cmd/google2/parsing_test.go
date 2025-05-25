package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestHardest(t *testing.T) {
	var flags parser.Mode
	flags |= parser.SkipObjectResolution
	files := []string{
		"/usr/lib/go/src/go/ast/ast.go",
		"/usr/lib/go/src/go/ast/commentmap.go",
		"/usr/lib/go/src/go/ast/ast_test.go",
		"/usr/lib/go/src/go/ast/commentmap_test.go",
		"/usr/lib/go/src/go/ast/example_test.go",
		"/usr/lib/go/src/go/ast/filter.go",
		"/usr/lib/go/src/go/ast/filter_test.go",
		"/usr/lib/go/src/go/ast/import.go",
		"/usr/lib/go/src/go/ast/issues_test.go",
		"/usr/lib/go/src/go/ast/print.go",
		"/usr/lib/go/src/go/ast/print_test.go",
		"/usr/lib/go/src/go/ast/resolve.go",
		"/usr/lib/go/src/go/ast/scope.go",
		"/usr/lib/go/src/go/ast/walk.go",
		"/usr/lib/go/src/go/ast/walk_test.go",
		"/usr/lib/go/src/go/types/alias.go",
		"/usr/lib/go/src/go/types/api.go",
		"/usr/lib/go/src/go/types/api_predicates.go",
		"/usr/lib/go/src/go/types/api_test.go",
		"/usr/lib/go/src/go/types/array.go",
		"/usr/lib/go/src/go/types/assignments.go",
		"/usr/lib/go/src/go/types/badlinkname.go",
		"/usr/lib/go/src/go/types/basic.go",
		"/usr/lib/go/src/go/types/builtins.go",
		"/usr/lib/go/src/go/types/builtins_test.go",
		"/usr/lib/go/src/go/types/call.go",
		"/usr/lib/go/src/go/types/chan.go",
		"/usr/lib/go/src/go/types/check.go",
		"/usr/lib/go/src/go/types/check_test.go",
		"/usr/lib/go/src/go/types/commentMap_test.go",
		"/usr/lib/go/src/go/types/const.go",
		"/usr/lib/go/src/go/types/context.go",
		"/usr/lib/go/src/go/types/context_test.go",
		"/usr/lib/go/src/go/types/conversions.go",
		"/usr/lib/go/src/go/types/decl.go",
		"/usr/lib/go/src/go/types/errorcalls_test.go",
		"/usr/lib/go/src/go/types/errors.go",
		"/usr/lib/go/src/go/types/errors_test.go",
		"/usr/lib/go/src/go/types/errsupport.go",
		"/usr/lib/go/src/go/types/eval.go",
		"/usr/lib/go/src/go/types/eval_test.go",
		"/usr/lib/go/src/go/types/example_test.go",
		"/usr/lib/go/src/go/types/expr.go",
		"/usr/lib/go/src/go/types/exprstring.go",
		"/usr/lib/go/src/go/types/exprstring_test.go",
		"/usr/lib/go/src/go/types/format.go",
		"/usr/lib/go/src/go/types/gccgosizes.go",
		"/usr/lib/go/src/go/types/gcsizes.go",
		"/usr/lib/go/src/go/types/generate.go",
		"/usr/lib/go/src/go/types/generate_test.go",
		"/usr/lib/go/src/go/types/gotype.go",
		"/usr/lib/go/src/go/types/hilbert_test.go",
		"/usr/lib/go/src/go/types/index.go",
		"/usr/lib/go/src/go/types/infer.go",
		"/usr/lib/go/src/go/types/initorder.go",
		"/usr/lib/go/src/go/types/instantiate.go",
		"/usr/lib/go/src/go/types/instantiate_test.go",
		"/usr/lib/go/src/go/types/interface.go",
		"/usr/lib/go/src/go/types/issues_test.go",
		"/usr/lib/go/src/go/types/iter.go",
		"/usr/lib/go/src/go/types/labels.go",
		"/usr/lib/go/src/go/types/literals.go",
		"/usr/lib/go/src/go/types/lookup.go",
		"/usr/lib/go/src/go/types/lookup_test.go",
		"/usr/lib/go/src/go/types/main_test.go",
		"/usr/lib/go/src/go/types/map.go",
		"/usr/lib/go/src/go/types/methodset.go",
		"/usr/lib/go/src/go/types/methodset_test.go",
		"/usr/lib/go/src/go/types/mono.go",
		"/usr/lib/go/src/go/types/mono_test.go",
		"/usr/lib/go/src/go/types/named.go",
		"/usr/lib/go/src/go/types/named_test.go",
		"/usr/lib/go/src/go/types/object.go",
		"/usr/lib/go/src/go/types/object_test.go",
		"/usr/lib/go/src/go/types/objset.go",
		"/usr/lib/go/src/go/types/operand.go",
		"/usr/lib/go/src/go/types/package.go",
		"/usr/lib/go/src/go/types/pointer.go",
		"/usr/lib/go/src/go/types/predicates.go",
		"/usr/lib/go/src/go/types/README.md",
		"/usr/lib/go/src/go/types/recording.go",
		"/usr/lib/go/src/go/types/resolver.go",
		"/usr/lib/go/src/go/types/resolver_test.go",
		"/usr/lib/go/src/go/types/return.go",
		"/usr/lib/go/src/go/types/scope.go",
		"/usr/lib/go/src/go/types/scope2.go",
		"/usr/lib/go/src/go/types/scope2_test.go",
		"/usr/lib/go/src/go/types/selection.go",
		"/usr/lib/go/src/go/types/self_test.go",
		"/usr/lib/go/src/go/types/signature.go",
		"/usr/lib/go/src/go/types/sizeof_test.go",
		"/usr/lib/go/src/go/types/sizes.go",
		"/usr/lib/go/src/go/types/sizes_test.go",
		"/usr/lib/go/src/go/types/slice.go",
		"/usr/lib/go/src/go/types/stdlib_test.go",
		"/usr/lib/go/src/go/types/stmt.go",
		"/usr/lib/go/src/go/types/struct.go",
		"/usr/lib/go/src/go/types/subst.go",
		"/usr/lib/go/src/go/types/termlist.go",
		"/usr/lib/go/src/go/types/termlist_test.go",
		"/usr/lib/go/src/go/types/token_test.go",
		"/usr/lib/go/src/go/types/tuple.go",
		"/usr/lib/go/src/go/types/type.go",
		"/usr/lib/go/src/go/types/typelists.go",
		"/usr/lib/go/src/go/types/typeparam.go",
		"/usr/lib/go/src/go/types/typeset.go",
		"/usr/lib/go/src/go/types/typeset_test.go",
		"/usr/lib/go/src/go/types/typestring.go",
		"/usr/lib/go/src/go/types/typestring_test.go",
		"/usr/lib/go/src/go/types/typeterm.go",
		"/usr/lib/go/src/go/types/typeterm_test.go",
		"/usr/lib/go/src/go/types/typexpr.go",
		"/usr/lib/go/src/go/types/under.go",
		"/usr/lib/go/src/go/types/unify.go",
		"/usr/lib/go/src/go/types/union.go",
		"/usr/lib/go/src/go/types/universe.go",
		"/usr/lib/go/src/go/types/util.go",
		"/usr/lib/go/src/go/types/util_test.go",
		"/usr/lib/go/src/go/types/validtype.go",
		"/usr/lib/go/src/go/types/version.go",
	}

	for _, file := range files {
		var fset token.FileSet
		file_ast, _ := parser.ParseFile(&fset, file, nil, flags)
		func_declarations := file_ast.Decls
		for _, declaration := range func_declarations {
			switch v := declaration.(type) {
			case *ast.FuncDecl:
				{
					function := parse_ast_func_decl(&fset, v)
					fmt.Println(function)
				}
			}
		}
	}
}

func TestQueryInput(t *testing.T) {
	var flags parser.Mode
	flags |= parser.SkipObjectResolution
	queries := []string{
		"(int) int{return foo}",
		"(*Comment) Pos() (token.Pos)",
		"(*Comment) End() (token.Pos)",
		"(*CommentGroup) Pos() (token.Pos)",
		"(*CommentGroup) End() (token.Pos)",
		"isWhitespace(byte) (bool)",
		"stripTrailingWhitespace(string) (string)",
		"(*CommentGroup) Text() (string)",
		"isDirective(string) (bool)",
		"(*Field) Pos() (token.Pos)",
		"(*Field) End() (token.Pos)",
		"(*FieldList) Pos() (token.Pos)",
		"(*FieldList) End() (token.Pos)",
		"(*FieldList) NumFields() (int)",
		"(*BadExpr) Pos() (token.Pos)",
		"(*Ident) Pos() (token.Pos)",
		"(*Ellipsis) Pos() (token.Pos)",
		"(*BasicLit) Pos() (token.Pos)",
		"(*FuncLit) Pos() (token.Pos)",
		"(*CompositeLit) Pos() (token.Pos)",
		"(*ParenExpr) Pos() (token.Pos)",
		"(*SelectorExpr) Pos() (token.Pos)",
		"(*IndexExpr) Pos() (token.Pos)",
		"(*IndexListExpr) Pos() (token.Pos)",
		"(*SliceExpr) Pos() (token.Pos)",
		"(*TypeAssertExpr) Pos() (token.Pos)",
		"(*CallExpr) Pos() (token.Pos)",
		"(*StarExpr) Pos() (token.Pos)",
		"(*UnaryExpr) Pos() (token.Pos)",
		"(*BinaryExpr) Pos() (token.Pos)",
		"(*KeyValueExpr) Pos() (token.Pos)",
		"(*ArrayType) Pos() (token.Pos)",
		"(*StructType) Pos() (token.Pos)",
		"(*FuncType) Pos() (token.Pos)",
		"(*InterfaceType) Pos() (token.Pos)",
		"(*MapType) Pos() (token.Pos)",
		"(*ChanType) Pos() (token.Pos)",
		"(*BadExpr) End() (token.Pos)",
		"(*Ident) End() (token.Pos)",
		"(*Ellipsis) End() (token.Pos)",
		"(*BasicLit) End() (token.Pos)",
		"(*FuncLit) End() (token.Pos)",
		"(*CompositeLit) End() (token.Pos)",
		"(*ParenExpr) End() (token.Pos)",
		"(*SelectorExpr) End() (token.Pos)",
		"(*IndexExpr) End() (token.Pos)",
		"(*IndexListExpr) End() (token.Pos)",
		"(*SliceExpr) End() (token.Pos)",
		"(*TypeAssertExpr) End() (token.Pos)",
		"(*CallExpr) End() (token.Pos)",
		"(*StarExpr) End() (token.Pos)",
		"(*UnaryExpr) End() (token.Pos)",
		"(*BinaryExpr) End() (token.Pos)",
		"(*KeyValueExpr) End() (token.Pos)",
		"(*ArrayType) End() (token.Pos)",
		"(*StructType) End() (token.Pos)",
		"(*FuncType) End() (token.Pos)",
		"(*InterfaceType) End() (token.Pos)",
		"(*MapType) End() (token.Pos)",
		"(*ChanType) End() (token.Pos)",
		"(*BadExpr) exprNode() ()",
		"(*Ident) exprNode() ()",
		"(*Ellipsis) exprNode() ()",
		"(*BasicLit) exprNode() ()",
		"(*FuncLit) exprNode() ()",
		"(*CompositeLit) exprNode() ()",
		"(*ParenExpr) exprNode() ()",
		"(*SelectorExpr) exprNode() ()",
		"(*IndexExpr) exprNode() ()",
		"(*IndexListExpr) exprNode() ()",
		"(*SliceExpr) exprNode() ()",
		"(*TypeAssertExpr) exprNode() ()",
		"(*CallExpr) exprNode() ()",
		"(*StarExpr) exprNode() ()",
		"(*UnaryExpr) exprNode() ()",
		"(*BinaryExpr) exprNode() ()",
		"(*KeyValueExpr) exprNode() ()",
		"(*ArrayType) exprNode() ()",
		"(*StructType) exprNode() ()",
		"(*FuncType) exprNode() ()",
		"(*InterfaceType) exprNode() ()",
		"(*MapType) exprNode() ()",
		"(*ChanType) exprNode() ()",
		"NewIdent(name string) (*Ident)",
		"IsExported(name string) (bool)",
		"(id *Ident) IsExported() (bool)",
		"(id *Ident) String() (string)",
		"(s *BadStmt) Pos() (token.Pos)",
		"(s *DeclStmt) Pos() (token.Pos)",
		"(s *EmptyStmt) Pos() (token.Pos)",
		"(s *LabeledStmt) Pos() (token.Pos)",
		"(s *ExprStmt) Pos() (token.Pos)",
		"(s *SendStmt) Pos() (token.Pos)",
		"(s *IncDecStmt) Pos() (token.Pos)",
		"(s *AssignStmt) Pos() (token.Pos)",
		"(s *GoStmt) Pos() (token.Pos)",
		"(s *DeferStmt) Pos() (token.Pos)",
		"(s *ReturnStmt) Pos() (token.Pos)",
		"(s *BranchStmt) Pos() (token.Pos)",
		"(s *BlockStmt) Pos() (token.Pos)",
		"(s *IfStmt) Pos() (token.Pos)",
		"(s *CaseClause) Pos() (token.Pos)",
		"(s *SwitchStmt) Pos() (token.Pos)",
		"(s *TypeSwitchStmt) Pos() (token.Pos)",
		"(s *CommClause) Pos() (token.Pos)",
		"(s *SelectStmt) Pos() (token.Pos)",
		"(s *ForStmt) Pos() (token.Pos)",
		"(s *RangeStmt) Pos() (token.Pos)",
		"(s *BadStmt) End() (token.Pos)",
		"(s *DeclStmt) End() (token.Pos)",
		"(s *EmptyStmt) End() (token.Pos)",
		"(s *LabeledStmt) End() (token.Pos)",
		"(s *ExprStmt) End() (token.Pos)",
		"(s *SendStmt) End() (token.Pos)",
		"(s *IncDecStmt) End() (token.Pos)",
		"(s *AssignStmt) End() (token.Pos)",
		"(s *GoStmt) End() (token.Pos)",
		"(s *DeferStmt) End() (token.Pos)",
		"(s *ReturnStmt) End() (token.Pos)",
		"(s *BranchStmt) End() (token.Pos)",
		"(s *BlockStmt) End() (token.Pos)",
		"(s *IfStmt) End() (token.Pos)",
		"(s *CaseClause) End() (token.Pos)",
		"(s *SwitchStmt) End() (token.Pos)",
		"(s *TypeSwitchStmt) End() (token.Pos)",
		"(s *CommClause) End() (token.Pos)",
		"(s *SelectStmt) End() (token.Pos)",
		"(s *ForStmt) End() (token.Pos)",
		"(s *RangeStmt) End() (token.Pos)",
		"(*BadStmt) stmtNode() ()",
		"(*DeclStmt) stmtNode() ()",
		"(*EmptyStmt) stmtNode() ()",
		"(*LabeledStmt) stmtNode() ()",
		"(*ExprStmt) stmtNode() ()",
		"(*SendStmt) stmtNode() ()",
		"(*IncDecStmt) stmtNode() ()",
		"(*AssignStmt) stmtNode() ()",
		"(*GoStmt) stmtNode() ()",
		"(*DeferStmt) stmtNode() ()",
		"(*ReturnStmt) stmtNode() ()",
		"(*BranchStmt) stmtNode() ()",
		"(*BlockStmt) stmtNode() ()",
		"(*IfStmt) stmtNode() ()",
		"(*CaseClause) stmtNode() ()",
		"(*SwitchStmt) stmtNode() ()",
		"(*TypeSwitchStmt) stmtNode() ()",
		"(*CommClause) stmtNode() ()",
		"(*SelectStmt) stmtNode() ()",
		"(*ForStmt) stmtNode() ()",
		"(*RangeStmt) stmtNode() ()",
		"(s *ImportSpec) Pos() (token.Pos)",
		"(s *ValueSpec) Pos() (token.Pos)",
		"(s *TypeSpec) Pos() (token.Pos)",
		"(s *ImportSpec) End() (token.Pos)",
		"(s *ValueSpec) End() (token.Pos)",
		"(s *TypeSpec) End() (token.Pos)",
		"(*ImportSpec) specNode() ()",
		"(*ValueSpec) specNode() ()",
		"(*TypeSpec) specNode() ()",
	}

	for _, user_query := range queries {
		query_ast, err := parse_user_function_query(user_query)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(parse_ast_func_decl(nil, query_ast))
		}
	}
}

func build_queries(t string) []string {
	return []string{
		fmt.Sprintf("(%s)", t),
		fmt.Sprintf("(%s)%s", t, t),

		fmt.Sprintf("(*%s)(%s)", t, t),
		fmt.Sprintf("(*%s)(%s)%s", t, t, t),

		fmt.Sprintf("([]%s)", t),
		fmt.Sprintf("([]%s)%s", t, t),

		fmt.Sprintf("[%s](%s)", t, t),
		fmt.Sprintf("[%s](%s)%s", t, t, t),
	}
}

func TestGoogle2(t *testing.T) {
	var flags parser.Mode
	path := "/home/marcig/personal/fun/google2/examples/parser.go"
	flags |= parser.SkipObjectResolution

	type test struct {
		queries  []string
		expected string //todo xd i think i can make the golang compiler just spit the struct form of the results, rn i know the application works and i dont plan to add new stuff to it BNUUT if i was doing something for a long time development doing this would be the way i think, because yeah, i know the tests will pass rn but if i break something somewhere i would catch them pretty quickly i think
	}

	types := []string{"int", "string", "float32", "float64", "bool", "byte", "rune"}
	queries := make([]test, 0, 12)
	for _, t := range types {
		queries = append(queries, test{build_queries(t), ""})
	}

	for _, user_query := range queries {
		for _, query := range user_query.queries {
			functions := google2(query, path, 5)
			output := stringify_array(functions)
			fmt.Println(output)
			// if output != user_query.expected {
			// 	t.Error("oops in this query", query)
			// }
		}
	}
}
