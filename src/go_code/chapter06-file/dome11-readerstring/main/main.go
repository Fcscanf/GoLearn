package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*统计文件的字符个数*/

type CharCount struct {
	/*英文的个数*/
	ChCount int
	/*数字的个数*/
	NumCount int
	/*空格的个数*/
	SpaceCount int
	/*其他字符的个数*/
	OtherCount int
}

func main() {
	fileName := "D:\\fcofficework\\DNS\\1.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数为：%v 数字的个数为：%v 空格的个数为：%v 其他字符的个数为：%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
