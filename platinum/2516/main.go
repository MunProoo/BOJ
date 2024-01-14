package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"os"
)

var (
	writer    *bufio.Writer
	reader    *bufio.Reader
	enemyList [][]int
	visited   []bool
	cage      Cage
)

type Cage struct {
	cage1 []int
	cage2 []int
}

func main() {

	/*
		2516 원숭이
		1. 원숭이 한마리에게 앙숙관계인 원숭이는 최대 3마리
		2. 원숭이 한마리와 앙숙관계인 원숭이는 같은 우리에 최대 1마리
		3. 빈 우리는 없다.
		4. 우리는 2개
		5. 시간제한 1초 (10^8) . 입력제한이 10만이므로 O(1) ~ O(logN) ~ O(N) ~ O(NlogN) 까지만 가능할 것

		입력 :
		N (3이상 10만 이하)
		번호대로 앙숙의 수(M), 앙숙관계 번호

		풀이 :
		1번 원숭이는 무조건 cage1에 넣는다
		1번 원숭이의 앙숙중 첫번째는 cage1에 넣는다.
			앙숙 2번째, 3번째는 cage2에 넣는다.
		cage에 넣어진 원숭이는 visited를 true로 한다.

		일단 해보자.




	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()

	N := input.([]int)[0]
	enemyList = make([][]int, N+1)
	visited = make([]bool, N+1)

	for i := 1; i <= N; i++ {
		input = readLineInt()
		// M := input.([]int)[0]
		enemys := input.([]int)[1:]

		enemyList[i] = make([]int, 0)
		enemyList[i] = append(enemyList[i], enemys...)
	}

	cage.cage1 = make([]int, 0)
	cage.cage2 = make([]int, 0, 3)

	for i := 1; i <= N; i++ {
		putMonkeyInCage(i)
	}

	fmt.Printf("%v", cage)

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

func putMonkeyInCage(i int) {
	// i : 대상     enemy : 앙숙
	if !visited[i] {
		cage.cage1 = append(cage.cage1, i)
		visited[i] = true
	}

	for idx, enemy := range enemyList[i] {
		if visited[enemy] {
			continue
		}
		putMonkeyInCage(enemy)
	}
}
