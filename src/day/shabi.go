package main

import (
	"day_01"
	"fmt"
)

func main() {
	xunhuan()
}

func xunhuan()  {
	day_01.Day_01_test()
	println("请输入")
	var scanln string
	fmt.Scanln(&scanln)

	if scanln != "exit" {
		printSB()
		xunhuan()
	}
}

func printSB() {
	fmt.Println("张杰是傻逼")
}
