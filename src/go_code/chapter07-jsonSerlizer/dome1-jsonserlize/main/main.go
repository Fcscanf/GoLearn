package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

/*结构体序列化*/
func NewMinsterStruct() {
	monster := Monster{
		Name:     "孙悟空",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "如意七十二变",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误err：%v\n", err)
	}
	fmt.Printf("Map序列化后=%v\n", string(data))
}

/*Map序列化*/
func MapSerlizer() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "牛魔王"
	a["age"] = 10
	a["address"] = "火云洞"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误err：%v\n", err)
	}
	fmt.Printf("monster序列化后=%v\n", string(data))
}

/*切片序列化*/
func SliceSerlizer() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "TGH"
	m1["age"] = "19"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "FCC"
	m2["age"] = "18"
	m2["address"] = [2]string{"华府", "影视帝国"}
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误err：%v\n", err)
	}
	fmt.Printf("切片序列化后=%v\n", string(data))
}

/*基本数据类型序列化*/
func FloatSerlize() {
	var num1 float64 = 245.56
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误err：%v\n", err)
	}
	fmt.Printf("基本数据类型序列化后=%v\n", string(data))
}

func main() {
	NewMinsterStruct()
	MapSerlizer()
	SliceSerlizer()
	FloatSerlize()
}
