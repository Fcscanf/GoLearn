package main

import "fmt"

type Ponit struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Ponit
}

type Rect1 struct {
	leftUp, rightDown *Ponit
}

func main() {
	r1 := Rect{Ponit{1, 2}, Ponit{3, 4}}
	fmt.Printf("r1.leftUP.x address=%p r1.leftUp.y address=%p r1.rightDown.x address=%p r1.rightDown.y address=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	r2 := Rect1{&Ponit{10, 20}, &Ponit{30, 40}}
	fmt.Printf("r2.leftUp address=%p r2rightDown address=%p \n", &r2.leftUp, &r2.rightDown)
}
