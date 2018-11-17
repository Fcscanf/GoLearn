package main

import "fmt"

/*
指针相关操作
*/

func main() {
	/*
		基本数据类型在内存的布局
	*/
	var i int = 10
	fmt.Println("i在内存的地址是：", &i)

	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr在内存的地址是：%v\n", &ptr)
	fmt.Printf("ptr指向的值是：%v\n", *ptr)

	/*
		定义变量并使用指针修改变量的值
	*/
	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr1 *int
	ptr1 = &num
	*ptr1 = 10
	fmt.Println("num=", num)
}
