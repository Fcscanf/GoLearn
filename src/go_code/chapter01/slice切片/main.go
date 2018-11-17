package main

import "fmt"

/*数组和切片*/

func main() {
	var intArr [5]int = [...]int{1, 22, 33, 54, 88}
	/*
		slice := intArr[1:3]
		1.slice
		2.intArr[1:3]
		3.引用intArr数组的起始下标为1，最后下标为3（但是不包含3）
	*/
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice = ", slice)
	fmt.Println("slice 的长度= ", len(slice))
	fmt.Println("slice 的容量= ", cap(slice))
	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p,slice[0]==%v\n", &slice[0], slice[0])
	fmt.Println()
	slice[1] = 45
	fmt.Println("intArr=", intArr)
	fmt.Println("slice = ", slice)

	/*string进行切片处理*/
	str := "hello@fcsca"
	slice1 := str[6:]
	fmt.Println("slice1=", slice1)

	arr1 := []byte(str)
	arr1[0] = 'f'
	str = string(arr1)
	fmt.Println("str=", str)

	/*字符串切片处理*/
	arr2 := []rune(str)
	arr2[0] = '樊'
	str = string(arr2)
	fmt.Println("str=", str)

	fbnSlice := fbn(10)
	fmt.Println("fbnSlice=", fbnSlice)
}

func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
