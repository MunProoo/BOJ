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
		11565 바이너리 게임

		문자열 a -> b 만들기.

		규칙 :
		1. a의 맨앞에 문자를 뺄 수 있다. 1001 -> 001 (빈 문자열은 못뺸다)
		2. a의 맨끝에 parity 추가가능 (짝수 parity. 1이 홀 수개면 1추가.  1이 짝 수개면 0추가)

		01110101
		101010

		a의
		1) 1이 홀수개라면 1 1개는 추가가 가능함
		2) 1이 짝수개라면 1의 개수가 a보다 b가 많을 수 없음

		이 조건을 토대로 작성
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	aCount := strings.Count(readLine(), "1")
	bCount := strings.Count(readLine(), "1")

	result := ""
	flag := true
	if aCount%2 == 1 {
		if aCount+1 < bCount {
			flag = false // 불가능
		}
	} else {
		if aCount < bCount {
			flag = false // 불가능
		}
	}

	if flag {
		result = "VICTORY"
	} else {
		result = "DEFEAT"
	}

	fmt.Println(result)
}

func readLine() string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
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

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
