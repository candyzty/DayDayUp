package main

import (
	"fmt"
	"os"
)

var parmInit = make(map[string]string)

func forStatic() {
	var (
		y        string
		userName string
		userKey  string
		number   int
	)
	fmt.Println("玩游戏的次数为:")
	fmt.Scanln(&number)
	for i := 1; i < number+1; i++ {
		fmt.Printf("第%d次输入的值\n", i)
		createUser(userName, userKey)
	}
	fmt.Println(parmInit)
	fmt.Println("是否继续玩下去(y/n):")
	fmt.Scanln(&y)
	if y == "y" {
		fmt.Println("继续进行:")
		forStatic()
	}
}

func createUser(value, key string) {
	fmt.Println("第一个参数输入的是:value,第二个参数输入的是:key")
	fmt.Scanln(&value, &key)
	parmInit[value] = key
	fmt.Println(parmInit)
}

func getMapInfo() {
	getValue := ""
	fmt.Println("查询key是否存在！")
	fmt.Scanln(&getValue)
	kayName, ok := parmInit[getValue]

	if ok == true {
		fmt.Println("parmInit = ", kayName)
	} else {
		fmt.Println("key不存在")
	}
}

func updateMapInfo() {
	updateKey := ""
	updateValue := ""
	fmt.Println("更新map参数:")
	fmt.Println("请输入更新key， value:")
	fmt.Scanln(&updateKey)
	fmt.Scanln(&updateValue)
	keyName, ok := parmInit[updateKey]
	if ok == true {
		fmt.Println("parmInit = ", keyName)
	} else {
		fmt.Println("key不存在")
		os.Exit(0)
	}
	parmInit[updateKey] = updateValue
	fmt.Println(parmInit)
}
func delectMapInfo() {
	delKey := ""
	fmt.Println("删除map值:")
	fmt.Scanln(&delKey)
	keyName, ok := parmInit[delKey]
	if ok == true {
		fmt.Println("parmInit = ", keyName)
	} else {
		fmt.Println("key不存在")
		os.Exit(0)
	}
	delete(parmInit, delKey)
	fmt.Println(parmInit)
	//delectMapInfo()
}

func buildParameters() {
	parameters := ""
	airs := `
	GET)   getMapInfo
	PUT)   updateMapInfo
	DEL)   delectMapInfo 
    HELP)  Select the parameters of the build
    EXIT)  Exit is   0 
	`
	fmt.Println(airs)
	fmt.Scanln(&parameters)
	switch parameters {
	case "GET":
		getMapInfo()
		buildParameters()
	case "PUT":
		updateMapInfo()
		buildParameters()
	case "DEL":
		delectMapInfo()
		buildParameters()
	case "EXIT":
		os.Exit(0)
	default:
		fmt.Println("参数错误")
	}
}
func main() {
	var x string
	enter := "yes"
	fmt.Println("请问是否开始答题(yes/no):")
	fmt.Scanln(&x) //读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
	if x == enter {
		fmt.Println("恭喜你,咱们开始了")
		forStatic()
		buildParameters()
	} else {
		fmt.Println("游戏退出！")
		os.Exit(0)
	}
}
