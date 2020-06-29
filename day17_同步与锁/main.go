package main

import (
	"fmt"
	"sync"
	"time"
)

//例子一
//func main() {
//	wg := sync.WaitGroup{}
//	var mutex sync.Mutex
//	fmt.Println("Locking (Go)")
//	mutex.Lock()
//	fmt.Println("Locked (Go)")
//	wg.Add(3)
//
//	for i:=1; i<4; i++{
//		go func(i int) {
//			fmt.Printf("Locking (G%d)\n", i)
//			mutex.Lock()
//			fmt.Printf("locked (G%d)\n", i)
//
//			time.Sleep(time.Second * 2)
//			mutex.Unlock()
//			fmt.Printf("unlocked (G%d)\n", i)
//			wg.Done()
//		}(i)
//	}
//	time.Sleep(time.Second * 5)
//	fmt.Println("ready unlock (Go)")
//	mutex.Unlock()
//	fmt.Println("unlocked (Go)")
//
//	wg.Wait()
//}

//例子二
//type Book struct {
//	BookName string
//	L *sync.Mutex
//}
//
//func (bk *Book) SetName(wg *sync.WaitGroup, name string)  {
//	defer func() {
//		fmt.Println("Unlock set name:", name)
//		bk.L.Unlock()
//		wg.Done()
//	}()
//	bk.L.Lock()
//	fmt.Println("Lock set name:", name)
//	time.Sleep(1 * time.Second)
//	bk.BookName = name
//}
//func main() {
//	bk := Book{}
//	bk.L = new(sync.Mutex)
//	wg := &sync.WaitGroup{}
//	books := []string{"《三国演义》", "《道德经》", "《西游记》"}
//	for _, book := range books{
//		wg.Add(1)
//		go bk.SetName(wg, book)
//	}
//	wg.Wait()
//}

//例子三
var m *sync.RWMutex

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	var rwMutex sync.RWMutex
	Data := 0
	for i := 0; i < 10; i++ {
		go func(t int) {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			fmt.Printf("Read data: %v\n", Data)
			wg.Done()
			time.Sleep(2 * time.Second)
			// 这句代码第一次运行后，读解锁。
			// 循环到第二个时，读锁定后，这个goroutine就没有阻塞，同时读成功。
		}(i)

		go func(t int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			Data += t
			fmt.Printf("Write Data: %v %d \n", Data, t)
			wg.Done()

			// 这句代码让写锁的效果显示出来，写锁定下是需要解锁后才能写的。
			time.Sleep(2 * time.Second)
		}(i)
	}
	time.Sleep(5 * time.Second)
	wg.Wait()
}
