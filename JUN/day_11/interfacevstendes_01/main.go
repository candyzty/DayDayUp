package main

import "fmt"

// Monkey 结构体

type Monkey struct {
	Name string
}

// 声明接口
type BirdAble interface {
	Flying()
}

type LittleMonKey struct {
	Monkey
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来回爬树")
}

func (sw *LittleMonKey) Flying() {
	fmt.Println(sw.Name, "通过学习过飞翔")
}

func (sw *LittleMonKey) Swimming() {
	fmt.Println(sw.Name, "通过学习可以进行游泳")
}

func main() {
	// 创建一个LittleMonKey实例
	monkey := LittleMonKey{
		Monkey{
			Name: "猴子",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
