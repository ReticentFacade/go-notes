package main

/*
#cgo LDFLAGS: -L./lib -lhello
#include "./lib/hello.h"
*/
import "C"

func main() {
	C.hello(C.String("world"))
	C.yell(C.String("YOU!"))
}
