package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*文件的拷贝*/

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}

func main() {
	srcFile := "D:\\Photos\\Datapicture\\mmexport1530688562488.jpg"
	dstFile := "D:\\Photos\\1.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝完成！")
	} else {
		fmt.Println("拷贝失败，err=", err)
	}
}
