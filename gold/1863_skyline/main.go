package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"os"
)

var (
	writer *bufio.Writer
	reader *bufio.Reader
)

func main() {

	/*
		1863 스카이라인 쉬운거

		왼쪽에서 스카이라인을 보아 갈 때, 고도가 바뀌는 지점의 좌표 x와 y가 주어진다.
		최소 건물의 개수를 출력하라

		0을 기준으로 고도가 몇 번 바뀌는지를 알면 건물의 최소개수를 알 수 있다?
		Set을 이용해본다.

		Set을 이용하면 다음 케이스에서 오답
		4
		1 4
		2 3
		3 5
		4 4
		출력:3,  정답:4

		스택을 이용, 고도가 낮아지면 연결될 건물이 없음.

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	stack := make(Stack, 0)

	N := readLineInt()[0]
	result := 0
	for i := 0; i < N; i++ {
		input := readLineInt()
		_, y := input[0], input[1]

		if y == 0 {
			result += len(stack)
			stack = make(Stack, 0)
		} else {
			if y < stack.Top() {
				for y < stack.Top() {
					stack.Pop()
					result++
				}
			}
			if y == stack.Top() {
				continue
			}
			stack.Push(y)
		}
	}

	if len(stack) != 0 {
		result += len(stack)
	}

	fmt.Println(result)

}

func readLineInt() []int {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	inputSlice := strings.Split(input, " ")

	numList := make([]int, len(inputSlice))
	for i, val := range inputSlice {
		numList[i], _ = strconv.Atoi(val)
	}

	return numList
}

// Set
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Top() int {
	if len(*s) == 0 {
		return -1
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	top := len(*s) - 1
	result := (*s)[top]
	*s = (*s)[:top]
	return result
}

// func abs(a int) int {
// 	if a > 0 {
// 		return a
// 	} else {
// 		return -a
// 	}
// }

// func readLine() string {
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSuffix(input, "\n")
// 	input = strings.TrimSuffix(input, "\r")

// 	return input
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }
