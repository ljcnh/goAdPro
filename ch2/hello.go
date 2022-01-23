package main

//void SayHello(char* s);
import "C"

import "fmt"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}

//export SayHello
func SayHello(s string) {
	fmt.Print(C.GoString(s))
}
