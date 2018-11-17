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

func unmarshalStruct() {
	str := "{\"name\":\"孙悟空\",\"age\":500,\"birthday\":\"2011-11-11\",\"sal\":8000,\"skill\":\"如意七十二变\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化失败err：%v\n", err)
	}
	fmt.Printf("反序列化后monster：%v\n", monster)
}

func unmarshallMap() {
	str := "{\"address\":\"火云洞\",\"age\":10,\"name\":\"牛魔王\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("反序列化失败err：%v\n", err)
	}
	fmt.Printf("反序列化Map后：%v\n", a)
}

func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"19\",\"name\":\"TGH\"}," +
		"{\"address\":[\"华府\",\"影视帝国\"],\"age\":\"18\",\"name\":\"FCC\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("反序列化失败err：%v\n", err)
	}
	fmt.Printf("反序列化Slice后：%v\n", slice)
}

func main() {
	unmarshalStruct()
	unmarshallMap()
	unmarshalSlice()
}
