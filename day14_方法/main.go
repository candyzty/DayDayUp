package main

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
//type Rectangle struct {
//	width, height int64
//}
//
////func (r *Rectangle)setVal()  {
//func (r Rectangle) setVal() {
//	r.height = 20
//}
//func main() {
//	p := Rectangle{1, 2}
//	s := p
//	p.setVal()
//	fmt.Println(p.height, s.height)
//}

//接口
import (
	"fmt"
	"strconv"
)

type Good interface {
	settleAccount() int
	orderInfo() string
}
type Phone struct {
	name     string
	quantity int
	price    int
}

func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}
func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}
func main() {
	iPhone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}
	goods := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}
