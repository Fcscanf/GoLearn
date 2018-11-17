package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

/*struct tag-`json:"name"`*/

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

/*结构体类型转换*/

func main() {
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)

	monster := Monster{"古赛昆", 500, "Tadk"}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
