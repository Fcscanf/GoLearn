package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "D:\\fcofficework\\DNS\\authorized_keys"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content))
}
