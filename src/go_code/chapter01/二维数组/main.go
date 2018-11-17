package main

import "fmt"

func main() {
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 1
	arr[2][3] = 1
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址%p\n", &arr2[0])
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1])

	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0])

	fmt.Println()
	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)

	/*二维数组的输出*/
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("%v\t", arr3[i][j])
		}
		fmt.Println()
	}

	for i, v := range arr3 {
		for j, v1 := range v {
			fmt.Printf("arr3[%v][%v]=%v \t", i, j, v1)
		}
		fmt.Println()
	}

	/*二维数组练习*/
	var scores [3][5]float64
	totalSum := 0.0
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班级的总分为%v，平均分为%v\n", i+1, sum, sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分为%v，所有班级平均分为%v \n", totalSum, totalSum/15)
}
