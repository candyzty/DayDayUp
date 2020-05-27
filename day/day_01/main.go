package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	day_05()
}

func day_01(){
	//变量的几种声明方式

	/**
	第一种，标准声明
	*/
	var a int = 10
	//自带换行输出标准输出控制台
	fmt.Println(a)

	/**
	第二种，自动推导类型
	*/
	b := 10
	//‘\n’windows的换行符
	fmt.Printf("b type is %T\n",b)
}

func day_02()  {
	//常量的定义
	const a int8 = 127
	const c  = "v"
	fmt.Printf("c type is %T\n",c)

	const (
		b = 10
		/**
		单引号和双引号区分string还是byte<uint8>
		自动类型推导
		:=将'c'默认为uint16
		 */
		e = 'c'
		r = "c"
	)
	fmt.Printf("e type : %T,r type is : %T",e,r)

}
/**
控制台输入输出
 */
func day_03(){
	a := 10
	b := "10"
	/**
	普通打印
	 */
	fmt.Print(a)
	fmt.Print(b)
	/**
	打印后换行
	 */
	fmt.Println(a)
	fmt.Println(b)
	/**
	格式化打印。使用占位符
	常用占位符
	%v  原样输出
	%d  10进制格式
	%s  字符串格式
	%t  bool类型
	%f  浮点类型float系列
	%c  字符打印byte。'c'
	%c  打印地址
	%T  数据类型输出
	 */
	fmt.Printf("a : %d type is %T\nb : %s type is %T\n",a,a,b,b)


	/**
	控制台输入
	1.fmt包下的scan系列。与print系列相似
	 */
	var s int
	var r string
	fmt.Scan(&s)
	fmt.Scan(&r)
	fmt.Println(r)
	fmt.Println(s)
	fmt.Scanln(&s)
	fmt.Scanln(&r)
	fmt.Println(s)
	fmt.Println(r)
	fmt.Scanf("%d,%s",&s,&r)
	fmt.Println(s,r)

	/**
	2.通过IO从控制台输入
	 */
	fmt.Print("请输入：")
	reader := bufio.NewReader(os.Stdin)
	s1, err := reader.ReadString('a')
	fmt.Printf("读取的数据：%s,err:%T",s1,err)
}

/**
GO的流程控制
顺序结构，选择结构，循环结构
 */
func day_04(){
	/**
	选择结构
	if，else，switch，select
	 */
	a,b := 3,4
	if a > b {
		fmt.Println("a大于b")
	}
	if s, r := 10, 11; s < r {
		fmt.Println("s小于r")
	}

	switch a {
	case 1:
		fmt.Println("选择1");
	case 3:
		fmt.Println("选择3");
		fallthrough
	case 4:
		fmt.Println("选择4");
		fallthrough
	case 5:
		fmt.Println("选择5");
	default:
		fmt.Println("hello！ 借宿一");
	}
	fmt.Println("结束了！！！");
}

func day_05(){
	for i := 0; i<10 ; i++ {
		fmt.Println(i);
	}
}