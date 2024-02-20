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
	inputs         string // brainfuck에 입력될 문자

	arrSize int // defaultArr 크기
)

const (
	commandLimit = 50000000
)

type Pointer int

func (p *Pointer) Increase() {
	if int(*p) == arrSize-1 {
		*p = 0
		return
	}
	*p++
}
func (p *Pointer) Decrease() {
	if *p == 0 {
		*p = Pointer(arrSize) - 1
		return
	}
	*p--
}

type CheckValue struct {
	value   byte
	pointer Pointer
}

func main() {

	/*
		3954 Brainfuck 인터프리터
		역대급 재미없었던 문제.. 무한루프의 정의부터 명확하지 않아서 많이 헤멨다.
		질문게시판 엄청 이용함..

		, : 문자 하나를 읽고 포인터가 가리키는 곳에 저장한다. (입력의 마지막이면 255를 저장한다)

		Brainfuck 프로그램이 무한루프인지 아닌지만 확인한다.
		(5천만번을 반복해야 끝나는 루프는 주어지지 않는다)

		풀이 1.
		(1)무한루프인 조건을 생객해내 걸러낸다.
		- 아무 명령어가 없다.
		- 아무리 반복해도 루프가 끝나지 않는 조건이다
		>> 처음엔 diff = (반복 이전 값 - 이후 값)을 비교하여 절대로 defaultArr[defaultPointer]가 0이 안되는 조건을 설정했었는데,
		ex ) 255 % diff != 0 이면 절대 닿을 수 없다.

		이 부분은 값이 오버플로우, 언더플로우되면 diff값이 달라지므로 반례가 생긴다.
		오버플로우 언더플로우의 경우를 배제하기 위해서 어떻게 할까 고민하다가, Threshold를 사용하여 해결하였다.


		- 명령어 실행이 5천만번이상이다.

		(2)만약 5천만번을 돌았다면, 무한루프 속에 있다고 판단하여 가장 바깥의 루프를 출력한다.


	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	// 테스트케이스
	input = readLineInt()
	T := input.([]int)[0]

	// 명령어 Input
	for i := 0; i < T; i++ {
		// 입력1
		input = readLineInt()
		ints := input.([]int)
		arrSize, _, _ = ints[0], ints[1], ints[2]

		defaultArr = make([]byte, arrSize)
		defaultPointer = 0

		// 입력2 (brainfuck 프로그램)
		input = readLine()
		command := input.(string)

		// 입력3 (brainfuck 프로그램의 입력)
		input = readLine()
		inputs = input.(string)

		// 대괄호 매핑 , 감싼 루프 체크
		bracketMap, enclosingBrackets := mappingBrackets(command)
		interpreter(command, bracketMap, enclosingBrackets)

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

func interpreter(command string, bracketMap map[int]int, enclosingBrackets map[int]int) {
	inputCnt := 0 // brainfuck 입력 사용 개수
	inputLength := len(inputs)
	// loopFlag := false
	// var initValue byte
	initValueMap := make(map[int]CheckValue) // 반복문이 시작할 때 defaultPointer 값. (반복문을 통해서 얼마의 차이가 생기는지 확인용)

	// 루프를 돌았는지 안돌았는지 확인
	loopFlag := make(map[int]bool) // 중첩 반복문인 경우 기존 변수 하나로는 덮어쓰기가 되므로 map으로 설정
	commandCount := 0              // 커맨드 카운트

	for i := 0; i < len(command); i++ {
		switch command[i] {
		case '>':
			commandCount++
			defaultPointer.Increase()
		case '<':
			commandCount++
			defaultPointer.Decrease()
		case '+':
			commandCount++
			if defaultArr[defaultPointer] == 255 {
				defaultArr[defaultPointer] = 0

			} else {
				defaultArr[defaultPointer]++
			}
		case '-':
			commandCount++
			if defaultArr[defaultPointer] == 0 {
				defaultArr[defaultPointer] = 255
			} else {
				defaultArr[defaultPointer]--
			}
		case '.':
			commandCount++
		// 루프 작업
		case '[':
			commandCount++
			if defaultArr[defaultPointer] == 0 { // 루프가 끝남
				i = bracketMap[i] - 1 // i++ 되므로
				// loopFlag = false
				continue
			}

		case ']':
			commandCount++
			start := bracketMap[i]
			end := i

			// 반복하기 전의 값 할당
			prevValue := initValueMap[start]
			initValue := prevValue.value
			initPointer := prevValue.pointer

			initValueMap[start] = CheckValue{defaultArr[defaultPointer], defaultPointer}
			if defaultArr[defaultPointer] != 0 {
				// 무한루프 판별 안되면 다시 반복해야함. 기존에는 이 코드의 위치가 이상해서 잘못나오고 있었음
				i = bracketMap[i]

				// ,로 인한 EOF라는 변수가 있어서 무조건 루프 돌려야함
				if strings.Contains(command[start:end], ",") {
					continue
				}

				// 이미 한 번은 돌고 왔음
				if loopFlag[i] {
					// 명령어가 아예 없어서 무한루프인 경우
					if (end - start - 1) == 0 {
						sb.WriteString(fmt.Sprintf("Loops %d %d\n", start, end))
						return
					}

					currentValue := int(defaultArr[defaultPointer])

					diff := abs(int(initValue) - currentValue)
					threshold := 130 // 255의 대략 반보다 크면 오버플로우나 언더플로우가 일어난 경우로 생각하여 건너뛰기
					if diff >= threshold {
						continue
					}

					// if diff == 0 && initPointer == defaultPointer { // 값의 변화가 없는 경우. -> 무한 루프
					// (1번째 반복 종료시 포인터 == 2번째 반복 종료 시 포인터) && (255 % diff) == 0이 아니라면, 무한루프다.
					if diff != 0 {
						expression := 255 % int(diff)
						// 오버,언더 플로우 후에 0을 지나침
						if expression != 0 && initPointer == defaultPointer {
							sb.WriteString(fmt.Sprintf("Loops %d %d\n", start, end))
							return
						}
					}

					// 생각해낸 무한루프를 걸러냈고, 그럼에도 5천만번 이상 명령이 실행됐으니 실행 루프를 감싸고 있는 루프를 무한루프로 생각
					if commandCount > commandLimit {
						if encStart, exist := enclosingBrackets[start]; exist {
							encEnd := bracketMap[encStart]
							sb.WriteString(fmt.Sprintf("Loops %d %d\n", encStart, encEnd))
							return
						} else {
							sb.WriteString(fmt.Sprintf("Loops %d %d\n", start, end))
							return
						}
					}

				}
				loopFlag[start] = true // 반복시키기 위해서 i의 값을 초기화했는데, 이 코드에선 반영을 안해줘서 루프 확인이 잘 안됐음
			} else {
				loopFlag[start] = false
			}

			// 루프 종료
		case ',':
			commandCount++
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
func mappingBrackets(s string) (bracketMap map[int]int, enclosingBrackets map[int]int) {
	stack := []int{}
	bracketMap = make(map[int]int)
	enclosingBrackets = make(map[int]int) // 가장 바깥쪽 루프
	encStart := 0

	for idx, char := range s {
		if char == '[' {
			if len(stack) == 0 {
				encStart = idx
			} else {
				enclosingBrackets[idx] = encStart
			}
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

	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
