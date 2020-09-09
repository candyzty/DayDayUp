package main

import "fmt"

func main() {
	xunhuan()
}

func xunhuan() {
	println("请输入")
	var scanln string
	fmt.Scanln(&scanln)

	if scanln != "exit" {
		printSB()
		xunhuan()
	}
}

func printSB() {
	fmt.Println("薛佳是傻逼")
}
