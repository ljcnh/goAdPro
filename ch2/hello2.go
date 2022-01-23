package main

//static const char* cs = "hello";
import "C"
import "github.com/ljcnh/goAdPro/ch2/cgo_helper"

//报错
func main() {
	cgo_helper.PrintCString(C.cs) // error
}
