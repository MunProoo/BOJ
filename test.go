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

const (
	target = "SPGED"
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

	var input interface{}
	input = readLineInt()

	N := input.([]int)[0]
	board := make([]string, N)

	impossible := false // 불가능한 경우 체크

	for i := 0; i < N; i++ {
		input = readLine()
		word := input.(string)

		if strings.Count(word, "G") == 4 && strings.Count(word, "Y") == 1 {
			impossible = true
		}
		board[i] = word
	}

	// 불가능한 조건
	if impossible {
		writer.WriteString("IMPOSSIBLE\n")
		return
	}

	writer.WriteString(target + "\n")

	// SPEED처럼 E가 2번 나오는걸 target으로 잡지 않고 더 쉬운걸 target으로 해도 되는지 테스트
	for _, word := range board {
		visited = make([]bool, 5) // G, R로 매칭되면 true로 폐기하도록

		answer := findPossibleAnswer(word)

		writer.WriteString(answer + "\n")
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

func readLine() string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

func findPossibleAnswer(word string) string {
	answer := make([]byte, 5)

	findGreen(word, &answer)
	findYellow(word, &answer)
	findGrey(word, &answer)
	return string(answer)
}

// Green이면 answer에 매칭
func findGreen(word string, answer *[]byte) {
	for i, char := range word {
		if char == 'G' {
			(*answer)[i] = target[i]
			visited[i] = true
		}
	}
}

// Yellow면 answer에 가능한 글자 매칭
func findYellow(word string, answer *[]byte) {
	for i, char := range word {
		if char == 'Y' {
			idx := findNoMatchIdx(i) // (Green에서 매칭된거 제외하고 매칭)
			(*answer)[i] = target[idx]
			visited[idx] = true
		}
	}
}

// Grey면 무조건 "F"로 매칭
func findGrey(word string, answer *[]byte) {
	for i, char := range word {
		if char == 'B' {
			(*answer)[i] = 'F'
		}
	}
}

// 사용 안한 글자면서 정답과 위치가 다른 글자 찾기
func findNoMatchIdx(i int) int {

	// 일단 본인 자리보다 뒤에부터 매칭해가도록 함
	// cause : (뒤쪽에서 본인자리랑 같은 글자가 매칭될수도 있어서)
	for idx := i + 1; idx < 5; idx++ {
		if !visited[idx] {
			return idx
		}
	}

	// 뒤에 매칭될 수 있는 글자가 없으면, 앞에서부터 검사
	for idx := range visited {
		if !visited[idx] {
			return idx
		}
	}

	return 99 // 문제 조건에 맞게 들어오면 절대 안탐. 탔으면 런타임에러 (생각 못한 반례가 있다는 것)

}
