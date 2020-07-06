package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}
func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
