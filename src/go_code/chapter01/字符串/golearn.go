package main

import (
	"fmt"
	"strconv"
)

/*
字符串相关操作
*/
func main() {
	//===============================基本类型转String===========================================
	var num1 int = 99
	var num2 float64 = 23.453
	var b bool = true
	var myChar byte = 'h'
	var str string

	//使用fmt.Sprintf转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%v\n", str, str)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%v\n", str, str)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%v\n", str, str)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%v\n", str, str)
	fmt.Printf("str type %T str=%q\n", str, str)

	//使用strconv包的函数
	var num3 int = 99
	var num4 float64 = 23.453
	var b1 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%v\n", str, str)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%v\n", str, str)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b1)
	fmt.Printf("str type %T str=%v\n", str, str)
	fmt.Printf("str type %T str=%q\n", str, str)

	var num5 int = 2434
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%v\n", str, str)
	fmt.Printf("str type %T str=%q\n", str, str)

	/*================================String转基本数据类型=====================================*/
	var str1 string = "true"
	var b2 bool
	/*
		b2 , _ = strconv.ParseBool(str1)
		strconv.ParseBool(str1)该函数会返回两个值，（value bool，err error）
		使用_忽略获取第二个值
	*/
	b2, _ = strconv.ParseBool(str1)
	fmt.Printf("b type %T b=%v\n", b2, b2)

	var str2 string = "123443"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	n2 = int(n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)
}
