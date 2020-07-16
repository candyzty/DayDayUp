package main

import "fmt"

func InitSlice() []string {
	var s1 []string
	return s1
}

// 调用append函数必须用原来的切片变量接收返回值
func AppendSlice(s []string, parameter []string) []string {
	//aa := append(s,"北京","上海","广州","深圳")
	s = append(s, parameter...)
	fmt.Println(len(s))
	return s
}

// 获取切片的
func SearchSlice(s []string, index int) (IndexName string) {
	//var err  error()
	IndexName = s[index]
	IndexName = fmt.Sprintf("获取到的切片为：%s，索引为:%d, 索引位置的参数为: %s\n", s, index, IndexName)
	return IndexName
}

func DeleteSlice(s []string, value string) []string {
	var s1 []int
	for i, v := range s {
		//fmt.Println(i,v)
		if v == value {
			s1 = append(s1, i)
		}
	}
	s = append(s[:1], s[2:]...)
	//fmt.Println(s)
	return s
}
func main() {
	var id int
	parameter := []string{"北京", "杭州", "深圳", "广州", "上海", "成都", "武汉", "杭州"}
	//for i,v := range parameter{
	//	fmt.Println(i,string(v))
	//}
	Slice := AppendSlice(InitSlice(), parameter)
	fmt.Println(Slice)
	fmt.Print("请输入获取索引来进行获取参数:")
	fmt.Scan(&id)
	GetIndex := SearchSlice(Slice, id)
	fmt.Println(GetIndex)
	fmt.Println(DeleteSlice(Slice, "杭州"))
}
