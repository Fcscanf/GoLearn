package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

/*显示成绩*/
func (stu *Student) ShowInfo() {
	fmt.Println("学生姓名：", stu.Name, "学生年龄：", stu.Age, "学生成绩：", stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student
}

type Graduate struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生考试中...")
}

func (g *Graduate) testing() {
	fmt.Println("大学生考试中...")
}

func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "Tom"
	pupil.Student.Age = 18
	pupil.testing()
	pupil.SetScore(80)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "Mary"
	graduate.Student.Age = 20
	graduate.testing()
	graduate.SetScore(90)
	graduate.Student.ShowInfo()
}
