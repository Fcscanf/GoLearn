package main

import "fmt"

/*冒泡排序*/

func main() {
	arr := [5]int{24, 69, 28, 78, 19}
	BubbleSort(&arr)
}

func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", *arr)
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后arr=", *arr)
}
