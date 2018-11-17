package main

import "fmt"

func main() {
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))

	var strSlice []string = []string{"tom", "fc", "fs"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice size=", len(strSlice))
	fmt.Println("strSlice cap=", cap(strSlice))

	/*切片的遍历*/
	for i := 0; i < len(strSlice); i++ {
		fmt.Printf("slice[%v]=%v ", i, strSlice[i])
	}
	fmt.Println()
	for i, v := range strSlice {
		fmt.Printf("i = %v v = %v \n", i, v)
	}

	fmt.Println()
	slice2 := slice[1:2]
	slice2[0] = 100
	fmt.Println("slice2 = ", slice2)
	fmt.Println("slice = ", slice)

	/*切片的追加*/
	var slice3 []int = []int{100, 200, 300}
	slice3 = append(slice3, 400, 500)
	slice4 := append(slice3, slice3...)
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)

	/*切片的拷贝*/
	var slice5 []int = []int{1, 2, 3, 4, 5, 6}
	var slice6 = make([]int, 10)
	fmt.Println("slice5=", slice5)
	fmt.Println("slice6=", slice6)
	copy(slice6, slice5)
	fmt.Println("slice5=", slice5)
	fmt.Println("slice6=", slice6)
}
