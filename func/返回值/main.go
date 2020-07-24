package main

import "fmt"

func incr(a int) int {
	var b int
	defer func() {
		a++
		fmt.Println(a)
		b++
		fmt.Println(b)
	}()
	a++
	b = a
	return b
}

func incr1(a int) (b int) {

	defer func() {
		a++
		fmt.Println(a)
		b++
		fmt.Println(b)
	}()
	a++
	b = a
	return b
}
func main() {
	a := 0
	//c := incr(a)
	c := incr1(a)
	fmt.Println(c, a)
}
