package main

type Queue []int

/*添加*/
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

/*读取*/
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

/*判断非空*/
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
