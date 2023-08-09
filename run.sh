#! /bin/bash

GOOS=js GOARCH=wasm go test -c -o test.wasm -run TestPersist

echo visit http://localhost:8000/wasm_exec.html
go run ./webserver.go
