package main

/*
#include <stdlib.h>
void* makeslice(size_t memsize) {
    return malloc(memsize);
}
*/
import "C"
import (
	"unsafe"
)

func makeByteSlize(n int) []byte {
	p := C.makeslice(C.size_t(n))
	return ((*[1 << 31]byte)(p))[0:n:n]
}

func freeByteSlice(p []byte) {
	C.free(unsafe.Pointer(&p[0]))
}

func main() {
	// 这是最大了 不知道为啥
	s := makeByteSlize(1 << 31)
	s[len(s)-1] = 255
	//print(s[len(s)-1])
	print(len(s))
	freeByteSlice(s)
}
