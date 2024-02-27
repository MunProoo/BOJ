package main

import (
	"bufio"
	"strconv"
	"strings"

	"os"
)

var (
	writer *bufio.Writer
	reader *bufio.Reader
	sb     strings.Builder

	// 풀이변수
)

func main() {

	/*
		9655 돌 게임 (게임 이론)

		2명이서 즐긴다.
		탁자위에 돌 N개가 있다. 턴을 번갈아가며 돌을 1개 혹은 3개를 가져간다.
		마지막 돌을 가져가면 진다.
		둘은 완벽한 게임을 한다.
		항상 상근이가 먼저 한다.

		9655와 결과 반대로하면 됨

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()
	N := input.([]int)[0] // N

	/*
		// 직관적 풀이
		if N%2 == 0 {
			sb.WriteString("CY")
		} else {
			sb.WriteString("SK")
		}
	*/

	// DP 풀이
	dp := make([]int, 1001)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 1
	dp[4] = 2
	for i := 5; i <= 1000; i++ {
		dp[i] = min(dp[i-1], dp[i-3]) + 1
	}

	if dp[N]%2 == 1 {
		sb.WriteString("CY\n")
	} else {
		sb.WriteString("SK\n")
	}

	writer.WriteString(sb.String())

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

func readLine() string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
