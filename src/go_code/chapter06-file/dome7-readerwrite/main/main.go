package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*在已存在文件清空原有内容进行追加*/
func main() {
	filePath := "D:\\fcofficework\\DNS\\1.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	/*关闭文件流*/
	defer file.Close()
	/*读取*/
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	/*写入文件*/
	str := "hello FCC您好！！！\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	/*因为writer是带缓存的，需要通过flush到磁盘*/
	writer.Flush()
}
