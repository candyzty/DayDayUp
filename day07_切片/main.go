package main

import "fmt"

//1.声明
//func main() {
//	a := [5]int{1, 2, 3, 4, 5}
//	//索引1到3
//	var b []int = a[1:4]
//	fmt.Println(b)
//}

//2.修改
//func main() {
//	darr := [...]int{1, 2, 3, 4, 5, 6}
//	dslice := darr[1:5]
//	fmt.Println("array before", darr)
//	for i := range dslice{
//		dslice[i] = dslice[i] + 10
//	}
//	fmt.Println("array after", darr)
//}
//array before [1 2 3 4 5 6]
//array after [1 12 13 14 15 6]

//3.多次修改
//func main() {
//	numa := [3]int{1, 2, 3}
//	nums1 := numa[:]
//	nums2 := numa[:]
//	fmt.Println("array before chage 1", numa)
//	nums1[0] = 100
//	fmt.Println("array after modification to slice nums1", numa)
//	nums2[1] = 200
//	fmt.Println("array after modification to slice nums2", numa)
//}

//4. len和cap
//func main() {
//	var numbers = make([]int, 3, 5)
//	printSlice(numbers)
//}
//func printSlice(x []int)  {
//	fmt.Printf("len is %d, slice is %d, slice is %v\n", len(x), cap(x), x)
//}

//5.空切片
//func main() {
//	//numbers := [] int{}
//	//numbers := make([]int, 0, 0)
//	var numbers []int
//	printSlice(numbers)
//	if(numbers == nil){
//		fmt.Println("切片是空的")
//	}
//}
//func printSlice(x []int)  {
//	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
//}

//6. append和copy
func main() {
	var numbers []int
	printSlice(numbers)
	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4, 5, 6, 7, 8)
	printSlice(numbers)
	//numbers = append(numbers, 2)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 3)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 4)
	//printSlice(numbers)

	//numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	//copy(numbers1, numbers)
	//printSlice(numbers1)
	//
	//numbers1 = append(numbers1, 5,6,7,8,9,10,11)
	//printSlice(numbers1)
	//
	//numbers1 = append(numbers1, 12)
	//printSlice(numbers1)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}
