package main

import "fmt"

//1. 函数的使用
//func main() {
//	var a int = 100
//	var b int = 200
//	var result int
//	result = max(a, b)
//	fmt.Printf("最大值是：%d \n", result)
//}
//
//func max(num1, num2 int) int {
//var result int
//if (num1 > num2){
//	result = num1
//}else {
//	result = num2
//}
//	return result
//}

//2. 值传递
//import "math"
//func main() {
//	getSquareRoot := func(x float64) float64 {
//		return math.Sqrt(x)
//	}
//	fmt.Printf("%.2f\n",getSquareRoot(9))
//	fmt.Println(getSquareRoot(9))
//}

//3.引用传递
//func main() {
//	x := 3
//	fmt.Println("x=", x)
//	x1 := add1(&x)
//	fmt.Println("x+1=", x1)
//	fmt.Println("x=", x)
//}
//func add1(x *int) int  {
//	*x += 1
//	return *x
//}

//4. 一个函数可以返回多个值
//func main() {
////	a, b := swap("Mike", "James")
////	fmt.Println(a, b)
////}
////func swap(x,y string) (string,string)  {
////	return y,x
////}

//5. 空白标识符
//func main() {
////	area,_ := rectProps(10.8, 5.6)
////	fmt.Printf("Area %f\n", area)
////}
////func rectProps(length,width float64) (float64,float64)  {
////	area := length * width
////	permimeter := (length + width) *2
////	return area,permimeter
////}

//6. 延迟函数
//func main() {
//	a := 1
//	b := 2
//	defer fmt.Println(b)
//	fmt.Println(a)
//}
//func main() {
//	nums := []int{4, 18, 12, 15}
//	largest(nums)
//}
//func finished()  {
//	fmt.Println("Finished")
//}
//func largest(ss []int)  {
//	defer finished()
//	fmt.Println("Started")
//	max := ss[0]
//	for _,x := range ss{
//		if x > max{
//			max = x
//		}
//	}
//	fmt.Println("Largest number in", ss, "is", max)
//}

//7. 延迟方法
//type person struct {
//	firstName string
//	lastName string
//}
//
//func (p1 person) fullName()  {
//	fmt.Printf("%s %s",p1.firstName, p1.lastName)
//}
//func main() {
//	p := person{
//		firstName: "John",
//		lastName:  "Smith",
//	}
//	defer p.fullName()
//	fmt.Printf("Welcome ")
//}

//8. 延迟参数
//func main() {
//	a := 5
//	defer printA(a)
//	a = 10
//	fmt.Println("value of a before defered function call", a)
//}
//func printA(a int)  {
//	fmt.Println("value of a in deffered function", a)
//}

//9.堆栈的推迟
func main() {
	name := "Naveen"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	//for _,v := range []rune(name){
	//	defer fmt.Printf("%c", v)
	//}
	for _, v := range []int32(name) {
		defer fmt.Printf("%c", v)
	}
}
