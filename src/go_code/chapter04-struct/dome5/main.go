package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (stu Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

/*重写String方法使用地址可以直接调用*/

func main() {
	stu := Student{
		Name: "Tom",
		Age:  80,
	}
	fmt.Println(&stu)
}
