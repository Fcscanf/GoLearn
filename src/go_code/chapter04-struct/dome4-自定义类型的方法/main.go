package main

import "fmt"

/*结构体的方法-自定义类型都有方法*/

type Person struct {
	Name string
}

/*指定该方法属于Person*/
func (person Person) test() {
	fmt.Println("Name:", person.Name)
}

func (person Person) speak() {
	fmt.Println(person.Name, "是一个好人！")
}

func (person Person) skill() {
	res := 0
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算的结果是：", res)
}

func (person Person) skillResult(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算的结果是：", res)
}

func (person Person) skillResultAdd(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var person1 Person
	person1.Name = "Tom"
	person1.test()
	person1.speak()
	person1.skill()
	person1.skillResult(78)
	res := person1.skillResultAdd(12, 34)
	fmt.Println(res)

	var c Circle
	c.radius = 4.0
	fmt.Println("面积是：", c.area())
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

/*一般为值传递，效率低下，常用方法绑定结构体的指针类型*/

func (c *Circle) area1() float64 {
	return 3.14 * c.radius * c.radius
}
