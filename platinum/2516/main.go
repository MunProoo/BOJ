package main

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"os"
)

var (
	writer    *bufio.Writer
	reader    *bufio.Reader
	enemyList []map[int]bool
	visited   []bool
	cage      Cage
)

const (
	Cage1 = iota
	Cage2
)

type Cage struct {
	cage1Map map[int]interface{}
	cage1    []int
	cage2Map map[int]interface{}
	cage2    []int
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
		반씩 1,2에 넣는다
		1에서 안되는 애는 2에 넣는다
		2에서 안되는 애는 1에 넣는다

		둘 다 문제 없이 되도록 계속 반복한다
		-> 무한 루프 돌게 됨.
		그럼 어카지

		일단 다 밀어넣고, 앞번호 에서부터 앙숙이 2마리인 애들은 걸러내면 되려나
		체크해야 하는 거
		1. cage1에 본인 앙숙이 2마리 이상 있다 -> cage1에서 빼야한다.
		2. cage2에 내가 들어가면, 안되는 애들이 있다. (본인 앙숙들의 앙숙카운트 필요) -> 패스해야한다.


		*** 최종 풀이
		그리디한 생각!!
		1. 한곳에 몰아 넣는다.
		2. cage1에서 안되는 애들을 cage2로 넣는다.
		3. cage2에서 안되는 애들을 cage1로 넣는다.

		작업 없을때까지 반복한다. (난 동시에 하려해서 안됨)

	*/

	// readFile, _ := os.Open("D:\\GoWorkspace\\src\\algorithm\\platinum\\2516\\input.txt")
	// readFile, _ := os.Open("D:\\GoWorkspace\\src\\algorithm\\platinum\\2516\\예제2.txt")
	// readFile, _ := os.ReadFile("input.txt")
	// reader = bufio.NewReader(readFile)

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input interface{}
	input = readLineInt()

	N := input.([]int)[0]
	enemyList = make([]map[int]bool, N+1)
	visited = make([]bool, N+1)

	for i := 1; i <= N; i++ {
		input = readLineInt()
		// M := input.([]int)[0]
		enemys := input.([]int)[1:]

		enemyList[i] = make(map[int]bool)
		for _, enemy := range enemys {
			enemyList[i][enemy] = true
		}
	}

	cage.cage1Map = make(map[int]interface{})
	cage.cage1 = make([]int, 0)
	cage.cage2Map = make(map[int]interface{})
	cage.cage2 = make([]int, 0)

	// 원숭이들 한우리에 밀어넣음
	putMonkeyInCage(N)

	cageType := Cage1 // 다음 작업은 어느 cage에서 해야하는가 값
	var repeatedFlag bool

	for {
		if cageType == Cage1 {
			cageType, repeatedFlag = cage1에서내쫓기()
		} else {
			cageType, repeatedFlag = cage2에서내쫓기()
		}

		if !repeatedFlag {
			break
		}
	}

	for monkey := range cage.cage1Map {
		cage.cage1 = append(cage.cage1, monkey)
	}

	for monkey := range cage.cage2Map {
		cage.cage2 = append(cage.cage2, monkey)
	}

	if len(cage.cage2) == 0 {
		// cage2에 아무것도 없으면 cage1에서 빼서 넣어준다
		cage.cage2 = append(cage.cage2, cage.cage1[0])
		cage.cage1 = cage.cage1[1:]
	}

	sort.Ints(cage.cage1)
	sort.Ints(cage.cage2)
	writer.WriteString(strconv.Itoa(len(cage.cage1)) + " ")
	for _, monkey := range cage.cage1 {
		writer.WriteString(strconv.Itoa(monkey) + " ")
	}

	writer.WriteString("\n" + strconv.Itoa(len(cage.cage2)) + " ")
	for _, monkey := range cage.cage2 {
		writer.WriteString(strconv.Itoa(monkey) + " ")
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

// 일단 전부 cage1로
func putMonkeyInCage(N int) {
	for i := 1; i <= N; i++ {
		cage.cage1Map[i] = true
	}
}

// Cage에 있어도 되는지 확인
func isPossibleInCage(monkey int, flag int) bool {
	var anyCage map[int]interface{}
	if flag == 1 { // cage1 체크
		anyCage = cage.cage1Map
	} else { // cage2 체크
		anyCage = cage.cage2Map
	}

	// 본인이 들어가도 되는지 여부 확인
	cnt := 0
	for enemy := range enemyList[monkey] {
		if _, isExist := anyCage[enemy]; isExist {
			cnt++
		}
	}

	return cnt < 2
}

func cage1에서내쫓기() (cageType int, repeatedFlag bool) {
	for monkey := range cage.cage1Map {
		// 원숭이가 cage1에서 못있는 애인지 확인
		if possible := isPossibleInCage(monkey, 1); !possible {
			// cage1에 못있으면 cage2로 이동시키고 확인
			cage.cage2Map[monkey] = true
			delete(cage.cage1Map, monkey)
			repeatedFlag = true
		}
	}

	return Cage2, repeatedFlag
}

func cage2에서내쫓기() (cageType int, repeatedFlag bool) {
	for monkey := range cage.cage2Map {
		// 원숭이가 cage2에서 못있는 애인지 확인
		if possible := isPossibleInCage(monkey, 2); !possible {
			// cage2에 못있으면 cage1로 이동시키고 확인
			cage.cage1Map[monkey] = true
			delete(cage.cage2Map, monkey)
			repeatedFlag = true
		}
	}
	return Cage1, repeatedFlag
}
