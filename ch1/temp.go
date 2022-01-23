// 随意写的代码

package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"
	"unicode/utf8"
	"unsafe"
)

func Add2Slice(s *[]int, t int) {
	(*s)[0]++
	*s = append(*s, t)
	(*s)[0]++
}

func AddSlice(s []int, t int) {
	s[0]++
	s = append(s, t)
	s[0]++
	fmt.Println(s)
}

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}

func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

//func f(x int) *int {
//	return &x
//}

func g() int {
	var x = new(int)
	return *x
}

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

//var a string
//var done bool
//
//func setup() {
//	a = "hello, world"
//	done = true
//}

var a string
var ch = make(chan struct{})

func f() {
	print(a)
	ch <- struct{}{}
}
func hello() {
	a = "hello, world"
	go f()
}

func main() {
	defer func() {
		if r := recover(); r != nil {

		}
		// 虽然总是返回nil, 但是可以恢复异常状态
	}()

	// 警告: 用`nil`为参数抛出异常
	panic(nil)
	//for i:=0;i<10000;i++{
	//	fmt.Println(<-ch)
	//}
	//hello()
	//<- ch
	//done := make(chan int)
	//
	//go func(){
	//	println("你好, 世界")
	//	done <- 1
	//}()
	//
	//<-done
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//go worker(&wg)
	//go worker(&wg)
	//
	//wg.Wait()
	//
	//fmt.Println(total)
	//fmt.Fprintln(os.Stdout, UpperString("hello, world"))
	//a := []int{0, 1, 2, 3}
	//b := copy(a[1:], a[1+1:])
	//fmt.Println(a)
	//fmt.Println(b)
	//a = a[:1+copy(a[1:], a[1+1:])]
	//fmt.Println(a)
	//a := []int{1, 2, 3}
	//a = append(a[:0], a[1:]...) // 删除开头1个元素
	//a = []int{1, 2, 3}
	//a = append(a[:0], a[2:]...) // 删除开头N个元素
	//fmt.Println(a)
	//a := []int{0, 1, 2, 3}
	//Add2Slice(&a, 4) // [2 1 2 3 4]
	//
	//b := make([]int,4,10)  //[0 0 0 0]
	//AddSlice(b, 4) //  [2 0 0 0 4]
	//fmt.Println(b) // [2 0 0 0]
	//
	//c := []int{0, 1, 2, 3}
	//AddSlice(c, 4) // [2 1 2 3 4]
	//fmt.Println(c)   //  [1 1 2 3]
}

// for range 的模拟
func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

// []byte(s) 转换模拟
func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}

// string(bytes)转换模拟
func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)
	return p
}

// []rune(s)转换模拟 : 字符串到[]rune的转换必然会导致重新分配[]rune内存空间
func str2runes(s string) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

// string(runes)转换模拟  []rune到字符串的转换也必然会导致重新构造字符串
func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}

func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}

func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}

//func FindPhoneNumber(filename string) []byte {
//	b, _ := ioutil.ReadFile(filename)
//	return regexp.MustCompile("[0-9]+").Find(b)
//}

// 应该这么写

func FindPhoneNumber(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	return append([]byte{}, b...)
}

func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	// 以int方式给float64排序
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr
	// 以int方式给float64排序
	sort.Ints(c)
}
