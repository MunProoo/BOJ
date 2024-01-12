package main

type Queue []int // 큐타입 추가

func (que *Queue) push(node int) {
	*que = append(*que, node)
}

func (que *Queue) pop() (answer int) {
	answer = (*que)[0]
	*que = (*que)[1:]
	return
}

func (que *Queue) isEmpty() bool {
	if len(*que) == 0 {
		return true
	} else {
		return false
	}
}
