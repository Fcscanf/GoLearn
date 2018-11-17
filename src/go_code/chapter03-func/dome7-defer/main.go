package main

import (
	"errors"
	"fmt"
)

/*异常处理*/

func main() {
	test()
	fmt.Println("main")

	test01()
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

/*自定义错误*/
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误...")
	}
}

func test01() {
	err := readConf("config.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test01()继续执行...")
}
