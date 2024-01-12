package main

import (
	"bufio"
	"strconv"
	"strings"

	"os"
)

var (
	writer  *bufio.Writer
	reader  *bufio.Reader
	visited []bool
)

func main() {

	/*
		26147 wordle
		초: G (71)
		노: Y (89)
		회: B (66)

		IMPOSSIBLE : GGGGY (Y가 1개고, 전부 G인 경우)

		YYYYY 면 target에 대한 모든 문자가 나와야함. 순서만 다른
		먼저 G를 체크하고 매칭 안된 문자들 중에서 Y를 체크해야함.!!

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// var input interface{}

	writer.WriteString("Hymn To Love")
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
