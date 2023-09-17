package main

// #cgo CFLAGS: -I/Users/Reticent/Desktop/Work/ProjectWorkspace/github/go-notes/rust-plus-golang/lib
// #cgo LDFLAGS: -L/Users/Reticent/Desktop/Work/ProjectWorkspace/github/go-notes/rust-plus-golang/lib/hello/target/release -lhello
// #include "hello.h"
import "C"



func main() {
	C.hello(C.CString("world"))
	C.yell(C.CString("YOU!"))
}
