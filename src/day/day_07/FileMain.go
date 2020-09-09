package main

/**
IO操作
1.io包中提供了I/O原始操作的一系列接口。主要包装了一些已有的实现，如os包中的file操作类。并将这些抽象成为实用性的功能和一些其他相关的接口
2.由于这些接口和原始的操作以不同的实现包装了低级操作，客户不应假定他们对于并行执行是安全的。



几大重要的接口定义
1.Reader接口的定义，Read()方法用于读取数据。
type Reader interface {
     Read(p []byte) (n int, err error)
}
2.Writer 接口的定义，Write()方法用于写出数据。
type Writer interface {
     Write(p []byte) (n int, err error)
}

3. Seeker接口的定义，封装了基本的Seek方法。
Seek 设置下一次读写操作的指针位置，每次的读写操作都是从指针位置开始的
whence 的含义：
     case 0 ： 表示从数据的开头开始移动指针
     case 1 ： 表示从数据的当前指针位置开始移动指针
     case 2 ： 表示从数据的尾部开始移动指针
offset 是指指针移动的偏移量
返回移动后的指针位置和移动过程中遇到的任何错误
*/
func main() {

}
