package main

import "fmt"

//定义全局变量
var k1 = 100
var k2 = 200
var k3 = "fgc"

//全局变量简洁定义
var (
	k4    = 500
	k5    = 14
	name1 = "lk"
)

/*
变量
*/
func main() {
	fmt.Println("Hello,Go")
	fmt.Println("D:\\fcofficework\\Acf-My work\\临时工作文件夹\\班级材料\\材料\\素质拓展申报表\\2018秋素拓分认定")
	fmt.Println("天龙八部雪山飞狐\r张飞")

	/*====================变量先定义再使用=======================*/
	//变量不设置值则使用默认值
	var i int
	fmt.Println(i)

	//编译器自行判断数据类型
	var num = 10.11
	fmt.Println(num)

	//省略变量定义关键字var，注意:=左侧的变量不应该是已经申明过的，否则编译报错
	name := "tom"
	fmt.Println(name)

	//多变量声明
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	//多变量定义赋值
	var n4, namenick, n5 = 100, "fc", 99
	fmt.Println(n4, namenick, n5)

	n6, tg := 55, "fg"
	fmt.Println(n6, tg)

	//输出全局变量
	fmt.Println(k1, k2, k3, k4, k5, name1)
}

func enums() {
	/*普通枚举类型*/
	const (
		/*iota可以自增*/
		a = iota
		a1
		a2
		a3
	)
	fmt.Println(a, a1, a2, a3)

	/*自增枚举类型*/
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
