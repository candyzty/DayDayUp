package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	s1 := Student{
		1,
		"mike",
		'w',
		23,
		"hz",
	}
	fmt.Println(s1)
}
