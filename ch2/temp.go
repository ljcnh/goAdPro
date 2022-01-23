package main

import "C"

func main() {
	//C.printString("hello")
}

//v, _ := C.noreturn()
//fmt.Printf("%#v", v)
//v := C.add(1, 1)
//v = C.div(6, 3)
//fmt.Println(v)
//
//v1, err1 := C.div(1, 0)
//fmt.Println(v1, err1)
//// 通过 reflect.SliceHeader 转换
//var arr0 []byte
//var arr0Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&arr0))
//arr0Hdr.Data = uintptr(unsafe.Pointer(&C.arr[0]))
//arr0Hdr.Len = 10
//arr0Hdr.Cap = 10
//
//// 通过切片语法转换
//arr1 := (*[31]byte)(unsafe.Pointer(&C.arr[0]))[:10:10]
//fmt.Println(arr1)
//
//var s0 string
//var s0Hdr = (*reflect.StringHeader)(unsafe.Pointer(&s0))
//s0Hdr.Data = uintptr(unsafe.Pointer(C.s))
//s0Hdr.Len = int(C.strlen(C.s))
//
//sLen := int(C.strlen(C.s))
//s1 := string((*[31]byte)(unsafe.Pointer(&C.s[0]))[:sLen:sLen])
//fmt.Println(s1)
///*
//#include <stdint.h>
//
//union B1 {
//    int i;
//    float f;
//};
//
//union B2 {
//    int8_t i8;
//    int64_t i64;
//};
//
//enum C {
//    ONE,
//    TWO,
//};
//*/
//import "C"
//import "fmt"
//
//func main() {
//	var c C.enum_C = C.TWO
//	fmt.Println(c)
//	fmt.Println(C.ONE)
//	fmt.Println(C.TWO)
//	//var b C.union_B1
//	//fmt.Println("b.i:", *(*C.int)(unsafe.Pointer(&b)))
//	//fmt.Println("b.f:", *(*C.float)(unsafe.Pointer(&b)))
//
//	//var b1 C.union_B1;
//	//fmt.Printf("%T\n", b1) // [4]uint8
//	//var b2 C.union_B2;
//	//fmt.Printf("%T\n", b2) // [8]uint8
//}
