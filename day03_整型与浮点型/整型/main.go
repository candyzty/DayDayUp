package main

import "fmt"

////使用三种进制表示12
//func main() {
//  var num01 int64 = 0b1100
//  var num02 int64 = 0o14
//  var num03 int64 = 0xC
//  fmt.Printf("2进制数 %b 表示的是：%d \n",num01,num01)
//  fmt.Printf("8进制数 %o 表示的是：%d \n",num02,num02)
//  fmt.Printf("16进制数 %x 表示的是：%d \n",num03,num03)
//}

var myfloat01 float32 = 100000182
var myfloat02 float32 = 100000187

func main() {
	fmt.Println("myfloat: ", myfloat01)
	fmt.Println("myfloat: ", myfloat01+5)
	fmt.Println(myfloat02 == myfloat01+5)
}
