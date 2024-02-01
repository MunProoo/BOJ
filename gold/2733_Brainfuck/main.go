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
	defaultArr     [32768]byte
	defaultPointer Pointer
	commands       []string
)

type Pointer int

func (p *Pointer) Increase() {
	if *p == 32767 {
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
		2733 Brainfuck

		기초 지식
		> : 포인터 증가. 32767이면 오버플로우 (0)
		< : 포인터 감소. 0이면 언더플로우 (32767)
		+ : 포인터가 가리키는 값 증가. 255면 오버플로우 (0)
		- : 포인터가 가리키는 값 감소. 0이면 언더플로우 (255)
		. : 포인터가 가리키는 값을 아스키문자로 출력

		반복문
		[ : 포인터가 가리키는 값이 0이면 ]로
		] : 포인터가 가리키는 값이 0이 아니면 [로

		초기에 0으로 초기화된 32768바이트 배열을 가지고 있다.

		end만 적혀있으면 프로그램 종료
		%는 주석

		Brainfuck 언어가 입력되면 결과물을 출력한다.


		3
		++++++++[>+++++++++ % hello-world.
		<-]>.<+++++[>++++++<-]>-.+++++++..
		+++.<++++++++[>>++++<<-]>>.<<++++[>
		------<-]>.<++++[>++++++<-]>.+++.
		------.--------.>+.
		end
		+++[>+++++++[.
		end
		%% Print alphabet, A-Z.
		+ + + + + +++++++++++++++++++++>
		++++++++++++++++++++++++++++++++
		++++++++++++++++++++++++++++++++
		+< [ >.+<- ]
		end

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()
	T := input.([]int)[0] // 프로그램 명령어 갯수

	commands = make([]string, T)

	// 명령어 Input
	for i := 0; i < T; i++ {
		nextFlag := false
		for !nextFlag {
			input = readLine()
			command := input.(string)

			// BrainFuck 프로그램 끝
			if strings.EqualFold(command, "end") {
				nextFlag = true
				continue
			}
			// 프로그램에서 주석부분 제거
			if index := strings.Index(command, "%"); index != -1 {
				command = command[:index]
			}

			commands[i] += command
		}

	}

	// Output
	for idx, command := range commands {
		sb = strings.Builder{}
		sb.WriteString(fmt.Sprintf("PROGRAM #%d:\n", idx+1))
		// [ , ]가 문법에 맞는지 확인
		if !isVaildCommand(command) {
			sb.WriteString("COMPILE ERROR")
			writer.WriteString(sb.String() + "\n") // 여기 안해줬었네;
			continue
		}

		// BrainFuck 프로그램 기본 초기 설정
		defaultArr = [32768]byte{}
		defaultPointer = 0

		// 대괄호 매핑
		bracketMap := mappingBrackets(command)

		// brainfuck 해석
		interpreter(command, bracketMap)

		writer.WriteString(sb.String() + "\n")
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

func interpreter(command string, bracketMap map[int]int) {
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
			sb.WriteString(string(defaultArr[defaultPointer]))
		// 루프 작업
		case '[':
			if defaultArr[defaultPointer] == 0 {
				i = bracketMap[i] - 1 // i++ 되므로
				continue
			}
		case ']':
			if defaultArr[defaultPointer] != 0 {
				i = bracketMap[i]
				continue
			}
			// 루프 종료
		}
	}
}

// 명령어 괄호가 문제없는지 확인
func isVaildCommand(s string) bool {
	stack := []rune{}

	for _, char := range s {
		if char == '[' {
			stack = append(stack, char)
		} else if char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				// ']'가 나왔는데 짝이 맞지 않거나 스택이 비어있는 경우
				return false
			}
			// 짝이 맞아 스택에서 '['를 제거
			stack = stack[:len(stack)-1]
		}
	}

	// 스택이 비어있으면 모든 짝이 맞음
	return len(stack) == 0
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
