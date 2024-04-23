package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"os"
)

var (
	writer *bufio.Writer
	reader *bufio.Reader
	sb     strings.Builder

	// 풀이변수
	N, M int
	// result [][]int // map보단 슬라이스가 더 빠르겠지라 생각했지만 고려해야하는 인덱스가 2천만개면 얘기가 달라지지
)

func main() {

	/*
		21278. 호석이 두마리치킨
		치킨집을 2개 연다.
		모든 건물과 접근성이 가장 좋게 짓는 방법은?
		여러 가지라면, 건물 번호가 작게.

		첫번째로 , 가장 자식이 많은 노드를 선택, 해당 자식들에 방문처리를 하고
		두번째로 방문하지 않은 자식이 가장 많은 노드를 선택.
				방문하지 않은 자식이 동일하다면 숫자가 작게. 이렇게 해볼까

		출력 : 건물 2개가 지어질 건물 번호 오름차순 , 모든 도시에서의 왕복 시간의 합
		이렇게 풀기엔, 가장 작은 건물번호를 골라내기 위해서 노가다가 많아짐..

		<다른 풀이>
		1. A->E 등 노드간의 최소 거리 구하기 (플로이드-워셜 정리)
		2. 건물 2개 선택해서 모든 노드간의 거리 계산 (브루트포스)
		3. 풀이과정 2가 최소가 되는 2개의 점과 합을 구한다.


		[플로이드-워셜]
		그래프에서 가능한 모든 노드에 대해 최단거리를 구하는 알고리즘
		임의의 노드 s에서 임의의 노드 t까지 거리를 구한다면,
		s와 t사이의 노드 m에 대해 s-m 까지의 최단거리, m-t까지의 최단거리를 구해서 테이블에 기입
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input := readLineInt()
	N, M := input[0], input[1]

	infinite := math.MaxInt     // 무한을 의미 (연결 안됨)
	graph := make([][]int, N+1) // 건물 그래프

	for i := 0; i <= N; i++ {
		graph[i] = make([]int, N+1)
		for j := range graph[i] {
			graph[i][j] = infinite // 기본값 무한으로 초기화
		}
	}

	for i := 0; i < M; i++ {
		input := readLineInt()
		building1, building2 := input[0], input[1]
		graph[building1][building2] = 2 // 건물 연결 (왕복시간으로)
		graph[building2][building1] = 2
	}

	fwTable := FloydWarshall(graph)

	now := infinite
	building1, building2 := 0, 0

	for i := 1; i <= N; i++ { // 건물1
		for j := 1; j <= N; j++ { // 건물2
			if i == j {
				continue
			}
			// 건물 2개를 골랐을 때 최소 거리
			minDistance := CalcDistance(fwTable, i, j)
			if minDistance < now {
				now = minDistance
				building1 = i
				building2 = j
			}
		}
	}

	fmt.Println(building1, building2, now)

}

// 플로이드-워셜 정리
// return : 플로이드-워셜 테이블
func FloydWarshall(graph [][]int) [][]int {
	n := len(graph)
	fwTable := make([][]int, n) // 플로이드-워셜 테이블
	infinite := math.MaxInt

	for i := range fwTable {
		fwTable[i] = make([]int, n)
		copy(fwTable[i], graph[i])
	}

	// 플로이드-워셜 알고리즘
	for k := 0; k < n; k++ { // 중간 노드
		for i := 0; i < n; i++ { // 출발 노드
			for j := 0; j < n; j++ { // 도착지점 노드
				if fwTable[i][k] != infinite && fwTable[k][j] != infinite && fwTable[i][j] > fwTable[i][k]+fwTable[k][j] {
					// i -> k -> j가 가능하면서 더 짧은거리라면 갱신
					fwTable[i][j] = fwTable[i][k] + fwTable[k][j]
				}
			}
		}
	}

	return fwTable
}

// 선택된 건물 2개에 대해 모든 노드와의 최단 거리합 계산
// return : 최단 거리의 합
func CalcDistance(table [][]int, building1, building2 int) (distance int) {
	length := len(table)

	for i := 1; i < length; i++ {
		if i == building1 || i == building2 {
			// 같은 건물은 왕복 거리 0
			continue
		}
		// 두 건물중 가까운 거로 채택
		distance += min(table[building1][i], table[building2][i])
	}

	return
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
