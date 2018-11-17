package main

import (
	"fmt"
	"sort"
)

/*map-引用类型*/

func main() {
	/*map赋值的定义*/
	var a map[string]string
	/*make分配内存*/
	a = make(map[string]string, 10)
	a["no1"] = "FC"
	a["no2"] = "TR"
	fmt.Println(a)

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "南京"
	cities["no4"] = "江苏"
	fmt.Println(cities)

	heroes := map[string]string{
		"hero1": "RT",
		"hero2": "YT",
		"hero3": "GT",
	}
	fmt.Println("heroes=", heroes)

	/*map的遍历*/
	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
	fmt.Println("cities有", len(cities), "对key-value")

	/*操作练习*/
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["name"] = "江苏省南京市"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Grt"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["name"] = "北京市"
	fmt.Println(studentMap)

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
		fmt.Println()
	}

	/*map的crud
	存在即更新
	存在即新增
	存在即删除
	*/
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("有no1 key 值为%v \n", val)
	} else {
		fmt.Printf("没有no1 key\n")
	}

	delete(cities, "no1")
	fmt.Println(cities)
	/*刷一个空的进去即删除所有map*/
	cities = make(map[string]string)

	/*map切片*/
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "TGH"
		monsters[0]["age"] = "500"
	}
	fmt.Println(monsters)

	newMonster := map[string]string{
		"name": "DFG",
		"age":  "700",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)

	/*map的排序*/
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)

	/*key放到切片进行遍历*/
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Printf("map[%v]=%v", k, map1[k])
	}
}
