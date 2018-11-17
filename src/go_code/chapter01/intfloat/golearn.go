package main

import (
	"fmt"
	"unsafe"
)

/*
整型-浮点型
*/
func main() {

	/*===================================整型==============================*/

	/*变量使用多次定义取最后的值-即在同一类型范围内不断的变换
	  在同一个范围内不能重复定义*/
	var i int = 10
	i = 30
	i = 50
	fmt.Println(i)

	var j int = 4
	/*整型相加*/
	fmt.Println(i + j)
	/*字符串拼接*/
	var name string = "Hello"
	var name1 string = "World"
	fmt.Println(name + name1)

	/*整型的范围*/
	var name2 int8 = 127
	fmt.Println(name2)
	/*
		格式化输出变量的类型
		返回变量的字节数
	*/
	fmt.Printf("name2的类型 %T\n占用的字节数是 %d\n", name2, unsafe.Sizeof(name2))

	/*===================================浮点型==============================*/
	var price float32 = 89.12
	fmt.Println(price)

	//	浮点型精度不同会造成精度损失
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1 = ", num1, "num2 ", num2)

	//GoLang的浮点型默认是float64位的
	var num3 = 1.1
	fmt.Printf("num3的数据类型为 %T\n", num3)

	//十进制数的形式必须有小数点
	num4 := 5.12
	num5 := .123
	fmt.Println("num4 = ", num4, ",num5 =", num5)

	//科学记数法
	num6 := 5.1234e2
	num7 := 5.1234E2
	fmt.Println("num6 = ", num6, ",num7 =", num7)

	/*===================================字符=================================*/
	var c1 byte = 'a'
	var c2 byte = '0'

	//直接输出byte值，输出对应的ASCII码值
	fmt.Println("c1=", c1, "c2=", c2)
	//输出对应的字符串需要格式化输出
	fmt.Printf("c1=%c, c2=%c\n", c1, c2)

	//var c3 byte = '北' overflow溢出
	var c3 int = '北'
	fmt.Printf("c3=%c", c3)

	/*===================================字符串=================================*/
	//字符串不可变
	//双引号-可识别转义字符
	//``使用反引号可以原样输出
}
