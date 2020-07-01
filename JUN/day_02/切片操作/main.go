package main

import "fmt"

func sliceCreate() {
	//var s  [] int
	//for i :=0;i<100;i++{
	//	printSlice(s)
	//	s = append(s,i)
	//}
	//fmt.Println(s)

	//  切片的三种方式
	// 1、s1 指定值来创建的
	// 2、s2 没有指定值，但是指定了长度为16
	// 3、s3 没有指定值，但是指定了长度为10，cap为32
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)
}
func sliceCopy() {
	fmt.Println("Slice copy")
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	//s3 := make([]int, 10, 32)
	copy(s2, s1)
	printSlice(s2)
}
func sliceDelete() {
	fmt.Printf("Delete elements from slice")
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	//s3 := make([]int, 10, 32)
	copy(s2, s1)
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2, len(s2), cap(s2))
}
func printSlice(s []int) {
	fmt.Printf(" value=%v,len=%d, cap=%d\n", s, len(s), cap(s))
}
func deleteFrontTail() {
	fmt.Println("--Pop  front and tail slice--")
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	copy(s2, s1)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println("Pop front slice ")
	s2 = s2[1:]
	fmt.Println(s2, len(s2), cap(s2))

	fmt.Println("Pop tail slice")
	s2 = s2[:len(s2)-1]
	fmt.Println(s2, len(s2), cap(s2))

}
func main() {
	//sliceCreate()
	sliceCopy()
	sliceDelete()
	deleteFrontTail()
}
