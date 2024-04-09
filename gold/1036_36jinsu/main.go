package main

import (
	"bufio"
	"fmt"
	"math/big"
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
)

func main() {

	/*
		1036. 36진수

		문자열에서 가장 많이 중복되는 수를 바꿔야함.
		근데, 자리수가 큰 숫자위주로 바뀌어야함

		합을 최대로 하기 위해서 어떤걸 바꿔야 하나?
		찾아야할 것 -> 문자열의 자리수가 큰 쪽에서 많이 중복되는 수
		가중치를 통해서 가면 좋을듯

		각 자릿수마다 10진수로 변경 시의 기댓값이 있다. (1의 자리면 36^1, 10의 자리면 36^2)
		1. N개의 문자를 순회한다.
		2. differ = 0~Z까지 문자가 나올 때 (Z의 10진수 변환값 - 문자의 10진수 변환값) * 36^자릿수
			-> 각 문자마다 differ를 합산한다.
		3. 0~Z까지 문자에 대해서 differ가 가장 큰 대로 K개를 골라서 변경한다.
		4. 문자를 Z로 변경한 후 36진수로 계산



	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()
	N := input.([]int)[0]

	commands := make([]string, 0)

	charCounts := make([]Pair, 0)
	charMap := make(map[rune]*big.Int)

	for i := 0; i < N; i++ {
		input = readLine()
		command := input.(string)
		commands = append(commands, command)

		pos := 0 // 자릿수
		for j := len(command) - 1; j > 0; j-- {
			// 차이 계산
			differ := calcDiffer(rune(command[j]), pos)

			// 기존 값에 합산
			updateValue(charMap, rune(command[j]), differ)
			pos++
		}
	}

	for char, value := range charMap {
		charCounts = append(charCounts, Pair{char, value})
	}

	// differ 내림차순 정렬
	sort.Slice(charCounts, func(i, j int) bool {
		return charCounts[i].Differ.Cmp(charCounts[j].Differ) > 0
	})

	input = readLineInt()
	K := input.([]int)[0]

	fmt.Println(N, commands, K)
	fmt.Println("")
	fmt.Println(charCounts)

}

type Pair struct {
	Char   rune // 문자
	Differ *big.Int
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

// 36진수를 10진수 big.Int로 변환
func base36ToDecimalBigInt(number string) *big.Int {
	decimal := new(big.Int)
	decimal.SetString(number, 36)
	return decimal
}

// 10진수를 36진수로 변환
func decimalToBase36(decimal *big.Int) string {
	return strings.ToUpper(decimal.Text(36))
}

/*
문자를 숫자로 변경
	0 : 48
	9 : 57
	A : 65
	Z : 90
*/
// "Z" - 특정 문자" 의 차이값

func calcDiffer(char rune, i int) *big.Int {
	base := big.NewInt(36)
	// 10진수로 변환
	num1, num2 := transCodeDesc(char), transCodeDesc('Z')
	differ := big.NewInt(int64(num2) - int64(num1))

	// 자릿수 계산
	exp := new(big.Int).Exp(base, big.NewInt(int64(i)), nil)

	result := new(big.Int).Mul(differ, exp)
	return result

}

// 36진수 1자릿수를 10진수 변환한 값 (int)
func transCodeDesc(char rune) int {
	num := 0
	if char >= '0' && char <= '9' {
		num = int(char - '0')
	} else if char >= 'A' && char <= 'Z' {
		num = int(char - 'A' + 10)
	}
	return num
}

// 맵에 새로운 값 추가 또는 기존 값 업데이트
func updateValue(charCounts map[rune]*big.Int, char rune, value *big.Int) {
	if _, ok := charCounts[char]; ok {
		charCounts[char].Add(charCounts[char], value) // 기존 값에 더하기
	} else {
		charCounts[char] = new(big.Int).Set(value) // 새로운 값 추가
	}
}
