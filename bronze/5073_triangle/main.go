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
	sb     strings.Builder
)

func main() {

	/*
		5073. 삼각형과 세 변

		세 변의 길이가 모두 같으면 Equilateral
		두 변의 길이만 같으면 Isosceles
		세 변의 길이가 모두 다르면 Scalene
		삼각형 조건을 만족 못하면 Invalid

		삼각형 조건 : 두 변의 길이 합 > 한 변의 길이
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		input := readLineInt()
		a, b, c := input[0], input[1], input[2]

		if a == 0 && b == 0 && c == 0 {
			break
		}

		if a+b <= c || a+c <= b || b+c <= a {
			sb.WriteString("Invalid\n")
		} else if a == b && b == c {
			sb.WriteString("Equilateral\n")
		} else if a == b || a == c || b == c {
			sb.WriteString("Isosceles\n")
		} else {
			sb.WriteString("Scalene\n")
		}

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
