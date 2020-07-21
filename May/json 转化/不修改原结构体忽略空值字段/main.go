package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User
	Password *struct{} `json:"password,omitempty"`
}

// 优雅处理字符串格式的数字
type Card struct {
	ID    int64   `json:"id,string"`
	Score float64 `json:"score,string"`
}

func intAndStringDemo() {
	jsonStr1 := `{"id": "12312312", "score": "88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarshal jsonStr1 failed, err: %v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1)
}

func omitPasswordDemo() {
	u1 := User{
		Name:     "杭州",
		Password: "123123",
	}
	//b, _ := json.Marshal(PublicUser{&u1})
	b, _ := json.Marshal(PublicUser{User: &u1})
	fmt.Printf("%s\n", b)
}

//  json 整数变浮点数
func jsonDemo() {
	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))
	fmt.Printf("type:%T\n", m["count"])
	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	_ = json.Unmarshal(b, &m2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", m2["count"])
	fmt.Printf("type:%T\n", m2["count"])
}

func decoderDemo() {
	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))
	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	// 使用decoder方式反序列化，指定使用number类型
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&m2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("type:%T\n", m2["count"])  // json.Number
	// 将m2["count"]转为json.Number之后调用Int64()方法获得int64类型的值
	count, err := m2["count"].(json.Number).Int64()
	if err != nil {
		fmt.Printf("parse to int64 failed, err:%v\n", err)
		return
	}
	fmt.Printf("type:%T\n", int(count)) //
}
func main() {
	//omitPasswordDemo()
	//intAndStringDemo()
	//jsonDemo()
	decoderDemo()
}
