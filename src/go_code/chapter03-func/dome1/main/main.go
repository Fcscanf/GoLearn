package main

import (
	"fmt"
	"strings"
)

/*
累加器
*/
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n += x
		str += "a"
		fmt.Println("str = ", str)
		return n
	}
}

/*闭包实践*/
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))

	f1 := makeSuffix(".jpg")
	fmt.Println("文件处理后=", f1("winter"))
	fmt.Println("文件处理后=", f1("bg.jpg"))
}
