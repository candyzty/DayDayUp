package main

//	创建 make(map[string]int)
//	获取元素 m[key]
//	key 不存在时，获得value类型的初始值
//	用value,ok:= m[key] 来判断是否存在key
//	用delete删除一个key
//	使用range遍历key，或者遍历key,value对
//	map是无序的
//	使用len获得map的个数
//	除了slice，map,function的内建类型都可以作为key
//	Struct类型不包含上述字段，也可作为key
import (
	"fmt"
)

func main() {
	// 第一种创建的方法
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	// 2、 定义一个空map,m2 == empty map
	m2 := make(map[string]int)

	// 3、 定义一个空map，m3 == nil
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	fmt.Println("遍历map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting values")
	coureName := m["course"]
	fmt.Println(coureName)

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
}
