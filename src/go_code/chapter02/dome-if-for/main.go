package main

import (
	"fmt"
	"math"
)

func main() {
	if age := 20; age > 18 {
		fmt.Println("你年龄大于18，要对自己负责！")
	} else {
		fmt.Println("Hello,Body")
	}

	var year int = 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("2020是闰年")
	} else if year < 100 {
		fmt.Println("输入错误")
	}

	var a float64 = 3
	var b float64 = 100.0
	var c float64 = 6.0

	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x2=%v\n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v\n", x1)
	} else {
		fmt.Println("无解")
	}

	var height int32
	var money float32
	var handsome bool
	fmt.Println("请输入身高（厘米）")
	fmt.Scanln(&height)
	fmt.Println("请输入财富（千万）")
	fmt.Scanln(&money)
	fmt.Println("请输入是否帅（true/false）")
	fmt.Scanln(&handsome)
	if height > 180 && money > 1.0 && handsome {
		fmt.Println("嫁给他")
	} else if height > 180 || money > 1.0 || handsome {
		fmt.Println("勉强")
	} else {
		fmt.Println("不嫁")
	}

	var second float64
	fmt.Println("请输入秒数：")
	fmt.Scanln(&second)
	if second <= 8 {
		var gender string
		fmt.Println("请输入性别：")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入决赛男子组")
		} else {
			fmt.Println("进入决赛女子组")
		}
	} else {
		fmt.Println("out...")
	}

	var month byte
	var age byte
	var price float64 = 60
	fmt.Println("请输入游玩的月份：")
	fmt.Scanln(&month)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("%v月 %v 年龄 票价：%v", month, age, price/3)
		} else if age >= 18 {
			fmt.Printf("%v月 %v 年龄 票价：%v", month, age, price)
		} else {
			fmt.Printf("%v月 %v 年龄 票价：%v", month, age, price/2)
		}
	} else {
		if age >= 18 && age < 60 {
			fmt.Println("淡季成人票价：40")
		} else {
			fmt.Println("淡季儿童票价：20")
		}
	}

	/*
		fallthrough穿透功能
	*/
	var key byte
	fmt.Println("请输入一个字符：a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("赵乾坤")
		fallthrough
	case 'b':
		fmt.Println("古赛昆")
	case 'c':
		fmt.Println("戴维超")
	default:
		fmt.Println("输入错误！")
	}

	/*
		for循环
	*/
	for i := 1; i < 10; i++ {
		fmt.Println("Hello,Fc")
	}
	var str string = "helloworld"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	fmt.Println()
	str = "abc"
	for index, val := range str {
		fmt.Printf("index=%d,val=%c \n", index, val)
	}

	var classNum int = 2
	var stuNum int = 5
	var totalSum float64 = 0.0
	var passCount int = 0
	for j := 1; j <= classNum; j++ {
		sum := 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d个班第%d个学生的成绩：\n", j, i)
			fmt.Scanln(&score)
			sum += score
			if score >= 60 {
				passCount++
			}
		}
		fmt.Printf("第%d个班级的平均分是%v\n", j, sum/float64((stuNum)))
		totalSum += sum
	}
	fmt.Printf("各个班的总成绩%v 所有班级的平均分是%v\n",
		totalSum, totalSum/float64(stuNum*classNum))
	fmt.Printf("及格人数为：%v\n", passCount)

	var totalLevel int = 9
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v 8 %v = %v", j, i, j*i)
		}
		fmt.Println()
	}

	/*
		指定标签形式使用break
	*/
lable2:
	for i := 0; i <= 4; i++ {
		//lable1:
		for j := 0; j <= 10; j++ {
			if j == 2 {
				break lable2
			}
			fmt.Println("j=", j)
		}
	}

}
