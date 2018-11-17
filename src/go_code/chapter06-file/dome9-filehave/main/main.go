package main

import (
	"fmt"
	"os"
)

/*判断文件以及目录是否存在*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("当前文件存在！")
		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Println("当前文件不存在！")
		return false, nil
	}
	return false, nil
}

func main() {
	path := "D:\\fcofficework\\2.txt"
	PathExists(path)
}
