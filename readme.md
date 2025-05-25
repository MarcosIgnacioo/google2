https://www.youtube.com/watch?v=fv-wlo8yVhk
because golang is an AMAZING language it allow us to do some cool stuff like this

func foo(a,b,c, int) float32

which is good for programming but for our amazin query language is kinda of a pain in the asssss

and u might be wondering WHYYY
because in the query language we just wanna specify the types

and this function translates to in our stringifying
(int) float32
which is incorrect!!!

it should be 
(int, int, int) float32
so when signature stringifying even if we do not account the names we also need to keep in mind the amount of parameters of
the same type so it outputs the right thing, which is the thing above
so thats our goal today

why maybe the lesbian distance formula might not be the fit for this

```go
sort.Slice old but a lot of text
if user_function_search.Generics != nil {
	generics := stringify_field_list_without_names(user_function_search.Generics)
	generics_left := stringify_field_list_without_names(left.Generics)
	generics_right := stringify_field_list_without_names(right.Generics)
	left_changes_needed_for_match += lev_distance(generics_left, generics)
	right_changes_needed_for_match += lev_distance(generics_right, generics)
}
if user_function_search.Receiver != nil {
	receiver := stringify_field_list_without_names(user_function_search.Receiver)
	receiver_left := stringify_field_list_without_names(left.Receiver)
	receiver_right := stringify_field_list_without_names(right.Receiver)
	left_changes_needed_for_match += lev_distance(receiver_left, receiver)
	right_changes_needed_for_match += lev_distance(receiver_right, receiver)
}

if user_function_search.Params.List != nil {
	params := stringify_field_list_without_names(user_function_search.Params)
	params_left := stringify_field_list_without_names(left.Params)
	params_right := stringify_field_list_without_names(right.Params)
	left_changes_needed_for_match += lev_distance(params_left, params)
	right_changes_needed_for_match += lev_distance(params_right, params)
}

if user_function_search.Results.List != nil {
	results := stringify_field_list_without_names(user_function_search.Results)
	results_left := stringify_field_list_without_names(left.Results)
	results_right := stringify_field_list_without_names(right.Results)
	left_changes_needed_for_match += lev_distance(results_left, results)
	right_changes_needed_for_match += lev_distance(results_right, results)
}

```

because if i put this input (float) int

the sumf is listed wayyy below, which kinda makes sense because the parameter 

yeah in the golang case specifically the lesbian distance doesnt work quite well because when we have functions that just kj kj a

take this example, here we need to specify the string and int parameters so we get the AddFile method which is the one the user wanna know about, which is kinda a bummer, but because of the existence of the method syntax we need to compare the methods types and generics types, param types and result types separatly and if they have like a lesbian distance (because we will use the string representation of those types im not doing anything deeper than that) and they have like 2 or something we call it a match and give them certain weigth, we might priorize the parts the user gives us like if they only give the method well we give the points to the method thing or generic types, is a kinda cumbersome because well what happens if the function doenst exist as a method but it does exist as a regular function, we do not show it? i think there we could manage like a total matches or something, that if ther e have not been that many matches we lower the requsisites for matchiung fucntions so we just add some that might be like that or flag them as a MAYBE and the other ones we did compare hard as HIGHLY or something like that idk
```go
2025/05/20 21:03:34 (string, int): (*File) 
/usr/lib/go/src/go/token/position.go:594:1 searchInts(a int[], x int): (int) 
/usr/lib/go/src/go/token/position.go:270:1 (f *File) fixOffset(offset int): (int) 
/usr/lib/go/src/go/token/position.go:527:1 searchFiles(a *File[], x int): (int) 
/usr/lib/go/src/go/token/position.go:464:1 (s *FileSet) AddFile(filename string, base, size int): (*File) 
/usr/lib/go/src/go/token/position.go:302:1 (f *File) Pos(offset int): (Pos) 
/usr/lib/go/src/go/token/position.go:150:1 (f *File) MergeLine(line int): () 
/usr/lib/go/src/go/token/position.go:225:1 (f *File) LineStart(line int): (Pos) 
/usr/lib/go/src/go/token/position.go:138:1 (f *File) AddLine(offset int): () 
/usr/lib/go/src/go/token/position.go:325:1 searchLineInfos(a lineInfo[], x int): (int) 
/usr/lib/go/src/go/token/position.go:118:1 (f *File) Base(): (int) 
```

for the function we are looking for is a float and byte is closer to it than int, (both require changing the four characters, maybe 3 if the lesbian distance tryhards) but in general some stuff is just not quite good for a fuzzy search when i kinda expect to show me numbery parameters, kinda of a bummer but it will be more fun to implement right?

also the float32 return type doesnt help because for it to be a match requires a whole change which is not good for the filtering
```console
2025/05/20 18:27:47 (float): (int) 
/home/marcig/personal/fun/google2/internal/dummy.go:1147:1 sum(x, y int): (int) 
/home/marcig/personal/fun/google2/internal/dummy.go:82:1 isWhitespace(ch byte): (bool) 
/home/marcig/personal/fun/google2/internal/dummy.go:1109:1 IsGenerated(file *File): (bool) 
/home/marcig/personal/fun/google2/internal/dummy.go:1137:1 Unparen(e Expr): (Expr) 
/home/marcig/personal/fun/google2/internal/dummy.go:596:1 IsExported(name string): (bool) 
/home/marcig/personal/fun/google2/internal/dummy.go:84:1 stripTrailingWhitespace(s string): (string) 
/home/marcig/personal/fun/google2/internal/dummy.go:564:1 (*Ident) exprNode(): () 
/home/marcig/personal/fun/google2/internal/dummy.go:163:1 isDirective(c string): (bool) 
/home/marcig/personal/fun/google2/internal/dummy.go:1151:1 sumf(x, y int): (float32) 
/home/marcig/personal/fun/google2/internal/dummy.go:877:1 (*IfStmt) stmtNode(): () 
```

# TODO

this week

add testings for query inputs and file inputs and expected outputs[]
change lev distance to be just a for loop instead of a recursive stuff[]
clean up the main code and make two builds one for user and one for dkjbugging[]

future features
make a custom parser so we dont rely on the golang one, and also we can allow a little bit of more goofy syntax by the user and still resulting in a correct AST or at least kinda correct []
make a hashmap with the standard golang packages or a way to just indicate a package name and search in all the files u know[]
allow for math types like decimal for any kind of decimal/floating point  value number []




Golang has the Node interface which, from there we have 3 kind of nodes
        Expressions and type nodes
        statement nodes
        declaration nodes

We want to find the functions in source file so we are insterested in the declaration nodes


And the `Decl` interface is the one for the declarations (yay) and the struct that implements a Declariont interface that is for the functions is the one we need for finding the functions
declNode()

```go
    /usr/lib/go/src/go/ast/ast.go:994
	FuncDecl struct {
		Doc  *CommentGroup // associated documentation; or nil
		Recv *FieldList    // receiver (methods); or nil (functions)
		Name *Ident        // function/method name
		Type *FuncType     // function signature: type and value parameters, results, and position of "func" keyword
		Body *BlockStmt    // function body; or nil for external (non-Go) function
	}
```

Great now we have access to the function declaration of a file

```go
// TODO: Change this please to just use the FuncDecl from teh ast package
type Function2 struct {
	FunctionName string
	Params       []Param
	Results      []Result
}```
