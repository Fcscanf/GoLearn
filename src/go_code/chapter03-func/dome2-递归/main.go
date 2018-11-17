package main

import "fmt"

/*
函数递归-斐波那契数
*/
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func sum(n1 int, args ...int) int {
	sum := 1
	for i := 1; i < len(args); i++ {
		sum += args[0]
	}
	return sum
}

/*
初始化函数，main执行之前调用
*/
func init() {
	fmt.Println("初始化。。。")
}

/*
defer处理机制
1.defer进栈-先进后出，值入栈前后不变
*/
func sum1(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok1 n2=", n2)

	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := fbn(9)
	fmt.Println(res)

	res4 := sum(10)
	fmt.Println("res4=", res4)

	res1 := sum1(10, 30)
	fmt.Println("res=", res1)
}
