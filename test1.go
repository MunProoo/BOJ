package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// /*
// 인스턴스 문제
// 10799 쇠막대기

// 레이저는 항상 ()꼴로 옴.
// ')'가 나왔을 때, 이전 문자가 '(' 가 아니면 쇠막대기의 끝을 의미한다.

// 스택의 길이 : 쇠막대기 갯수

// 레이저가 나오면 -> barCount += len(stack)
// 쇠막대기의 끝이면 -> 해당 쇠막대기의 짜투리를 barCount에 추가 barCount++

// */

// func main() {
// 	// reader := bufio.NewReader(os.Stdin)
// 	writer := bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()

// 	// var input string
// 	// fmt.Fscanln(reader, &input)
// 	a := "INTP"
// 	b := "ISFJ"
// 	c := "ENTP"

// 	fmt.Println(getDistance(a, b, c))
// }
// func getDistance(a, b, c string) int {
// 	distance := 0
// 	for i := 0; i < 4; i++ {
// 		if a[i] != b[i] {
// 			distance++
// 		} else if b[i] != c[i] {
// 			distance++
// 		} else if a[i] != c[i] {
// 			distance++
// 		}
// 	}
// 	return distance
// }

// 큐
/*
type Que []int // 정수형 슬라이스로 Que 타입 추가

func (s *Que) push(n int) {
	*s = append(*s, n) // 끝에 값 추가
}

func (s *Que) pop() int {
	if len(*s) == 0 {
		return 0
		// writer.WriteString("-1\n")
	} else {
		// 처음값 := fmt.Sprintf("%v", (*s)[0])
		처음값 := (*s)[0]
		// writer.WriteString(처음값 + "\n")

		*s = (*s)[1:] // 처음값 제거하고 남은 애들 할당
		return 처음값
	}
}
*/
