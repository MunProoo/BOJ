package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"os"
)

var (
	// 입출력
	writer *bufio.Writer
	reader *bufio.Reader
	sb     strings.Builder
	sb2    strings.Builder

	// 풀이 변수
	N, M      int
	prefixSum [][]int
	board     [][]int

	maxVisited [][]bool // 이미 사용한 최대값은 제외
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

		2. 최대 맛이 1개라면 정답을 찾는 듯한데, 최대 맛이 중복이고, 첫 번째로 고른 최대맛으로는 M개를 못고른다면 실패함
		M개를 잘라내는 경우를 찾을 수 있도록 백트래킹을 사용해야 할듯



	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	// 1st Input
	input = readLineInt()
	N, M = input.([]int)[0], input.([]int)[1]

	// Make Board
	board = make([][]int, N+1)
	for i := 0; i <= N; i++ {
		board[i] = make([]int, 1)
	}

	// 2nd Input
	for i := 1; i <= N; i++ {
		input = readLineInt()
		board[i] = append(board[i], input.([]int)...)
	}

	visited := make([][]bool, N+1) // 잘라낸건지 확인.
	for i := 0; i <= N; i++ {
		visited[i] = make([]bool, N+1)
	}

	prefixSum = make([][]int, N+1) // 직사각형 단위의 누적합
	for i := 0; i <= N; i++ {
		prefixSum[i] = make([]int, N+1)
	}

	// 직사각형 영역별 누적합 구하기
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + board[i][j]
		}
	}

	// 양파깡 구하기 (arg로 startX 등은 Max가 중복이 나는 경우만 의미 있음. 0을 넣으면 초기 시작이라는 것.)
	result := findOnionGGang(1, 1, 3, 3, 0, visited)
	if !result {
		sb.WriteString("0\n")
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

func calcOnion(x1, x2, y1, y2 int, prefixSum, board [][]int) int {
	// 최종 누적 - x1,y1 ~ x2,y2의 누적합
	totalFlavor := prefixSum[x2][y2] - prefixSum[x1-1][y2] - prefixSum[x2][y1-1] + prefixSum[x1-1][y1-1]

	// emptyFlavor -> x1+1, y1+1 ~ x2-1, y2-1까지의 합
	x1, y1 = x1+1, y1+1
	x2, y2 = x2-1, y2-1
	emptyFlavor := prefixSum[x2][y2] - prefixSum[x1-1][y2] - prefixSum[x2][y1-1] + prefixSum[x1-1][y1-1]

	return totalFlavor - emptyFlavor
}

// 현재 판에서 가장 맛있는 맛을 만드는 조건이 여러개라면 각각 경우를 전부 확인 (브루트포스라 해야하나)
// M개의 양파깡을 만들 수 있다면 true를 반환함.
func findOnionGGang(startX, startY, endX, endY, depth int, visited [][]bool) bool {

	for i := depth; i < M; i++ {

		a, b, c, d := 0, 0, 0, 0
		max := -999999999
		checkFlag := false

		// 현재 상황에서 가장 맛있는 값 return
		maxTaste := findMaxTaste(visited)

		for x1 := 1; x1 <= N-2; x1++ {
			for y1 := 1; y1 <= N-2; y1++ {

				if visited[x1][y1] {
					// 이미 사용한 양파깡 판임
					continue
				}
				// ContinueLoop:
				for x2 := x1 + 2; x2 <= N; x2++ {
				ContinueLoop: // Label을 한칸 위로 했어서 x2가 올라가는게 아니라 y1부터 올라갔기때문에 너무 과하게 스킵했었음
					for y2 := y1 + 2; y2 <= N; y2++ {

						// Max값이 중복인 경우 다른 좌표를 보내왔으므로 그 지점부터 시작하도록.
						if x1 == 1 && y1 == 1 && x2 == 3 && y2 == 3 && i == depth {
							x1, y1, x2, y2 = startX, startY, endX, endY
						}

						// 이미 사용한 양파깡 판이므로 사용 불가함
						for x := x1; x <= x2; x++ {
							for y := y1; y <= y2; y++ {
								if visited[x][y] {
									break ContinueLoop
								}
							}
						}

						flavor := calcOnion(x1, x2, y1, y2, prefixSum, board)
						if flavor > max {
							checkFlag = true
							max = flavor
							a, b, c, d = x1, y1, x2, y2
						} else if flavor == maxTaste { // 최고 맛이 중복인 경우

							argVisited := make([][]bool, N+1)
							for i, val := range visited {
								argVisited[i] = make([]bool, 0)
								argVisited[i] = append(argVisited[i], val...)
							}
							// 중복 탐색 배제
							visitMarker(x1, y1, x2, y2, argVisited)
							tempResult := sb.String()

							resultSave(max, x1, y1, x2, y2)
							res := findOnionGGang(x1, y1, x2, y2, i+1, argVisited)
							if res {
								return true
							}
							sb.WriteString(tempResult)
						}

					}
				}
			}
		}

		if checkFlag {
			// 직사각형 테두리 사용 체크
			visitMarker(a, b, c, d, visited)

			resultSave(max, a, b, c, d)
		} else {
			sb.Reset()
			return false
		}
	}

	return true
}

func visitMarker(a, b, c, d int, visited [][]bool) {
	for x := a; x <= c; x++ {
		for y := b; y <= d; y++ {
			if x == a || x == c {
				visited[x][y] = true
			}
			if y == b || y == d {
				visited[x][y] = true
			}
		}
	}
}

func resultSave(max, a, b, c, d int) {
	result := fmt.Sprintf("%d %d %d %d %d\n", max, a, b, c, d)
	sb.WriteString(result)
}

func findMaxTaste(visited [][]bool) int {
	max := -999999999
	for x1 := 1; x1 <= N-2; x1++ {
		for y1 := 1; y1 <= N-2; y1++ {
			// ContinueLoop:
			for x2 := x1 + 2; x2 <= N; x2++ {
			ContinueLoop: // Label을 한칸 위로 했어서 x2가 올라가는게 아니라 y1부터 올라갔기때문에 너무 과하게 스킵했었음
				for y2 := y1 + 2; y2 <= N; y2++ {
					// 이미 사용한 양파깡 판이므로 사용 불가함
					for x := x1; x <= x2; x++ {
						for y := y1; y <= y2; y++ {
							if visited[x][y] {
								break ContinueLoop
							}
						}
					}

					flavor := calcOnion(x1, x2, y1, y2, prefixSum, board)
					if flavor > max {
						max = flavor
					}
				}
			}
		}
	}

	return max
}
