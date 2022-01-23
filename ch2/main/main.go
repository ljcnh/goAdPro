package main

import "C"

import (
	"fmt"
	_ "github.com/ljcnh/goAdPro/ch2/go2c_1"
)

func main() {
	println("Done")
}

//export goPrintln
func goPrintln(s *C.char) {
	fmt.Println("goPrintln:", C.GoString(s))
}
