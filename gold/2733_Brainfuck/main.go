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
	// defaultArr []byte
	sb             strings.Builder
	defaultArr     [32768]byte
	defaultPointer Pointer
	// roopMemory     []int  // 이부분은 스택으로 해야겠다. 추후 문법에러는 없는지 확인 후 작업들어가도록
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
	sb = strings.Builder{}
	defer writer.Flush()

	var input interface{}
	input = readLineInt()
	T := input.([]int)[0] // 프로그램 명령어 갯수

	defaultArr = [32768]byte{}
	defaultPointer = 0

	for T > 0 {
		input = readLine()
		command := input.(string)

		// brainfuck 프로그램 종료
		if strings.EqualFold(command, "end") {
			T--
		}

		// 명령어 한 줄을 단위 명령어로 바꿔서 확인
		interpreter(command)
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

func interpreter(command string) {
	for _, unitCommand := range command {
		switch unitCommand {
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
		case '[':
			// 루프 작업 -> 스택으로 해야할듯?
		case ']':
			// 루프 작업
		case '%':

		}
	}
}
