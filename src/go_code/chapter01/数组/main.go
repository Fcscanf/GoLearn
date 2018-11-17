package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*数组和切片-数组改值通过寻址改值-引用传递*/
func main() {
	var hens [6]float64
	hens[0] = 1.9
	hens[1] = 4.9
	hens[2] = 5.9
	hens[3] = 2.9
	hens[4] = 8.9
	hens[5] = 5.9
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	fmt.Printf("总体重：%v，平均体重：%v\n", totalWeight, int(totalWeight/6))

	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请%v号学生成绩：", i)
		fmt.Scanln(&score[i])
	}
	for i := 0; i < len(score); i++ {
		fmt.Printf("student[%v]`score is %v\n", i, score[i])
	}

	/*初始化数组*/
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{1, 2, 3}
	fmt.Println("numArr02=", numArr02)

	var numArr03 = [...]int{1, 2, 3}
	fmt.Println("numArr03=", numArr03)

	/*通过下标指定值*/
	var numArr04 = [...]int{1: 800, 0: 900, 2: 456}
	fmt.Println("numArr04=", numArr04)

	/*for-range访问数组*/
	var heroes [3]string = [3]string{"FC", "CK", "JK"}
	for i, v := range heroes {
		fmt.Printf("i = %v,v = %v\n", i, v)
	}

	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c\n", myChars[i])
	}

	var intArr [5]int = [...]int{1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxValIndex := 0
	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v\n", maxVal, maxValIndex)

	var intArr1 [5]int = [...]int{1, -1, 9, 90, 11}
	sum := 0
	for _, val := range intArr1 {
		sum += val
	}
	fmt.Printf("sum=%v ,平均值=%v\n", sum, float64(sum)/float64(len(intArr1)))

	/*随机生成五个数，并将其反转打印*/
	var intArr2 [5]int
	len := len(intArr2)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr2[i] = rand.Intn(100)
	}
	fmt.Println("交换前", intArr2)
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr2[len-1-i]
		intArr2[len-1-i] = intArr2[i]
		intArr2[i] = temp
	}
	fmt.Println("交换后", intArr2)
}
