package main

import (
	"bufio"
	"fmt"
	"os"
)

/*在已存在文件清空原有内容重新写入*/
func main() {
	filePath := "D:\\fcofficework\\DNS\\1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	defer file.Close()
	str := "hello FCC\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	/*因为writer是带缓存的，需要通过flush到磁盘*/
	writer.Flush()
}
