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
	// defaultArr []byte
	sb             strings.Builder
	defaultArr     []byte
	defaultPointer Pointer
	commands       []string // brainfuck 프로그램
	inputs         string   // brainfuck에 입력될 문자
	// inputs         []rune   // brainfuck에 입력될 문자

	arrSize int // defaultArr 크기
)

type Pointer int

func (p *Pointer) Increase() {
	if int(*p) == arrSize {
		*p = 0
		return
	}
	*p++
}
func (p *Pointer) Decrease() {
	if *p == 0 {
		*p = 32767
		return
	}
	*p--
}

func main() {

	/*
		3954 Brainfuck 인터프리터

		, : 문자 하나를 읽고 포인터가 가리키는 곳에 저장한다. (입력의 마지막이면 255를 저장한다)

		Brainfuck 프로그램이 무한루프인지 아닌지만 확인한다.
		(5천만번을 반복해야 끝나는 루프는 주어지지 않는다)

		풀이 1.
		(1)무한루프가 된다는 것은 포인터가 가리키는 값이 0이 안된다는 것.
			- 그냥 단순히 루프를 통해 0과의 거리가 멀어지는지만 체크했는데, 오버플로우가 있는걸 생각했어야함
			- 루프의 시작값과 루프 후 값의 변화량이 중요

			1
			1 9 1
			++[[++]+]
			a


			이런 경우라면, 안쪽 루프가 1번 반복하고 무한루프에 빠짐.
			01, 23, 45, 67, 80     시작:2, diff : 2
			12, 34, 56, 78, 01     시작:3, diff : 2

			(arrSize-1-시작 ) % diff == 0이 아니라면, 무한루프다.

		... 위의 식과 다른데, EOF에 의해 현재 값이 바뀌는 경우가 있음.
		변수가 brainfuck 입력값이므로 입력한 값을 다 쓰고나서 수식을 사용하도록 구현하면?
		-> 제출하니 시간초과 (9%)


	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	// 테스트케이스
	input = readLineInt()
	T := input.([]int)[0]

	// commands = make([]string, sc)
	// inputs = make([]rune, si) // brainfuck 입력

	// 명령어 Input
	for i := 0; i < T; i++ {
		// 입력1
		input = readLineInt()
		ints := input.([]int)
		arrSize, _, _ = ints[0], ints[1], ints[2]

		defaultArr = make([]byte, arrSize)
		defaultPointer = 0

		// 입력2 (프로그램)
		input = readLine()
		command := input.(string)

		// 입력3 (프로그램에 입력)
		input = readLine()
		inputs = input.(string)

		// 대괄호 매핑
		bracketMap := mappingBrackets(command)

		// brainfuck 해석
		interpreter(command, bracketMap)

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

func interpreter(command string, bracketMap map[int]int) {
	inputCnt := 0 // brainfuck 입력 사용 개수
	inputLength := len(inputs)
	loopFlag := false
	var initValue byte

	for i := 0; i < len(command); i++ {
		switch command[i] {
		case '>':
			defaultPointer.Increase()
		case '<':
			defaultPointer.Decrease()
		case '+':
			defaultArr[defaultPointer]++
		case '-':
			defaultArr[defaultPointer]--
		case '.':
			// sb.WriteString(string(defaultArr[defaultPointer]))
		// 루프 작업
		case '[':

			if defaultArr[defaultPointer] == 0 { // 루프가 끝남
				i = bracketMap[i] - 1 // i++ 되므로
				// loopFlag = false
				continue
			}

		case ']':
			if defaultArr[defaultPointer] != 0 {
				i = bracketMap[i]

				// ,로 인한 EOF라는 변수가 있어서 수식으로 계산. 무조건 루프 돌려야함
				if strings.Contains(command, ",") {
					continue
				}

				// 이미 한 번은 돌고 왔음
				if loopFlag {

					// byte는 255에서 오버플로우가 나는걸 생각 못함..
					diff := abs(int(initValue) - int(defaultArr[defaultPointer]))
					if diff == 0 {
						end := bracketMap[i]
						start := bracketMap[end]
						sb.WriteString(fmt.Sprintf("Loops %d %d\n", start, end))
						return
					}

					expression := (arrSize - 1 - int(initValue)) % int(diff)
					// fmt.Println(expression)

					// 오버,언더 플로우 후에 0을 지나침
					if expression != 0 {
						end := bracketMap[i]
						start := bracketMap[end]
						sb.WriteString(fmt.Sprintf("Loops %d %d\n", start, end))
						return
					}
				}

				loopFlag = true
				initValue = defaultArr[defaultPointer]
			} else {
				loopFlag = false
			}

			// 루프 종료
		case ',':
			if inputCnt < inputLength {
				defaultArr[defaultPointer] = inputs[inputCnt]
			} else {
				defaultArr[defaultPointer] = 255
			}
			// 입력 사용한 것 체크를 뺴먹으니 애먹은 곳에서 무한루프 돌았던 삽질..
			inputCnt++
		}
	}
	sb.WriteString("Terminates\n")
}

// 대괄호 짝맺기
func mappingBrackets(s string) map[int]int {
	stack := []int{}
	bracketMap := make(map[int]int)

	for idx, char := range s {
		if char == '[' {
			stack = append(stack, idx)
		} else if char == ']' {
			start := stack[len(stack)-1]
			end := idx
			bracketMap[start] = end
			bracketMap[end] = start
			// 짝이 맞아 스택에서 '['를 제거
			stack = stack[:len(stack)-1]
		}
	}

	return bracketMap
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
