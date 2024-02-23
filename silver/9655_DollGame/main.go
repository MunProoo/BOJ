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
		마지막 돌을 가져가면 이긴다.
		둘은 완벽한 게임을 한다.
		항상 상근이가 먼저 한다.

		돌이 1개면 상근 Win
		돌이 2개면 창용 Win
		돌이 3개면 상근 Win
		돌이 4개면 창용 Win
		돌이 5개면 상근 Win
		돌이 6개면 창용	Win

		...
		돌이 홀수면 상근이가 이기네

		-- DP
		dp를 게임이 진행된 횟수, N을 돌 갯수라고 하면
		dp[1] = 1
		dp[2] = 2
		dp[3] = 1
		dp[4] = 2 (상근이가 3개 또는 1을 가져가고, 이후 창용이가 마무리) , 둘은 완벽한 게임을 하니까 최대한 적은 수로 게임할 것임

		dp[5] = 3 = min(dp[5-1],dp[5-3]) + 1
		dp[6] = min(dp[6-1], dp[6-3]) + 1

		...
		dp[N] = min(dp[N-1] , dp[N-3]) + 1

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

	if dp[N]%2 == 0 {
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
