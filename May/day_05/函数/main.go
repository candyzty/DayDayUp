package main

import (
	"fmt"
	"math"
)

func operate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic(fmt.Sprintf("Not support operate:%s", op))
	}
}

func swap(a, b int) (int, int) {
	return b, a
}

func compute(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func posInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

type Greeting func(name string) string

func say(g Greeting, word string) {
	fmt.Println(g(word))
}
func english(name string) string {
	return "hello," + name
}

func french(name string) string {
	return "Bonjour," + name
}

func sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s = nums[i]
		//fmt.Println(s)
	}
	//fmt.Println(s)
	return s
}

func main() {
	a := 10
	b := 20
	//c := operate(10,5,"/")
	//fmt.Println(c)
	fmt.Println(swap(a, b))
	//file ,err := os.Open("main.go")
	//if err != nil{
	//	fmt.Println("youwu",err)
	//}else {
	//	fmt.Println("文件名称是:", file)
	//}
	fmt.Println(compute(posInt, 3, 4))
	say(english, "world")
	say(french, "Her name is Tom!")
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
