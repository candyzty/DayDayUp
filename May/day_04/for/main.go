package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
		//fmt.Println(sum)
	}
	//fmt.Println(sum)

	a := 15
	b := 10
	for a >= b {
		b++
		fmt.Println(b)
	}
	d := []int{20, 43, 54}
	for _, v := range d {
		fmt.Println(v)
	}
}
