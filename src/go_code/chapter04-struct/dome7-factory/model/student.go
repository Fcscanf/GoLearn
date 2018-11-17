package model

type student struct {
	Name  string
	score float64
}

/*定义方法访问私有结构体*/
func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

/*定义方法访问私有属性*/
func (s *student) GetScore() float64 {
	return s.score
}
