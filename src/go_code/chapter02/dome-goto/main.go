package main

import "fmt"

/*
goto相关操作
*/
func main() {
	fmt.Println("ok1")
	fmt.Println("ok2")
	goto lable1
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
lable1:
	fmt.Println("ok6")
	fmt.Println("ok7")
	fmt.Println("ok8")
	fmt.Println("ok9")
}
