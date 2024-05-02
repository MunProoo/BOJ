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
	sb     strings.Builder
)

func main() {

	/*
		25907. 양과 늑대
		인터랙티브 문제.

		질문하면 해당 날짜의 양의 수를 알려줌.
		(1)N일 동안 회의장에 양 or 늑대가 온다.
		(2)양의 수 == 늑대의 수 -> 무지 출근
		(3)첫 날은 양이 오고, 마지막날에는 항상 늑대가 양보다 많음
			-> 무조건 양 == 늑대인 날이 있다

		20번 안에 무지 출근 날짜 중 아무거나 1개 찾아서 출력
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input := readLineInt()
	N := input[0] // 총 출석일

	/*
		1. 마지막날 양의 수를 물어본다 (각각 얼마나 오는지 확인)
		2. 늑대 - 양을 구해서 해당 일수를 마지막에서 빼서 물어본다
		2-1. 늑대 != 양이고, 양 < 늑대면 2.를 반복한다.
		2-2. if (늑대 < 양) : 2.에서 구한 일 - 현재 구한 일 사이를 이분탐색

		아니 근데 이렇게 해도 양 수가 늘어나버리면 이분탐색이 안돼.

		그냥, 조건을
		늑대 > 양인가?
		늑대 < 양인가?
		이렇게만 해야할듯
	*/

	start := 0
	end := N
	reqDay := (start + end) / 2
	for i := 0; i < 19; i++ {
		fmt.Fprintf(os.Stdout, "? %d\n", reqDay)
		sheeps := readLineInt()[0]

		wolfs := reqDay - sheeps // 늑대 수

		if wolfs == sheeps {
			break
		} else if wolfs < sheeps {
			// 양이 더 많음 : 늑대 > 양인 날짜로 가야함 : 현재 기준 "미래"로
			start = reqDay + 1
			reqDay = (start + end) / 2
		} else {
			// 늑대가 더 많음 : 늑대 < 양인 날짜로 가야함 : 현재 기준 "과거"로
			end = reqDay - 1
			reqDay = (start + end) / 2
		}
	}

	fmt.Fprintf(os.Stdout, "! %d\n", reqDay)
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

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func readLine() string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
