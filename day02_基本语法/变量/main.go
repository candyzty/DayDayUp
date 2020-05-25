package 变量

import "fmt"
import "log"

var a = "Hello"
var b string = "World"
var c bool

func main() {
	d := 10
	d = 11
	println(a, b, c, d)
	x := 140
	fmt.Println(&x)
	x, y := 200, "abc"
	fmt.Println(&x, x)
	fmt.Println(y)
	log.Println(a, b, c, d)
}
