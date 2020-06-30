package main

import "fmt"

var a int

func test() {
	fmt.Println(a)
}
func main() {
	fmt.Println(a)
	a++
	fmt.Println(a)
	test()
}
