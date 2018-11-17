package main

import (
	"fmt"
	"os"
)

/*解析命令行参数*/
func main() {
	fmt.Println("命令行参数有：", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}
