package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1 %v  len(s1) %d cap(s1) %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2 %v  len(s2) %d cap(s2) %d \n", s2, len(s2), cap(s2))

	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3 //  s3 和s4  都指向了同一个底层数组
	fmt.Println(s3, s4)
	s3[1] = 340
	fmt.Println(s3, s4)

	//  切片的遍历
	// 索引遍历
	//log.Println("索引 遍历")
	for i := 0; i < len(s3); i++ {
		//fmt.Print("索引 遍历")
		fmt.Print(s3[i])
	}
	// range 遍历
	//log.Println("range 遍历")
	for i, v := range s3 {
		//log.Println("range 遍历")
		fmt.Println(i, v)
	}
}
