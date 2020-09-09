package main

import (
	"fmt"
	"os"
)

/**
文件的操作
file类在os包下封装了底层的文件描述符和相关信息，同时封装了Read和Write的实现。
*/
func main() {
	var path string
	fmt.Scanln(&path)
	FileStatTest(path)
	FileModeTest(new(os.FileMode))
}

/**
文件信息接口  FileInfo
type FileInfo interface {
   Name() string   //文件名称
   Size() int64    //文件大小
   Mode() FileMode  //文件权限 --rwx--rwx--rwx
   ModeTime() time.time  //文件修改时间
   IsDir() bool     //是否是文件夹
   Sys() interface{}   //基于数据源接口（can return nil）
}
*/

func FileModeTest(s *os.FileMode) {
	var a uint32 = 214
	//s = a
	fmt.Print(a)
}

/**
Stat():方法
返回描述文件f的FileInfo类型值。如果出错，错误底层类型是*PathError
主要是文件的属性，name,size,mode,updateTime
*/
func FileStatTest(path string) {
	fileInfo, err := os.Stat(path)

	if err != nil {
		fmt.Println("err: ", err)
		return

	}

	fmt.Printf("%T\n", fileInfo)
	//文件名
	fmt.Println("文件名字：", fileInfo.Name())
	//文件大小
	fmt.Println("文件大小：", fileInfo.Size())
	//是否是目录
	fmt.Println("是否是目录：", fileInfo.IsDir()) //IsDirectory
	//修改时间
	fmt.Println("文件最后修改时间：", fileInfo.ModTime())
	//权限
	fmt.Println("文件的权限：", fileInfo.Mode()) //-rw-
}

/**
文件操作：mkdir，close，open,remove
*/
func fileOperation(path []string) {
	//打开文件：让当前的程序，和指定的文件之间建立一个连接
	fileOperation, err := os.Open(path[0])
	if err != nil {
		fmt.Println("err:", err)
	}
	fileOperation.Stat()
	//创建文件夹
	mode := os.FileMode(777)
	os.Mkdir("aa", mode)

}
func mkdirTest() {

}
