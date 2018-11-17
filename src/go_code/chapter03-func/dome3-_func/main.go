package main

import "fmt"

var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 - n2
	}
)

/*匿名函数*/
func main() {
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	a := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println("a=", a(10, 30))

	fmt.Println("Fun1=", Fun1(20, 40))
}
