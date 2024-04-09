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
	// result [][]int // map보단 슬라이스가 더 빠르겠지라 생각했지만 고려해야하는 인덱스가 2천만개면 얘기가 달라지지
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

		----------------- 좀 더 그리디하게 -------------------
		맨 처음 스트레스 수치 고르기
		* 현재 고를 수 있는 스트레스 범위 : A ~ B
		* 다음에 고를 스트레스 범위 : C ~ D

		조건1. A ~ B와 C ~ D의 범위가 겹치는 부분이 있다.
		-> 겹치는 부분 중 어떤걸 선택하면 좋을 지는 다 돌려서 확인

		조건2. A ~ B와 C ~ D가 안겹침
		-> B와 C가 가까우면 B를, A와 D가 가까우면 D가 골라질 것인데. 이 후의 과정에서 최소 수명이 달라질 수 있으므로 어쩄든 A,B 2가지를 돌려서 확인
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

	// 좀 더 그리디한 풀이
	firstMin, firstMax := stress[0][0], stress[0][1]
	nextMin, nextMax := stress[1][0], stress[1][1]

	// 겹치는 부분의 범위 체크
	nestedMin, nestedMax, isNested := findNestedRange(firstMin, firstMax, nextMin, nextMax)
	var minLifeSpan int = 9000000000
	if isNested { // 겹치는 부분이 있다면 그 부분만 확인
		for i := nestedMin; i <= nestedMax; i++ {
			presentStress = i
			val := calcLifeSpan(stress, presentStress)
			if minLifeSpan > val {
				minLifeSpan = val
			}
		}
	} else { // 겹치는 부분 없으면 현재 스트레스 둘 중 하나
		for i := 0; i < 2; i++ {
			presentStress = stress[0][i] // 처음은 최소, 두번짼 최대
			val := calcLifeSpan(stress, presentStress)
			if minLifeSpan > val {
				minLifeSpan = val
			}
		}
	}

	/*
		var minLifeSpan int = 9000000000
		for i := stress[0][0]; i <= stress[0][1]; i++ {
			presentStress = i
			val := calcLifeSpan(stress, presentStress)
			if minLifeSpan > val {
				minLifeSpan = val
			}
		}
	*/

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

// 겹치는 범위 찾는 함수
// return : nestedMin, nestedMax , bool(겹치는지 여부)
func findNestedRange(min1, max1, min2, max2 int) (int, int, bool) {
	nestedMin := max(min1, min2)
	nestedMax := min(max1, max2)

	if nestedMin <= nestedMax {
		return nestedMin, nestedMax, true
	}

	return 0, 0, false
}
