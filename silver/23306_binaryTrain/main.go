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
)

func main() {

	/*
		23306. 비내리는 호남선
		인터랙티브 문제.

		철로의 길이 N이 주어진다.
		0은 저지대, 1은 고지대
		01이면 오르막 , 10이면 내리막

		질문 : k번째 철로가 0인지 1인지
		답 : 0 or 1

		오르막이 많으면 1 내리막이 많으면 -1, 같으면 0 출력

		※  log2N보다 많이 질문하면 안된다.
			8<=N<=2048 (질문 수 : 3~11번)

		N이 짝수면 맨 앞과 맨 뒤만 물어봐도 될듯 0,1인 경우 1 , 1,0인 경우 0 , 둘이 같은 경우는 똑같 !
		N이 짝수면 오르막과 내리막이 "같지 않은 경우" 항상 1씩 차이남.

		N이 홀수면?
		N-1도 물어봐서 N번 자리와 비교해서 오르막, 내리막 추가

		69퍼? 까지 갔다가 틀림..
		구현을 잘못해서였음
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input := readLineInt()
	N := input[0] // 철로의 길이

	trainRoad := make([]int, 3) // 철로의 1번째, N-1번째, N번째 지대 저장
	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			fmt.Fprintf(os.Stdout, "? %d\n", 1)
		case 1:
			fmt.Fprintf(os.Stdout, "? %d\n", N-1)
		case 2:
			fmt.Fprintf(os.Stdout, "? %d\n", N)
		}
		answer := readLineInt()[0]
		trainRoad[i] = answer
	}

	first, last := 0, 0
	landFlag := ""

	if N%2 == 0 {
		// N이 짝수면 1번째와 N번째만 비교
		first, last = trainRoad[0], trainRoad[2]
	} else {
		// N이 홀수면 1번째와 N-1번째 먼저 비교
		first, last = trainRoad[0], trainRoad[1]
	}

	if first > last {
		// 1 ~~ 0 이므로 내리막이 더 많음
		landFlag = "내리막"
	} else if first < last {
		// 0 ~~ 1 이므로 오르막이 더 많음
		landFlag = "오르막"
	} else {
		// 평지임
		landFlag = "같음"
	}

	if N%2 != 0 {
		realLast := trainRoad[2]

		// last가 realLast보다 크려면 last가 1인거고. last가 1이면 오르막 혹은 같음임.
		if last > realLast {
			// 마지막 2개가 10 이므로 내리막 추가
			copyLandFlag := landFlag
			switch copyLandFlag {
			case "오르막":
				landFlag = "같음"
			case "같음":
				landFlag = "내리막"
			}

		} else if last < realLast {
			// last가 realLast보다 크려면 last가 0인거고. last가 0이면 내리막 혹은 같음임

			// 마지막 2개 01 이므로 오르막 추가
			copyLandFlag := landFlag
			switch copyLandFlag {
			case "내리막":
				landFlag = "같음"
			case "같음":
				landFlag = "오르막"
			}
		}
	}

	var result int
	switch landFlag {
	case "오르막":
		result = 1
	case "내리막":
		result = -1
	case "같음":
		result = 0
	}

	fmt.Fprintf(os.Stdout, "! %d\n", result)
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
