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
		14254 비밀번호 변경

		규칙 : 처음 K개 글자는 마지막 K개의 글자와 같아야 한다
		출력 : 규칙을 만족하는 글자의 최소 개수

		※ 자리를 바꾸는게 아닌, 글자를 바꾸는 것

		amavckdkz
		7
		:5

		a m a v c k d k z
		7개가 같아야함

		0 1 2 3 4 5 6
		a m a v c k d [:7]
		a v c k d k z [9-7:]
		2 3 4 5 6 7 8

		현재 같은 애들 : [0]과 [2] , [5]와 [7] 은 안바꾼다 치면

		[3]자리 v -> m       [3]변경 (+1)
		a m a m c k d
		a m c k d k z

		[4]자리 c -> a       [4]변경 (+2)
		a m a m a k d
		a m a k d k z

		[5]자리 안바꾸기로 함. 그렇다면 [3]을 다른애로 변경      [3]변경 m->k
		a m a k a k d
		a k a k d k z

		[6]자리 d -> a     (+3)
		a m a k a k a
		a k a k a k z

		[8]자리 z -> a (+4)
		a m a k a k a
		a k a k a k a

		아직 동일하지 않으니 다시 처음부터.
		[1]자리 m -> k    (+5)
		a k a k a k a
		a k a k a k a
		이 부분에서 [3]자리가 아닌, [1]자리를 변경하는 방법은 ..?

		최종 변경본 : a k a k a k a k a
		변경한 글자 : [1], [3], [4], [6], [8]


		글자를 변경하는 횟수가 아닌, 바꾸는 글자의 개수를 세는 것.

		기본적으로 뒤에 문자열을 변경하는 것이 순차적으로 갈 때 막힘이 없다.
		단, 기존에 같았던 문자열이라면, 변경 없이 가도록 한다.
		-> 기존에 앞에서 변경한 문자를 다른 것으로 변경해준다

		계속 틀리는데, 이유를 못찾겠음. 다른 로직 생각해야함

		다른 로직)
		문자를 1개 바꿨을 때, 그 영향을 계산할 수 있다..
		0 1 2 3 4 5 6
		a m a v c k d [:7]
		a v c k d k z [9-7:]
		2 3 4 5 6 7 8

		일 때, 인덱스 3번을 v -> m 으로 바꾸면, 인덱스 5번, 7번에 영향을 준다.
		5번 7번이 동일하므로 5번 7번에 맞춰 3번을 바꿔야 한다면, 인덱스 1번에 영향을 줘야한다...

		-> 구간 length =  len(word) - K
		첫 시작이 되는 문자에서 각각을 영향주는 문자 중 많이 나오는 애로 전부 변경한다. ?

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	word := readLine()
	K := readLineInt()[0]

	length := len(word) - K // 비밀번호 뒤에서부터 셌을 때 시작하는 인덱스

	// 계속 값을 변경해줘야하는데,
	// string은 인덱스로 변경이 불가능하니까 rune타입 배열로 작업
	var arr []rune
	for _, val := range word {
		arr = append(arr, val)
	}

	result := 0
	// 1개를 바꾸면 영향이 오는 건 length 간격으로 감.
	// 즉 length 간격에 있는 문자 중 가장 많은 것으로 바꾸면 효율적인 변경이다
	for i := 0; i < length; i++ {
		// 기준 := arr[:K]
		// 변경될것 := arr[length:]

		// 가장 많은 문자 찾기
		charMap := make(map[rune]int)
		for j := i; j < len(arr); j += length {
			// 구간의 각 문자의 수 저장
			charMap[arr[j]]++
		}

		max := 0 // 문자의 최대 수
		keyChar := rune(0)
		for key, val := range charMap {
			if max < val {
				max = val
				keyChar = key
			}
		}

		for j := i; j < len(arr); j += length {
			if arr[j] != keyChar {
				arr[j] = keyChar
				result++
			}
		}

	}

	fmt.Println(result)
	// fmt.Println(string(arr))
}

func readLine() string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
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
