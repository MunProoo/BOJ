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
		탁자위에 돌 N개가 있다. 턴을 번갈아가며 돌을 1개 혹은 3개 혹은 4개를 가져간다.
		마지막 돌을 가져가면 이긴다.
		둘은 완벽한 게임을 한다.
		항상 상근이가 먼저 한다.

		판 수가 홀 수면 상근Win, 짝수면 창용Win


		9655에서 dp 조건을 추가하면 될듯?

		둘은 완벽한 게임을 하니까 최대한 적은 판 수로 게임할 것임

		-- DP
		상근이 승리 유무를 dp값으로 놓는다.

		dp[N]을 "내 차례에 돌의 개수가 N개일 때 승리가 가능한가"라고 치면
		dp[N]이 true면 무조건 승리

		dp[1] = true
		dp[2] = false
		dp[3] = true
		dp[4] = true

		N이 5개이상이면 1,3,4를 빼면서 "상대가 무조건 승리하지 못하는 개수"를 넘겨주어야한다.
		N이 5일 때, 돌을 3개 빼서 dp[2]를 넘겨주면 상근이가 승리하게 된다.
			dp[5] = true

		N이 6이면
			1개를 빼면 dp[5]가 나오고 dp[5]는 무조건 승리가 가능하므로 (패배)
			3개를 빼면 dp[3]이고 dp[3]은 무조건 승리 가능하므로 (패배)
			4개를 빼면 dp[2]고 dp[2]는 상대방이 무조건 승리가 안되므로 내턴이 돌아옴 (승리)
			dp[6] = true

		N이 7이면
			1개를 빼면 dp[6]이 나오고, dp[6]을 상대가 먹으면 무조건 승리이므로 (패배)
			3개를 빼면 dp[4]이고 dp[4]을 상대가 먹으면 무조건 이기니까 (패배)
			4개를 빼면 dp[3]고 dp[3]을 상대가 먹으면 무조건 이기니까 (패배)
			dp[7] = false


	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()
	N := input.([]int)[0] // N

	// DP 풀이
	dp := make([]bool, 1001)
	dp[1] = true
	dp[2] = false
	dp[3] = true
	dp[4] = true

	// 상근이의 입장에서 항상 최선의 수를 구하는 로직
	for i := 5; i <= 1000; i++ {
		if dp[i-1] && dp[i-3] && dp[i-4] {
			// 후공에게 무조건 승리를 넘겨주는 경우
			dp[i] = false
		} else {
			// 승리할 방법이 있음
			dp[i] = true
		}
	}

	if !dp[N] {
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
