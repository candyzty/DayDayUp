package main
import "fmt"
func main() {
    var a string = "Runoob"
    fmt.Println(a)

    var b, c int = 1, 2
	fmt.Println(b, c)
	// 声明一个变量并初始化
    var d = "RUNOOB"
    fmt.Println(d)

    // 没有初始化就为零值
    var e int
    fmt.Println(e)

    // bool 零值为 false
    var f bool
	fmt.Println(f)
	
	var g int
    var h float64
    var i bool
    var j string
    fmt.Printf("%v %v %v %q\n", g, h, i, j) // 0 0 false ""
}