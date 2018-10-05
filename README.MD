# Simple Web Assembly Golang Demo

This project takes user input from the front end, passes it to a golang function, which encrypts the text with RSA, and
returns it to the front end. The go function communicates with the front end with the help of Web Assembly and Javascript.


This project is not intended to be used for anything other than to demonstrate Web Assembly with Golang. Do not use this
project to encrypt anything important.

index.html and wasm_exec.js are taken from the official golang repo on github.
https://github.com/golang/go/tree/master/misc/wasm


## How to run it

on windows:

```set GOOS=js```
```set GOARCH=wasm```
```go build -o main.go main.wasm```
```go run server.go```