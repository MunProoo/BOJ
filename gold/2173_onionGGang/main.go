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
	// defaultArr []byte
	sb      strings.Builder
	arrSize int // defaultArr 크기
)

func main() {

	/*
		2173 양파깡 만들기

		문제 :
		N*N 보드가 주어진다.
		M개의 양파깡을 만들어야한다.
		양파깡은 최소 3*3의 가운데 공간이 빈 직사각형이다.
		양파깡은 현재 board에서 가장 합이 크도록 만들어야한다.

		풀이 :

		1. 2차원 배열의 누적합 구하기
		https://nahwasa.com/entry/%EB%88%84%EC%A0%81-%ED%95%A9prefix-sum-2%EC%B0%A8%EC%9B%90-%EB%88%84%EC%A0%81%ED%95%A9prefix-sum-of-matrix-with-java





	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	// 1st Input
	input = readLineInt()
	N, M := input.([]int)[0], input.([]int)[1]

	// Make Board
	board := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		board[i] = make([]int, 1)
	}

	// 2nd Input
	for i := 1; i <= N; i++ {
		input = readLineInt()
		board[i] = append(board[i], input.([]int)...)
	}

	prefixSum := make([][]int, N+1) // 직사각형 단위의 누적합
	for i := 0; i <= N; i++ {
		prefixSum[i] = make([]int, N+1)
	}

	// 직사각형 영역별 누적합 구하기
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + board[i][j]
		}
	}

	// 최고 맛 구해보기
	max := 0
	for x1 := 1; x1 <= N; x1++ {
		for y1 := 1; y1 <= N; y1++ {
			for x2 := x1 + 2; x2 <= N; x2++ {
				for y2 := y1 + 2; y2 <= N; y2++ {
					flavor := calcOnion(x1, x2, y1, y2, prefixSum, board)
					if flavor > max {
						max = flavor
					}
				}
			}
		}
	}

	fmt.Printf("%d\n", max)
	// fmt.Printf("%v", board)
	fmt.Println(M)

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

func calcOnion(x1, x2, y1, y2 int, prefixSum, board [][]int) int {
	// 최종 누적 - x1,y1 ~ x2,y2의 누적합
	totalFlavor := prefixSum[x2][y2] - prefixSum[x1-1][y2] - prefixSum[x2][y1-1] + prefixSum[x1-1][y1-1]

	// emptyFlavor -> x1+1, y1+1 ~ x2-1, y2-1까지의 합
	x1, y1 = x1+1, y1+1
	x2, y2 = x2-1, y2-1
	emptyFlavor := prefixSum[x2][y2] - prefixSum[x1-1][y2] - prefixSum[x2][y1-1] + prefixSum[x1-1][y1-1]

	return totalFlavor - emptyFlavor
}
