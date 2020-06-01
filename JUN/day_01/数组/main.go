package main

import "fmt"

// 定义数组的几种方法，数量写在数组的前面
// 数组是值类型
// 调用func f(arr [10]int) 会拷贝数组
// 在go 语言中一般不直接使用数组

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int // 代表5个int
	arr2 := [5]int{1, 2, 3, 4}
	arr3 := [...]int{5, 6, 7, 8, 5}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println("------------")
	// 1、数组的遍历
	//for i :=0;i<len(arr3);i ++{
	//	fmt.Println(i)
	//}
	// 2、数组的遍历
	//for b,v := range arr3{
	//	fmt.Println(b,v)
	//}
	printArray(arr3)
	printArray(arr1)
}
