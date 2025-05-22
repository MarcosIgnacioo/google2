# github.com/Mariiel/classmoods/cmd/api
# for testing individual modules
# go test -run TestMyFunction ./...
compile:
	go build -o ./bin/google2 -gcflags='all=-N -l' ./cmd/google2   2> ./errors.err
	
c:
	go build -o ./bin/google2 -gcflags='all=-N -l' ./cmd/google2
	
run:
	./bin/google2

compile_test:
	go test -v ./cmd/google2 2> ./errors.err

test:
	go test -v ./cmd/google2
	
test_f:
	go test -v -run $(FN) ./cmd/google2

debug:debug
	env --chdir=./cmd/google2 gdlv debug
	
debugt:debugt
	env --chdir=./cmd/google2 gdlv test
	
debug_f:
	env --chdir=./cmd/google2 gdlv test -run $(FN)
	
	
# implement to gdlv debugging a specific test function
# debug_f:debug
# 	dlv test --build-flags='github.com/Mariiel/classmoods/cmd/google2' -- -test.run ^TestSiiaScrapper$
