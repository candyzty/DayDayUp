package main

import "fmt"

//1. 方法的语法
//type Employee struct {
//	name string
//	salary int64
//	currency string
//}
//
//func (e Employee)displaySalary()  {
//	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
//
//}
//func main() {
//	emp1 := Employee{
//		name:     "Jack",
//		salary:   5000,
//		currency: "$",
//	}
//	emp1.displaySalary()
//}

//2. 相同方法名
//type Rectangle struct {
//	width,height float64
//}
//type Circle struct {
//	radius float64
//}
//
//func (r Rectangle)area() float64  {
//	return r.width * r.height
//}
//
//func (c Circle)area()float64  {
//	return c.radius * c.radius * math.Pi
//}
//func main() {
//	r1 := Rectangle{12, 2}
//	r2 := Rectangle{9, 4}
//	c1 := Circle{10}
//	c2 := Circle{25}
//	fmt.Printf("Area of r1 is : %f\n", r1.area())
//	fmt.Printf("Area of r2 is : %f\n", r2.area())
//	fmt.Printf("Area of c1 is : %f\n", c1.area())
//	fmt.Printf("Area of c2 is : %f\n", c2.area())
//}

//3. 指针作为接收者
type Rectangle struct {
	width, height int64
}

//func (r *Rectangle)setVal()  {
func (r Rectangle) setVal() {
	r.height = 20
}
func main() {
	p := Rectangle{1, 2}
	s := p
	p.setVal()
	fmt.Println(p.height, s.height)
}
