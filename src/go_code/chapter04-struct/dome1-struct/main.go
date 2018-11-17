package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
}

type Person struct {
	Name  string
	Age   int
	Score []float64
	ptr   *int
	slice []int
	map1  map[string]string
}

func main() {
	var cat1 Cat
	fmt.Printf("cat1 address=%p\n", &cat1)
	var cat2 Cat
	fmt.Printf("cat2 address=%p\n", &cat2)
	cat1.Name = "GH"
	cat1.Age = 56
	cat1.Color = "red"
	fmt.Println("cat1=", cat1)
	fmt.Println("猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("age=", cat1.Age)
	fmt.Println("color=", cat1.Color)

	var person1 Person
	fmt.Println(person1)
	/*切片使用前make分配空间*/
	person1.slice = make([]int, 10)
	person1.slice[0] = 100
	/*map使用前make分配空间*/
	person1.map1 = make(map[string]string)
	person1.map1["key1"] = "tom"
	fmt.Println(person1)

	person2 := person1
	person2.Name = "CK"
	fmt.Printf("pserson1 address=%p\n", &person1)
	fmt.Printf("pserson2 address=%p\n", &person2)
	fmt.Println("person1=", person1)
	fmt.Println("person2=", person2)

	var person3 Person = Person{
		Name: "FHGT",
		Age:  18,
	}
	fmt.Println("psrson3=", person3)

	var person4 *Person = &person3
	person4.Name = "FHGT~"
	person4.Age = 20
	fmt.Println("person3", person3)
	fmt.Println("person4", person4)
	fmt.Printf("person3=%p\n", &person3)
	fmt.Printf("person4=%p\n", &person4)
}
