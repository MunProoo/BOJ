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

	// 풀이 변수
	maze [][]bool
)

const (
	maxCount = 101
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type Direction int

// 명령에 따라 방향 전환
func (d *Direction) Turn(direction rune) {
	switch direction {
	case 'R':
		(*d) += RIGHT
	case 'L':
		(*d) += LEFT
	}

	(*d) = (*d) % 4
}

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

		방향)
		UP : 0
		RIGHT : 1
		DOWN : 2
		LEFT : 3
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

	maze = make([][]bool, maxCount)
	for i := 0; i < maxCount; i++ {
		maze[i] = make([]bool, maxCount)
	}

	x, y := 50, 50    // 좌표
	maze[x][y] = true // 현재 위치도 길이니깐

	// 노트에 적힌대로 이동
	moveToNote(x, y, note)

	// 이동한 경로만 포함하는 미로 구역 구하기
	minX, minY, maxX, maxY := findRectangle()

	// 그려내기
	printMaze(minX, minY, maxX, maxY)

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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 노트에 맞춰 이동
func moveToNote(x, y int, note string) {
	var direction Direction = DOWN // 방향
	for _, char := range note {
		// 방향 전환
		if char != 'F' {
			direction.Turn(char)
			continue
		}

		// 앞으로 이동
		switch direction {
		case DOWN:
			x += 1
		case RIGHT:
			y += 1
		case UP:
			x -= 1
		case LEFT:
			y -= 1
		}

		maze[x][y] = true
	}
}

// 이동한 경로만 포함하는 미로 구역 구하기
func findRectangle() (minX, minY, maxX, maxY int) {
	minX, minY, maxX, maxY = 111, 111, 0, 0
	for i := 0; i < maxCount; i++ {
		for j := 0; j < maxCount; j++ {
			if maze[i][j] {
				minX = min(minX, i)
				minY = min(minY, j)
				maxX = max(maxX, i)
				maxY = max(maxY, j)
			}
		}
	}
	return
}

// 그려내기
func printMaze(minX, minY, maxX, maxY int) {
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if maze[i][j] {
				fmt.Fprintf(writer, "%s", ".")
			} else {
				fmt.Fprintf(writer, "%s", "#")
			}
		}
		writer.WriteString("\n")
	}
}
