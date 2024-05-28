package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"os"
)

var (
	writer                  *bufio.Writer
	reader                  *bufio.Reader
	sb1, sb3, sb4           strings.Builder
	initialBomb, secondBomb [][]bool
)

func main() {

	/*
		16919 봄버맨2

		직사각형 격자판이 있다. 비었거나 폭탄이 들어있다.

		같은 모양 반복함
		0,1 초기상태
		2 모든칸
		3 초기 폭탄 폭발모습
		4 모든칸
		5 초기상태
		6 모든칸
		7 초기 폭탄 폭발모습
		8 모든칸
		9 초기상태

		N이 짝수 : 모든칸 폭탄
		N이 홀수 : 초기 상태 or 초기 폭탄 폭발 모습

		틀림1 첫 번째 폭발로 모든 땅이 .이 되버리는 경우
		초기상태가 아닐 수도 있다.

		틀림2 직접 계산하니까 시간초과

		1번만 반복하고, 즉시 제출하자
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input := readLineInt()
	R, C, N := input[0], input[1], input[2]

	initialBomb = make([][]bool, R) // 첫 번째 폭탄
	secondBomb = make([][]bool, R)  // 2초에 설치하는 폭탄
	for i := 0; i < R; i++ {
		initialBomb[i] = make([]bool, C)
		secondBomb[i] = make([]bool, C)
	}

	pan := make([][]string, R)
	for i := 0; i < R; i++ {
		pan[i] = make([]string, C)
		input := strings.Split(readLine(), "")
		for j, val := range input {
			pan[i][j] = val
			if val == "O" {
				initialBomb[i][j] = true
			}
		}
	}

	if N == 1 {
		// N이 1인 경우와 5,9,13 번째의 판모양이 달라지는 경우가 있으므로 따로 처리해야했음
		savePan(pan, 1)
		fmt.Println(sb1.String())
		return
	}

	for i := 2; i <= 8; i++ {
		if i%4 == 0 {
			// 남은 칸 모두에 폭탄설치 및 기록(1번째 폭탄)
			BombPlanting(pan, 1)
			savePan(pan, 1)
		} else if i%4 == 2 {
			// 남은 칸 모두에 폭탄설치 및 기록(2번째 폭탄)
			BombPlanting(pan, 2)
			savePan(pan, 1)
		} else if i%4 == 3 {
			// 1번째 폭탄 터진 후 반영
			BombBomb(pan, initialBomb, 1)
			savePan(pan, 3)
		} else if i%4 == 1 {
			// 2번째 폭탄 터진 후 반영
			BombBomb(pan, secondBomb, 2)
			savePan(pan, 4)
		}
	}

	if N%2 == 0 {
		fmt.Println(sb1.String())
	} else if N%4 == 3 {
		fmt.Println(sb3.String())
	} else if N%4 == 1 {
		fmt.Println(sb4.String())
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

// 모든 땅에 폭탄 설치
func BombPlanting(pan [][]string, order int) {
	R := len(pan)
	C := len((pan)[0])

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if (pan)[i][j] != "O" {
				(pan)[i][j] = "O"

				if order == 1 {
					initialBomb[i][j] = true
				} else {
					secondBomb[i][j] = true
				}
			}
		}
	}

}

// 폭탄 터짐
func BombBomb(pan [][]string, bomb [][]bool, order int) {
	R := len(pan)
	C := len((pan)[0])

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if bomb[i][j] {
				(pan)[i][j] = "."

				상, 하, 좌, 우 := i-1, i+1, j-1, j+1
				if 상 <= 0 {
					상 = 0
				}
				if 하 >= R {
					하 = R - 1
				}
				if 좌 <= 0 {
					좌 = 0
				}
				if 우 >= C {
					우 = C - 1
				}

				(pan)[상][j] = "."
				(pan)[하][j] = "."
				(pan)[i][좌] = "."
				(pan)[i][우] = "."

				if order == 1 { // 터진 후에 영향 반영
					secondBomb[상][j] = false
					secondBomb[하][j] = false
					secondBomb[i][좌] = false
					secondBomb[i][우] = false
				} else {
					initialBomb[상][j] = false
					initialBomb[하][j] = false
					initialBomb[i][좌] = false
					initialBomb[i][우] = false
				}
			}
		}
	}

}

// 현재 판을 기록
func savePan(pan [][]string, order int) {
	var answersb strings.Builder
	for _, val := range pan {
		for _, val2 := range val {
			answersb.WriteString(val2)
		}
		answersb.WriteString("\n")
	}

	if order == 1 {
		sb1 = answersb
	} else if order == 3 {
		sb3 = answersb
	} else {
		sb4 = answersb
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

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
