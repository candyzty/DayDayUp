package main

import "fmt"

//自定义类型和类型别名示例
//1.自定义类型
//MyInt基于int类型的自定义类型
//type MyInt int
//func main() {
//	var i MyInt
//	fmt.Printf("type:%T value:%v\n", i, i)
//}
//2.类型别名
//type NewInt = int
//
//func main()  {
//	var x NewInt
//		fmt.Printf("type:%T value:%v\n", x, x)
//}

//定义结构体
//type person struct {
//	name string
//	age int8
//	city string
//}
type person struct {
	name, city string
	age        int8
}

func main() {
	var p person
	p.name = "Jack"
	p.city = "北京"
	p.age = 18
	fmt.Printf("p=%#v\n", p)
	fmt.Println(p.name)

	//匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.name = "小王子"
	user.married = false
	fmt.Printf("%#v\n", user)

}
