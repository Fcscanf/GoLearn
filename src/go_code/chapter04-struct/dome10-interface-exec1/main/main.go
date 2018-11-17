package main

import "fmt"

type Monkey struct {
	Name string
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "学会了游泳！")
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习学会了飞翔！")
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "会爬树！")
}

type LittleMonkey struct {
	Monkey
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
