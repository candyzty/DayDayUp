package main

import "log"

// 占位符 https://blog.csdn.net/HHTNAN/article/details/78607180
// 布尔 string
// int8 有符号 8 位整型 (-128 到 127) 长度：8bit
//
// int16 有符号 16 位整型 (-32768 到 32767)
//
// int32 有符号 32 位整型 (-2147483648 到 2147483647)
//
// int64 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
//
// uint8 无符号 8 位整型 (0 到 255) 8位都用于表示数值：
//
// uint16 无符号 16 位整型 (0 到 65535)
//
// uint32 无符号 32 位整型 (0 到 4294967295)
//
// uint64 无符号 64 位整型 (0 到 18446744073709551615)
const c = 66
const f int64 = 66

func main() {
	var a bool = false
	var b int8 = 61
	log.Println(a, b)
	log.Println(c + b) // c没有显示声明类型, 隐式转型为了int8
	var d int8 = 62    // 越界  127
	log.Println(c + d)
	log.Println(f + d) // 此处会报错
}

// 复合类型
// 1、指针类型（Pointer） 2、数组类型 3、结构化类型(struct) 4、Channel 类型 5、函数类型 6、切片类型 7、接口类型（interface） 8、Map 类型
// 优先级	运算符
// 7	~ ! ++ --
// 6	* / % << >> & &^
// 5	+ - ^
// 4	== != < <= >= >
// 3	<-
// 2	&&
// 1	||
