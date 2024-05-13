package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var (
	writer *bufio.Writer
	reader *bufio.Reader
	// sb     strings.Builder
)

func main() {

	/*
		11723. 집합

		비트 마스킹 이용하여 풀이
		+ 모든 입출력 byte로
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N := readLineInt()[0]
	var N int
	fmt.Fscanln(reader, &N)
	var out []byte = make([]byte, 0, 7376000)

	var S uint64

	for i := 0; i < N; i++ {

		// 기존에 사용하던 string 입력
		// input := strings.Split(readLine(), " ")
		// var command, number string
		// var x int
		// if len(input) > 1 {
		// 	command, number = input[0], input[1]
		// 	x, _ = strconv.Atoi(number)
		// } else {
		// 	command = input[0]
		// }

		// byte로 입력 및 처리
		input, _, _ := reader.ReadLine()
		var command []byte
		var x uint64
		sIndex := bytes.IndexByte(input, ' ')
		if sIndex > 0 {
			command = input[:sIndex]
			xbytes := input[sIndex+1:]
			if len(xbytes) == 2 {
				x = 10*uint64(xbytes[0]-'0') + uint64(xbytes[1]-'0')
			} else {
				x = uint64(xbytes[0] - '0')
			}
		} else {
			command = input
		}

		switch string(command) {
		case "add":
			S |= 1 << x
		case "remove":
			// X번째 비트가 1이고 나머지가 0인 값을 만들어서 반전시키면 X번째 비트만 뺄 수 있다.
			S &= ^(1 << x)
		case "check":
			if S&(1<<x) != 0 {
				out = append(out, '1', '\n')
				// sb.WriteString("1\n")
			} else {
				out = append(out, '0', '\n')
				// sb.WriteString("0\n")
			}

		case "toggle":
			S = S ^ (1 << x)
		case "all":
			S |= (1 << 21) - 1

		case "empty":
			S = 0
		}
	}

	writer.Write(out)
}
