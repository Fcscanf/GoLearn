package main

import (
	"fmt"
	"io/ioutil"
)

/*将文件1的内容拷贝到文件2*/
func main() {
	file1Path := "D:\\fcofficework\\DNS\\1.txt"
	file2Path := "D:\\fcofficework\\DNS\\2.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file err=%v\n", err)
	}
}
