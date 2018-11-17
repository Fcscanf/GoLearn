package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("d:\\Photos\\Screenshots\\暗物质\\IMG_20180927_194619.jpg")
	if err != nil {
		fmt.Println("open file err", err)
	}
	fmt.Printf("file=%v", file)
	err1 := file.Close()
	if err1 != nil {
		fmt.Println("close file err = ", err1)
	}
}
