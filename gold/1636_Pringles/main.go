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
	N      int
	result map[int][]int
)

func main() {

	/*
		1636 한번 열면 멈출 수 없어
		그리디 알고리즘. 애드혹

		프링글스를 먹으면 중독된다.
		중독스트레스가 변경될때마다 수명이 줄어든다.
		줄어드는 수명을 최소로 하자

		맨 처음에 프링글스 먹을 때 어떤 스트레스값을 고를지 고민하다가
		모든 경우 다 돌려서 가장 적은 수명 나오는거로 선택하도록..

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()
	N = input.([]int)[0] // N
	stress := make([][]int, N)

	presentStress := 0 // 현재 스트레스
	result = make(map[int][]int)

	for i := 0; i < N; i++ {

		input = readLineInt()
		// min, max := input.([]int)[0], input.([]int)[1]
		stress[i] = input.([]int)
	}

	var minLifeSpan int = 9000000000
	for i := stress[0][0]; i <= stress[0][1]; i++ {
		presentStress = i
		val := calcLifeSpan(stress, presentStress)
		if minLifeSpan > val {
			minLifeSpan = val
		}
	}

	firstMin, firstMax := stress[0][0], stress[0][1]
	nextMin, nextMax := stress[1][0], stress[1][1]
	if nextMin < 

	최소수명 := fmt.Sprintf("%d\n", minLifeSpan)
	sb.WriteString(최소수명)

	for _, val := range result[minLifeSpan] {
		sb.WriteString(fmt.Sprintf("%d\n", val))
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

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func calcLifeSpan(stress [][]int, presentStress int) int {
	lifespan := 0
	stressList := make([]int, 0)
	stressList = append(stressList, presentStress)

	for i := 1; i < N; i++ {
		min, max := stress[i][0], stress[i][1]
		if min <= presentStress && max >= presentStress {
			// 중독 스트레스를 같은 값으로 지정
			lifespan += 0
		} else {
			diffMin := abs(presentStress - min)
			diffMax := abs(presentStress - max)
			// Min, Max중 수명이 덜 줄어드는 쪽으로 변경
			if diffMin < diffMax {
				presentStress = min
				lifespan += diffMin
			} else {
				presentStress = max
				lifespan += diffMax
			}

		}
		stressList = append(stressList, presentStress)
	}

	// 결과물 저장
	result[lifespan] = stressList
	return lifespan
}
