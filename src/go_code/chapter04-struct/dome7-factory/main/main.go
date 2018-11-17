package main

import (
	"fmt"
	"go_code/chapter04-struct/dome7-factory/model"
)

func main() {

	/*var stu = model.Student{
		Name : "Tom",
		Score : 78.9,
	}
	fmt.Println(stu)*/

	var stu = model.NewStudent("Tom", 78.9)
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
