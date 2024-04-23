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

	// 풀이변수
	N, M int
	// result [][]int // map보단 슬라이스가 더 빠르겠지라 생각했지만 고려해야하는 인덱스가 2천만개면 얘기가 달라지지
)

func main() {

	/*
		23971. ZOAC4

		강의실을 예약하는데, 거리두기 수칙을 지키기 위해 일정 거리를 두고 앉아야함.
		최대 몇명을 수용할 수 있는가?

		H: 세로길이
		W: 가로 길이
		N: 세로 비워야하는 칸 (>= 조건)
		M: 가로 비워야하는 칸

		1) 최대한 많은칸이라면 1행1열부터 시작해야함.
		2) 세로로N칸, 가로로M칸만 떨어져야함

		N=2라면,
		H가 1,2,3일땐 1명만 가능
		H가 4,5,6일땐 2명만 가능
		H가 7,8,9일땐 3명만 가능
		N+1개의 수로 이루어짐.

		더 풀어서, N이 2라면 3가지의 경우에서 같은 값이 나오게됨. 따라서 3의 몫을 연관지을 수 있음
		H를 N으로 나눴을 때, 각각 묶음의 수가 공통된 값이 나오도록 하기위해서 H를 -1해준다.

		실제 앉을 수 있는 사람을 구하는 식 : (H-1 / N+1) +1

		세로에서 앉을수 있는 사람 * 가로에서 앉을 수 있는사람
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input := readLineInt()
	H, W, N, M := input[0], input[1], input[2], input[3]

	X := (H-1)/(N+1) + 1
	Y := (W-1)/(M+1) + 1
	fmt.Println(X * Y)

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
