package main

import "runtime"

func main() {
	gosched()
}

func init() {
	//获取go-root目录
	println(runtime.GOROOT())
	println(runtime.GOOS)
	//获取逻辑cpu的数量
	println(runtime.NumCPU())
	//设置go程序执行的最大的：【1，256】
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	println(n)
}
func gosched() {
	go func() {
		for i := 0;i < 5 ;i++  {
			println("test...")
		}
	}()

	for i := 0; i< 4 ; i++  {
		//让出时间片，先让被的协议执行，执行完在执行此协程
		runtime.Gosched()
		println("main...")
	}
}
