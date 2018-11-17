package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSilce []Hero

func (hs HeroSilce) Len() int {
	return len(hs)
}

func (hs HeroSilce) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSilce) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
	/*等价的交换，不使用中间变量*/
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var intSilce = []int{1, 78, 97, 45, 34}
	sort.Ints(intSilce)
	fmt.Println(intSilce)

	var heroes HeroSilce
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄：%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println("排序后：")
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}
}
