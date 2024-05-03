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
		27943. 가지 사진 찾기
		인터랙티브 문제.

		3 ~ 10^18개의 사진이 있다.
		가지는 절반 이상이다.

		>가지가 몇번부터 몇번까지인지 찾아내기
		>최대 120번 질문

		풀이 :
		(1) 중간은 무조건 gaji니까 앞뒤를 나눠서 처음과 중간, 중간과 끝을 start, end로 둔다
		(2) 앞범위 :
		(2-1) mid값이 gaji면 end를 mid로 변경
		(2-2) mid값이 != gaji면 start를 mid로 변경
		(2-3) start == end가 될때까지 반복?

		log2(10^9) ~= 30 이므로 횟수는 충분함

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input := readLineInt()
	N := input[0] // 총 사진 수

	// gaji 이름 알아내기 (가운데는 무조건 가지임)
	fmt.Fprintf(os.Stdout, "? %d\n", (1+N)/2)
	gaji := readLine()

	start := 1
	end := N / 2

	frontMid := 0

	// 앞구간 체크
	// 3~7이라 할때 2도 체크해야하는데 안하네? start==end여도 하도록 해야함. start > end면 break하고
	for {
		frontMid = (start + end) / 2

		if start > end { // 찾아냄
			frontMid += 1 // 가지의 시작은 +1부터
			break
		}

		fmt.Fprintf(os.Stdout, "? %d\n", frontMid)
		photo := readLine()
		if strings.EqualFold(gaji, photo) {
			end = frontMid - 1
		} else {
			start = frontMid + 1
		}
	}

	start = N / 2
	end = N
	backMid := 0
	for {
		backMid = (start + end) / 2

		if start > end {
			// backMid += 1 // 가지의 끝은 -1까지
			break
		}
		fmt.Fprintf(os.Stdout, "? %d\n", backMid)
		photo := readLine()

		if strings.EqualFold(gaji, photo) {
			start = backMid + 1
		} else {
			end = backMid - 1
		}
	}

	fmt.Fprintf(os.Stdout, "! %d %d\n", frontMid, backMid)
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
