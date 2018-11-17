package main

import "fmt"

/*二分法查找*/
func binaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	//判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("Not Find")
		return
	}
	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		binaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		binaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为%v \n", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 89, 1000, 1243}
	binaryFind(&arr, 0, len(arr)-1, 12)
}
