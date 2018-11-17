package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, x)
		case float64, float32:
			fmt.Printf("第%v个参数是float类型，值是%v\n", index, x)
		case int, int8, int16, int32, int64:
			fmt.Printf("第%v个参数是int类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数是Student类型，值是%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数是*Student类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数类型不确定，值是%v\n", index, x)
		}
	}
}

type Student struct {
}

func main() {
	var n1 float32 = 1.1
	var n2 int64 = 30
	var name string = "tom"
	address := "故宫"
	n4 := 500

	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, name, address, n4, stu1, stu2)
}
