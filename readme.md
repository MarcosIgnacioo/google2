# Google2

this a project *heavily* inspired by tsoding [coogle](https://www.youtube.com/watch?v=wK1HjnwDQng)

so what is this little freak

it is a function/methods greper

you use it this way
```console
google2 '(int) int' /path/to/file.go
```

and it will show the functions that match (with the amazin lesbian distance algorithm)
![chapell roan thank uu](https://media.tenor.com/hHli3fpPZGoAAAAe/chappell-roan-lesbian.png)

### Small xs a little bit example :3

```console
./bin/google2 '(int) int' ./examples/parser.go
[FUZZY SEARCHING]: (int): (int) 
./examples/parser.go:4:1 int_fn(arg int): (int) 
./examples/parser.go:52:1 generics_int_fn[K int](arg int): (int) 
./examples/parser.go:36:1 (X *int) method_int_fn(arg int): (int) 
./examples/parser.go:68:1 generics_2_int_fn[K int | Foo](arg int): (int) 
./examples/parser.go:76:1 generics_2_int_fn_NO_RETURN[K int | Foo](arg int): () 
./examples/parser.go:44:1 (X *int) method_int_fn_NO_RETURN(arg int): () 
./examples/parser.go:60:1 generics_int_fn_NO_RETURN[K int](arg int): () 
./examples/parser.go:12:1 int_fn_NO_RETURN(arg int): () 
./examples/parser.go:20:1 arr_int_fn(arg []int): ([]int) 
./examples/parser.go:28:1 arr_int_fn_NO_RETURN(arg []int): () 
```

cool rite?
you see how it is pretty ordered, with the functions that fit the signature the
most (this is done via string diffing the signature string of the function, so
sometimes the results might not be in the greatest order BUT it will be okay?
doing string diffing gives the advantage to the user of making typos which is
appreaciated)

the reason for this is a lot times you are like trying to use a package and
need a function, while also knowing that kind of stuff the function would
accept (in your hallucination)

# WHY

first of all, great exercise!, and second of all think about this scneario

you just `go getted` some 3 starred package that just does what you really need
for your amazing 10x spa, but dammit this little brat uses snake case as names
in the functions and also really bad ones LIKE WHAT THE HELL WERE YOU THINKING
WHEN YOU PUT A NUMBER BESIDES A FUNCTION NAME? but there is something really
cool in the programming world, even if the author of this GREAT package is a little bit cray cray, he still wrote valid golang code for it, using the same old int
string foo bar blah blah names and those are in your brain so, you *can* use that information, just searching by the
signature... Like *NORMALIZING THE INFORMATION!*, in a little big endian kinda way
if you know what i mean rite? so yeah this is it

# didn't convice you?
well here are some comments of our clients

`a serious project which respects a lot naming conventions of the Golang project` -John Doe
