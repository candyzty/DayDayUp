package main

import "fmt"

//1.数组
//func main() {
//	var n = [10]int{}
//	//0-9-->100-109
//	for i :=0; i < 10; i ++ {
//		n[i] = i + 100
//	}
//	for j := 0; j < 10; j ++ {
//		fmt.Printf("Element[%d] = %d \n", j, n[j])
//	}
//}
//2. 数组长度
//func main() {
//	a := [...]float64{57.7, 33.2, 21, 78}
//	fmt.Println("length of a is", len(a))
//}
//3.使用索引遍历数组
//func main() {
//	a := [...]float64{57.7, 33.2, 21, 78}
//	for i := 0; i < len(a); i++ {
//		fmt.Printf("%d th element of a is %.2f \n", i, a[i])
//	}
//}
//4. 使用range遍历数组
//func main() {
//	a := [...]float64{57.7, 33.2, 21, 78}
//	//sum := 0.0
//	sum := float64(0)
//	for i, x := range a {
//		fmt.Printf("%d th element of a is %.2f \n", i, x)
//		sum += x
//	}
//	fmt.Println("\n sum of all elements of a", sum)
//}
//5. 数组是值类型，非引用类型
//func main() {
//	a := [...]string{"USA", "China", "India"}
//	b := a
//	b[0] = "Singapore"
//	fmt.Println("a is ", a)
//	fmt.Println("b is ", b)
//}
//6. 数组不能被调整大小
//func main() {
//	a := [3]int{1, 2, 3}
//	var b [5]int
//	b = a
//}
