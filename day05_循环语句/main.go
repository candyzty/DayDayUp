package main

import "fmt"

//1.1 for 语句
//func main() {
//	for i := 1; i <= 10; i++{
//		fmt.Printf("%d\n", i)
//	}
//}
//1.2 for 变体
//func main() {
//	var b int = 15
//	var a int
//	numbers := [6]int{1, 2, 3, 5}
//	//打印0-9
//	for a := 0; a < 10; a++ {
//		fmt.Printf("a的值为：%d\n", a)
//	}
//	//打印1-15
//	for a < b {
//		a++
//		fmt.Printf("a 的值为： %d\n", a)
//	}
//	for i,x :=range numbers{
//		fmt.Printf("第%d位的值为：%d\n", i, x)
//	}
//}

//2.1break
//func main() {
//	//打印1-10，到5时退出
//	for i := 1; i <= 10; i++ {
//		if i > 5 {
//			break
//		}
//		fmt.Printf("%d ", i)
//	}
//	fmt.Printf("\n line alter for loop")
//}

//2.2 continue
//func main() {
//	//打印10以内奇数
//	for i := 0;i <= 10; i++ {
//		if i % 2 == 0 {
//			continue
//		}
//		fmt.Printf("%d ", i)
//	}
//}

//2.3 goto
//打印10-20，跳过15
func main() {
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("%d ", a)
		a++
	}
}
