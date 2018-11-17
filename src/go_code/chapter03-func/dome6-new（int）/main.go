package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("num1的类型%T,num1的值为=%v,num1的地址%v\n", num1, num1, &num1)
	num2 := new(int)
	*num2 = 100
	fmt.Printf("num2的类型%T,num2的值为=%v,num2的地址%v\n,num2这个指针指向的值为%v", num2, num2, &num2, *num2)
}
