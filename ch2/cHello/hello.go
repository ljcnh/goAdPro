// hello.go
package main

//void SayHello(const char* s);
import "C"

func main() {
	//println("hello cgo")
	C.SayHello(C.CString("Hello, World\n"))
}
