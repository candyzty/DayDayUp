package main

import "fmt"

func main() {
	parmInit := make(map[string]string)
	fmt.Println(parmInit)
	fmt.Println("插入数据！")
	parmInit["name"] = "Tom"
	parmInit["age"] = "twenty five"
	fmt.Println(parmInit)
	for key, value := range parmInit {
		fmt.Println(key, "-->", value)
	}
	fmt.Println("修改参数")
	parmInit["name"] = "Jack"
	fmt.Println(parmInit)

	value, ok := parmInit["name"]

	fmt.Println("hahah")
	if ok == true {
		fmt.Println("parmInit[1] = ", value)
	} else {
		fmt.Println("key不存在")
	}
}
