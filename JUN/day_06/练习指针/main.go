package main

import (
	"fmt"
	"log"
)

func typePointer() {
	astr := "hello"
	aint := 1
	abool := false
	arune := 'a'
	afloat := 1.2

	fmt.Printf("astr 指针类型是：%T\n", &astr)
	fmt.Printf("aint 指针类型是：%T\n", &aint)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("arune 指针类型是：%T\n", &arune)
	fmt.Printf("afloat 指针类型是：%T\n", &afloat)
}

func mytest(ptr *int) {
	fmt.Println(*ptr)
}

func nilPointer() {
	a := 25
	var b *int
	if b == nil {
		log.Println(b)
		b = &a
		log.Println(b)
	}
}
func modify(sls []int) {
	sls[0] = 90
}

func modifr(arr *[3]int) {
	arr[0] = 90
}
func main() {
	typePointer()
	nilPointer()

	a := [3]int{89, 90, 91}
	modify(a[:])
	log.Println(a)

	log.Println("使用指针")
	modifr(&a)
	fmt.Println(a)

}
