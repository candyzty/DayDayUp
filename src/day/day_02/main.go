package main

import (
    "fmt"
)
/**
第二天
数组
切片
map
string
 */
func main()  {
	day_04()
}

/**
数组的学习
 */
func day_01(){
    /**
    数组的学习分为几步骤
    1.初始化
    2.内置属性
    3.数组的下标访问
     */
    /*
    1.初始化
     */
    //初始化数组大小
    var name [10] int8 //默认为10个0
    fmt.Println("name:",name);
    var name1 = [5] int8{1,2,3,4,5}
    fmt.Println("name1 :",name1);
    //忽略[]设置数组的大小。自动默认值
    var name2 = []int8{1,2}
    fmt.Println("name2 :",name2);
    name3 := [...]byte{'a','b','b'}
    fmt.Println("name3 :",name3);
    //数组初始化的指定索引值
    name4 := [...]int8{1:4, 3:5, 7:8} //下标索引1的值为4，下标为3的值5.。
    fmt.Println("name4 :", name4);

    /**
    2.数组的内置属性
     */
    //<1. 数组的长度
    fmt.Println("name length is : ", len(name));
    //<2. 遍历数组
    for i := 0;i < len(name1) ; i++ {  //普通方法
        fmt.Println(name1[i]);
    }
    for i,v := range name2 {
        fmt.Printf("下标为%d的值为%v\n",i,v);
    }



}

/**
切片的学习
1.切片即原数组的传递引用
 */
func day_02() {
    /**
    从以下几个方面入手
    1.初始化
    2.内置属性
    3.强大功能
     */
    /**
    1.初始化
     */
    //初始化切片使用make([]type, len, capacity)函数初始化
    //type ：即数组元素类型 len:初始化长度
    //capacity：从文档中看似是一个数组长度最大值,cap会自动生长
    var name []int8 = make([]int8, 8)
    name1 := make([]int8, 2, 5)
    fmt.Println(name);
    fmt.Println(name1);
    //从已经初始化好的数组上分出一个切片
    var template = [8]int8{1,2,3,4,5,6,7,8}
    s := template[2:5]
    fmt.Println(s);

    /**
    修改切片
     */
    //在原数组的分片上对数据进行操作。原数组的值也会被处理
    //即分片就是原数组的传递引用
    for i,j := range s{
        s[i] = j+1
    }
    fmt.Println(template);

    /**
    2.内置属性
     */
    //len()  获取切片的长度
    //cap()  获取切片最长可以到达的长度
    fmt.Printf("len = %d cap = %d slice=%v\n",len(name1),cap(name1),name1)
    //append()  向切片里面追加一个或多个元素。返回一个数据类型相同的切片。
    //copy()    从原切片的数组中复制到目标切片中。注意TODO:是复制不是引用
    d := append(name1, 3, 4, 5, 6)
    d[0] = 1
    fmt.Printf("value : {%v},length:{%d},cap:{%d}\n",name1,len(name1),cap(name1));
    //fmt.Println(append(name1, 3, 4, 5, 6)+"length:"+len(name1)+"cap ："+cap(name1))
    printSlice(d)
    //只复制对应的索引对应的值的数据。
    copy(name1,d)
    printSlice(name1)
}

/**
map的学习。map也是Go的内置类型
map 引用类型。一个变化，另一个会有相应的变化
 */
func day_03() {
    /**
    1.map的初始化
    默认声明，默认是nil
    var name map[key_dataType]value_dataType
    使用make函数
    name = make(map[key_dataType]value_dataType)
     */
    //只是声明map并不对其进行初始化。map为nil不能添加元素
    var name1 map[int8]string
    fmt.Println(name1,len(name1));
    name2 := make(map[int8]string)
    print(name2,len(name2))

    name3 := map[string]float32{"c":5,"Go":7.5,"java":8}
    for count := range name3 {
        fmt.Printf("%s : %f\n",count,name3[count])
    }
    //增删改查
    name4 := map[int8]string{1:"cc"}
    delete(name4,1)
    name4[2] = "cand"
    fmt.Println(name4[2]);
    name4[2] = "candy"
    fmt.Println(name4[2]);
    delete(name4,2)
    _,ok := name4[2]
    fmt.Printf("%t\n",ok)

    /**
    2.内置函数
     */
    //1.查看元素在集合中是否存在
    cas,ok := name3["c"]
    if ok {
        fmt.Println(cas);
    }
    //2.获取map的长度
    fmt.Printf("%d\n",len(name4))

}

/**
string的学习
Go的字符串是一个字节的切片。可以通过将其内容封装在""中来创建。
Go的字符串是Unicode兼容的，并且是UTF-8编码
3.工具类1 strings
    类似java strings。封装操作string的函数
4.工具类2 strconv
    实现string和其他数值类型之间的转换****很重要。
 */
func day_04(){
    /**
    1.初始化
     */
    name := "hello!"

    /**
    2.遍历
     */
    for i := 0; i < len(name); i++ {
        fmt.Printf("%d,",name[i]);
    }
    fmt.Println();
    for c := range name{
        fmt.Printf("%c,",name[c]);
    }
    /**
    3.工具类1 strings
    类似java strings。封装操作string的函数
     */

    /**
     4.工具类2 strconv
    实现string和其他数值类型之间的转换****很重要。
     */

}

func printSlice(x []int8){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}