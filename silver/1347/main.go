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
)

func main() {

	/*
		1347 미로만들기

		홍준이의 움직임을 적은 노트만으로 미로 지도 만들기
		홍준이는 미로의 남쪽을 보고 있다.

		sol)
		상하좌우로 최대 50씩 갈 수 있으므로 가져야할 미로배열은 최대 100 * 100
		-> 101 * 101이 맞겠네. 딱 (50,50)에서 이동한다 치고

		음.. 문자 그대로 쭉 진행한다음에 이동한 부분만 체크해서
		미로로 만들어낸다?
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()
	_ = input.([]int)[0] // N
	input = readLine()
	note := input.(string)

	maxCount := 101
	maze := make([][]bool, maxCount)
	for i := 0; i < maxCount; i++ {
		maze[i] = make([]bool, maxCount)
	}

	// 이동 기록
	x, y := 50, 50
	for _, char := range note {
		switch char {
		case 'F':

		}
	}

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
