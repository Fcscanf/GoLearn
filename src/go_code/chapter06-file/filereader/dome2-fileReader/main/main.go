package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*缓冲式读取文件*/
func main() {
	file, err := os.Open("d:\\Photos\\Screenshots\\暗物质\\IMG_20180927_194619.jpg")
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束!")
}
