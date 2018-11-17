package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		go func(i int) {
			for {
				fmt.Println("FCChina", i)
				/*手动交出控制权*/
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
