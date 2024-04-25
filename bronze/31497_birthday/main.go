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
		31497. 생일 축하합니다~
		인터랙티브 문제.
		서버와 내 코드가 상호작용해야함.
		서버에게서 오늘 생일인 사람을 찾아내기
		서버는 1번은 거짓말을 할 수 있으니, 모두 2번씩 물어봐야함

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input := readLineInt()
	N := input[0] // 사람수
	peoples := make([]string, N)
	for i := 0; i < N; i++ {
		peoples[i] = readLine()
	}

	target := ""

	for i := 0; i < N; i++ {
		escapeFlag := false
		for j := 0; j < 2; j++ {
			fmt.Fprintf(os.Stdout, "? %s\n", peoples[i])
			answer := readLineInt()[0]
			if answer == 1 {
				target = peoples[i]
				escapeFlag = true
				break
			}
		}
		if escapeFlag {
			break
		}
	}

	fmt.Fprintf(os.Stdout, "! %s\n", target)
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
