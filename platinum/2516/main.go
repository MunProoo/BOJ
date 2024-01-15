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
		1번 원숭이는 무조건 cage1에 넣는다
		1번 원숭이의 앙숙중 첫번째는 cage1에 넣는다.
			앙숙 2번째, 3번째는 cage2에 넣는다.
		cage에 넣어진 원숭이는 visited를 true로 한다.

		* cage1에서 앙숙관계에 대해 확인이 필요하다.
		** cage2에 적어도 1마리는 넣어야 한다.


		로직 문제점
		enemy가 대상보다 크면 아직 안들어갔을 확률이 있음. ->  뒤늦게 들어올 수 있음
		enemy의 앙숙이 나밖에 없으면 입구컷 안당하고 cage1 들어온다.

		ex )4 <-> 5
		    5 <-> 4,6
			6 <-> 5
			일 때, 6을 아직 확정짓지 않은 상태에서 5를 체크해서 cage1에 넣은경우,
			추후 6의 가능성에서 'cage1 넣어도 된다고' 계산됨 (cage1에 앙숙은 5 1마리밖에 없으니까).

			before : 6이 들어온다 -> 앙숙을 체크했을 때 넣어도 되는가? 에서
			after : 6이 들어온다 -> 대상이 들어갔을 때 위법되는가? -> 본인이 들어갔을 때 앙숙이 위법되는가?



	*/

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

	for i := 1; i <= N; i++ {
		putMonkeyInCage(i)
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

func putMonkeyInCage(i int) {
	// i : 대상     enemy : 앙숙

	if !visited[i] {
		isPossibleCage1 := isPossiblePushCage1(i) // cage1에 넣어도 되는지 앙숙관계 확인
		if isPossibleCage1 {
			cage.cage1Map[i] = true
			cage.cage1 = append(cage.cage1, i) // 실제 케이지에 넣는다
			visited[i] = true
		} else {
			cage.cage2Map[i] = true
			cage.cage2 = append(cage.cage2, i) // 실제 케이지에 넣는다
			visited[i] = true
		}
	}

	for enemy := range enemyList[i] {
		if visited[enemy] {
			continue
		}
		putMonkeyInCage(enemy)
	}
}

func isPossiblePushCage1(monkey int) bool {
	// 앙숙을 기준 가능성 체크 (대상을 cage1에 넣어도 되는가)
	cnt := findCage1EnemyCnt(monkey)

	if cnt >= 2 {
		return false
	}

	// 대상 기준으론 넣어도 됨. 앙숙들은 대상이 들어와도 괜찮은지 확인
	cage.cage1Map[monkey] = true // 임시로 넣어놓는다
	for enemy := range enemyList[monkey] {
		enemyCount := findCage1EnemyCnt(enemy)
		if enemyCount >= 2 && cage.cage1Map[enemy] != nil { // 기존에 넣은 원숭이에게 문제생김
			cage.cage1Map[monkey] = nil // 넣으면 안되니까 다시 뺀다
			return false
		}
	}

	return true
}

// Cage1에 앙숙이 몇 마리 들어있는지 체크
func findCage1EnemyCnt(monkey int) int {
	cnt := 0
	for enemy := range enemyList[monkey] {
		if _, isExist := cage.cage1Map[enemy]; isExist {
			cnt++
		}
	}
	return cnt
}
