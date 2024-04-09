package main

import (
	"bufio"
	"fmt"
	"sort"
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
		1461. 도서관
		도서관 책 정리할 때 최소 걸음으로 정리하도록 계산하는 프로그램

		가장 먼 거리는 마지막에 가져간다. (2배 X)
		가장 먼 거리부터 순차적으로 덩어리로 묶는다.

		조건을 나눠야 하나? 양수만 있는경우, 음수만 있는 경우, 두 조건이 같이 있는 경우

		(1)
		음수일때는 밑에서부터 가면 되고,
		양수일때는 위에서부터 묶어야함

		(2)
		음수가 가장 크다면, 처음값만 걸음 *1
		양수가 가장 크다면, 마지막 값만 걸음*1

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()
	N, M = input.([]int)[0], input.([]int)[1]
	input = readLineInt()
	Location := input.([]int)
	sort.Ints(Location)

	MinusIsBiggest := false
	if abs(Location[0]) > abs(Location[N-1]) {
		// 음수값이 가장 크면
		MinusIsBiggest = true
	}

	changeSignIndex := 0 // 책의 위치가 -에서 +로 바뀌는 순간.

	for i, val := range Location {
		if val > 0 {
			changeSignIndex = i
			break
		}
	}

	걸음수 := 0
	// (1) 책의 위치가 음수인 경우일 때 계산
	// changeSignIndex : 음수 -> 양수가 되는 구간 체크되면, 종료
	// 해당 구간에서 끊고 다시 덩어리를 묶어 가야한다.
	for i := 0; i < N; i += M {
		if Location[i] > 0 {
			// 음수일때만 계산할거니까 양수만 있으면 pass
			break
		}

		signFlag := false // 음수 -> 양수로 전환됐으면 반복문 탈출

		tempLange := i + M // 덩어리 나눌 범위
		var tempLocation []int

		// 배열의 끝이라면 배열의 마지막 인덱스로 범위 조정
		if tempLange > N {
			tempLange = N
		}

		// 임시 슬라이스에 값 할당 (덩어리 생성)
		tempLocation = Location[i:tempLange]

		// 책의 위치가 음수 -> 양수로 전환되었다면 덩어리를 새롭게 구성해야함
		if i <= changeSignIndex && tempLange >= changeSignIndex && changeSignIndex != 0 {
			tempLocation = Location[i:changeSignIndex]

			// 음수에서 마지막 걸음을 해야하는데, 이 곳에 빠져서 걸음을 2배하는 경우가 생김
			// maxStep := max(abs(tempLocation[0]), abs(tempLocation[len(tempLocation)-1]))
			// 걸음수 += maxStep * 2

			// 이 공간은 그저 음수 -> 양수 전환 시 덩어리 범위 조정만 하도록 변경
			signFlag = true
		}

		// tempLocation = Location[tempLange:i]

		// 걸음수 계산
		if i == 0 && MinusIsBiggest {
			// 음수가 가장 클 때, 마지막 책을 갖다 놓는 작업 -> 2번 걸을 필요X
			maxStep := abs(tempLocation[0])
			걸음수 += maxStep
			if signFlag { // 남은 값들은 양수이므로 반복문 더 돌 필요X
				break
			}
		} else {
			// 일반적인 작업
			maxStep := max(abs(tempLocation[0]), abs(tempLocation[len(tempLocation)-1]))
			걸음수 += maxStep * 2
		}
	}

	// (2) 책의 위치가 양수일 때 계산
	for i := N; i > changeSignIndex; i -= M {
		if Location[i-1] < 0 {
			// 양수일 때만 계산할 거니까 음수만 있으면 pass
			break
		}

		tempLange := i - M // 덩어리 나눌 범위
		var tempLocation []int

		// 들고갈 수 있는 책이 전체 책 개수보다 많다면, 전부 들고 가도록 조정
		if tempLange < 0 {
			tempLange = 0
		}

		// 임시 슬라이스에 값 할당 (덩어리 생성)
		tempLocation = Location[tempLange:i]

		// 책의 위치가 음수 -> 양수로 전환되는 구간이 있을 수 있으므로, 해당 구간 배제
		if i >= changeSignIndex && tempLange <= changeSignIndex && changeSignIndex != 0 {
			tempLocation = Location[changeSignIndex:i]
		}

		// 걸음수 계산
		if i == N && !MinusIsBiggest {
			// 양수가 가장 클 때, 마지막 책을 갖다 놓는 작업 -> 2번 걸을 필요X
			maxStep := abs(tempLocation[len(tempLocation)-1])
			걸음수 += maxStep
		} else {
			// 일반적인 작업
			maxStep := max(abs(tempLocation[0]), abs(tempLocation[len(tempLocation)-1]))
			걸음수 += maxStep * 2
		}
	}

	fmt.Println(걸음수)
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
