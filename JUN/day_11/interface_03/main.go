package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//1. 声明一个Hero 结构体
type Hero struct {
	Name string
	Age  int
}

//2. 声明一个Hero 结构体切片类型

type HeroSlice []Hero

//3. 实现Interface

func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 这个方式就是决定你使用什么标准进行排序
// 按age 从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main() {
	intSlice := []int{0, -1, 10, 7, 90}
	// 使用sort 排序
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	// 测试我们自己写的结构体切片
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		// 将hero 添加到切片中
		heros = append(heros, hero)
	}
	//看看排序前的顺序
	for _, v := range heros {
		fmt.Println(v)
	}
	// 调用sort 包下面的方法
	sort.Sort(heros)
	fmt.Println("排序后的顺序")
	for _, v := range heros {
		fmt.Println(v)
	}
}
