package main

import "fmt"

func main() {
	//变量的几种声明方式

	/**
	第一种，标准声明
	 */
	var a int = 10
	//自带换行输出标准输出控制台
	fmt.Println(a)

	/**
	第二种，自动推导类型
	 */
	b := 10
	//‘\n’windows的换行符
	fmt.Printf("b type is %T\n",b)
}
