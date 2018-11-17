package main

import "fmt"

type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) Print1(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println("是奇数！")
	} else {
		fmt.Println("是偶数！")
	}
}

func (mu *MethodUtils) Print2(n int, m int, key string) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (calcuator *Calcuator) getSum() float64 {
	return calcuator.Num1 + calcuator.Num2
}

func (calcuator *Calcuator) getSub() float64 {
	return calcuator.Num1 - calcuator.Num2
}

func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	}
	return res
}

func main() {
	var mu MethodUtils
	mu.Print()
	mu.Print1(10, 2)
	fmt.Println("面积为：", mu.area(34, 12))
	mu.JudgeNum(11)

	var calcuator Calcuator
	calcuator.Num1 = 1.2
	calcuator.Num2 = 2.2
	fmt.Println("sum=", calcuator.getSum())
	fmt.Println("sub=", calcuator.getSub())
}
