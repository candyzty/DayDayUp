package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())

	// 时间戳
	time1 := now.Unix()
	fmt.Println(time1)
	timenano := now.UnixNano()
	fmt.Println(timenano)
	// time.Unix()
	ret := time.Unix(1592985349, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	fmt.Println(ret.Month())

	//时间格式化
	// 二十四小时制
	fmt.Println(now.Format("[2006-01-02 15:04:05.000]  "))

	// 十二小时制
	fmt.Println(now.Format("[2006-01-02 03:04:05.000]  "))

	//	 时间间隔
	nex, err := time.Parse("2006-01-02", "2019-08-04")
	if err != nil {
		fmt.Println("时间错误")
		return
	}

	a := nex.Sub(now)
	fmt.Println(a)
}
