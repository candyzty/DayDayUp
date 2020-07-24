package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
}
func swap1(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	a, b := 1, 2
	//swap(a,b)
	swap1(&a, &b)
	fmt.Println(a, b)
}
