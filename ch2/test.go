package main

//extern int go_qsort_compare(void* a, void* b);
import "C"

import (
	"fmt"
	"github.com/ljcnh/goAdPro/ch2/qsort"
	"unsafe"
)

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println(values)
	//qsort.Sort(unsafe.Pointer(&values[0]), len(values), int(unsafe.Sizeof(values[0])),
	//	func(a, b unsafe.Pointer) int {
	//		pa, pb := (*int32)(a), (*int32)(b)
	//		return int(*pa - *pb)
	//	},
	//)
	//qsort.Sort(unsafe.Pointer(&values[0]),
	//	len(values), int(unsafe.Sizeof(values[0])),
	//	qsort.CompareFunc(C.go_qsort_compare),
	//)
}
