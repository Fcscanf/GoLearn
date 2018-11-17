package main

import "fmt"

func main() {
	var name string
	var age byte
	var sala float32
	var isPass bool
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水：")
	fmt.Scanln(&sala)

	fmt.Println("请输入是否通过考试：")
	fmt.Scanln(&isPass)

	fmt.Println("名字是：", name,
		"年龄是：", age,
		"薪水是：", sala,
		"是否通过考试：", isPass)
	fmt.Printf("名字是：%v,年龄是：%v,薪水是：%v,是否通过考试：%v\n", name, age, sala, isPass)

	/*
		按照指定的格式输入
	*/
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sala, &isPass)
	fmt.Printf("名字是：%v,年龄是：%v,薪水是：%v,是否通过考试：%v", name, age, sala, isPass)
}
