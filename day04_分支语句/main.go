package main

import "fmt"

//1.1if语句
//func main() {
//	var a int = 10
//	if a < 20 {
//		fmt.Println("a 小于 20")
//	}
//	fmt.Printf("a 的值为 : %d\n", a)
//}
//1.2 if变体
//func main() {
//	if num := 10; num % 2 == 0 {
//		fmt.Println(num,"is even")
//	}else{
//		fmt.Println(num, "is odd")
//	}
//}

//2.1 switch
//func main() {
//	var grade string = "B"
//	switch grade {
//	case "A" :
//		fmt.Println("优秀！")
//	case "B", "C":
//		fmt.Println("良好")
//	case "F":
//		fmt.Println("不及格")
//	default:
//		fmt.Println("差")
//	}
//	fmt.Printf("你的等级是: %s\n", grade)
//}

//2.2 省略表达式
//func main() {
//	num := 75
//	switch  {
//	case num >= 0 && num <= 50:
//		fmt.Println("num is greater than 0 and less than 50")
//	case num >=51 && num <= 100:
//		fmt.Println("num is greater than 51 and less than 100")
//	case num >= 101:
//		fmt.Println("num is greater than 100")
//
//	}
//}

//2.3 fallthrough
//func main() {
//	switch x :=5; x {
//	default:
//		fmt.Println(x)
//	case 5:
//		x += 10
//		fmt.Println(x)
//		fallthrough
//	case 6:
//		x += 20
//		fmt.Println(x)
//	case 7:
//		x += 30
//		fmt.Println(x)
//
//	}
//}

//3. type switch
func main() {
	var x interface{} = 3.4
	switch i := x.(type) {
	case nil:
		fmt.Println("x 的类型是：%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")

	}
}
