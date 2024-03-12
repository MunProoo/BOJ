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

	// 풀이변수
)

func main() {

	/*
		11694 님게임

		* 님게임 필승전략 *
		요약 : 자기차례에, 각 돌 무더기에 있는 돌의 개수를 XOR한 값을 0으로 만들어주면 필승

		설명 : 무더기가 있을 때 각 무더기의 돌 수를 똑같이 맞춰주면 다음턴이 아무리 애를 써도 마지막 돌은 내 차례에 집을 수 있음
			   똑같이 맞추는 작업을 XOR을 통해 손쉽게 확인할 수 있음.

		-> Nim-sum : 각 더미의 돌 개수를 전부 XOR 연산한 값. ( nimsum = a ^ b ^ c )
		-> Nim-sum == 0 인 상태면 패배가 결정난 상태
		내 차례에 Nim-sum을 0으로 만들어 상대방에게 보내면 승리






		* 스프라그 그런디 (Sprague-Grundy) 정리 *
		"어떤 게임을 Nim 게임으로 환원시켜 필승법을 알아내는 방법"
		그런디 수 (Grundy Number)라는 걸 두는데, 이 수가 Nim 게임의 돌 더미에 대응된다.

		현재 그런디 수 = mex (Grundy Numbers...) , (mex함수는 Grundy Numbers에 속하지 않은 가장 작은 0 이상의 정수이다.)




	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()
	N := input.([]int)[0] // N
	input = readLineInt()

	var xor int = 1
	for i := 0; i < N; i++ {
		stones := input.([]int)[i]
		xor ^= stones
	}

	if xor != 0 {
		sb.WriteString("koosaga")
	} else {
		sb.WriteString("cubelover")
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
