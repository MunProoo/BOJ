package main

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"os"
)

var (
	writer *bufio.Writer
	reader *bufio.Reader
	sb     strings.Builder

	// 풀이변수
)

func main() {

	/*
		5544. 리그

		축구 경기에서 승리 시 +3점
		무승부 시 +1점
		패배 시 0점

		첫째 줄에 팀의 수 N (2 ≤ N ≤ 100)가 주어진다. 다음 N(N-1)/2개 줄에는 각 경기의 결과가 주어진다.
		경기의 결과는 A B C D와 같이 네 개의 정수로 나타내며, A팀이 C점, B팀이 D점을 획득했음을 의미한다.
		A와 B는 항상 다르다. 한 경기의 결과가 여러 번 주어지는 경우는 없다.

		승점이 동일한 경우 공동2위, 3위...
		각 팀의 순위를 구하라.

		입력 받자마자 승점을 계산한다.
		팀의 승점을 map에 담아 놓는다?
		순위를 어떻게 매길것인가 고민해야댐
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	N := readLineInt()[0]
	points := make(map[int]int) // 각 팀의 승점 누적

	gameCount := N * (N - 1) / 2
	for i := 0; i < gameCount; i++ {
		input := readLineInt()
		if input[2] > input[3] {
			points[input[0]] += 3
		} else if input[2] < input[3] {
			points[input[1]] += 3
		} else {
			points[input[0]] += 1
			points[input[1]] += 1
		}
	}

	league := make([]Team, 0)
	for key, val := range points {
		team := Team{
			Team:  key,
			Point: val,
		}
		league = append(league, team)
	}

	// 내림차순 정렬
	sort.Slice(league, func(i, j int) bool {
		return league[i].Point > league[j].Point
	})

	Ranks := make(map[int]int)

	for i, val := range league {
		if i != 0 {
			// 승점이 중복되는 경우
			if league[i-1].Point == val.Point {
				continue
			}
		}

		Ranks[val.Point] = i + 1
	}

	sort.Slice(league, func(i, j int) bool {
		return league[i].Team < league[j].Team
	})

	for _, val := range league {
		sb.WriteString(strconv.Itoa(Ranks[val.Point]) + "\n")
	}

	writer.WriteString(sb.String())

}

type Team struct {
	Team  int
	Point int
	Rank  int
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
