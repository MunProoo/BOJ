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
	sb     strings.Builder
)

func main() {

	/*
		9082 지롸칮기
		https://www.acmicpc.net/problem/9082

		숫자는 8방향중에 지뢰가 존재함을 알려줌
		가능한 지뢰의 수 중 최댓값을 출력하라

		혼자 힘으로 풀지 못해서 참고
		https://baby-dev.tistory.com/entry/Python-%EB%B0%B1%EC%A4%80-9082%EB%B2%88-%EC%A7%80%EB%A2%B0-%EC%B0%BE%EA%B8%B0

		<로직>
		number : 숫자 배열
		result : 출력할 답안 (지뢰 개수)

		number의 각 인덱스 숫자를 볼 때마다 지뢰가 있을지 없을지를 확인하는 방법임.
		각 인덱스의 숫자는 idx-1, idx, idx+1 중에 지뢰가 있다는 의미임. 앞에서부터 지뢰가 있다치고 (result++), 지뢰가 있을 위치들에 -1개 한다. 범위내에 0이 하나라도 있으면 지뢰는 없다.

		ex) 1 1 1 2 2 라면
		number[0] = 1  >> 0번, 1번 인덱스중에 지뢰가 1개 있으므로 result++ , number[0]-- , number[1]--
		number[1] = 0  >> 0이 됐으니까 패스
		number[2] = 1  >> 1번, 2번, 3번 인덱스 중에 지뢰가 1개 있음. 그러나 number[1]이 0으로 바뀌었으니 지뢰는 0개
		number[3] = 2  >> 2번, 3번, 4번 인덱스 중에 지뢰가 2개 있음. result++ , number[2]--, number[3]--, number[4]--
		number[4] = 1  >> 3번, 4번 중에 지뢰가 1개 있음. result++, number[3]--, number[4]--

		결과 : 0에 1개, 3에 1개, 4에 1개
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input := readLineInt()
	T := input[0]
	for T > 0 {
		input = readLineInt()
		N := input[0]

		number := readLineIntSplit2()     // 지뢰 수
		_ = strings.Split(readLine(), "") // 지뢰 string
		result := 0

		// <로직>
		for idx, val := range number {
			if val == 0 {
				continue
			} else {
				if idx == 0 {
					if number[idx] != 0 && number[idx+1] != 0 { // 범위 내에 지뢰가 1개 이상 있는 경우만
						result++

						// 범위내 지뢰 개수 줄이기
						number[idx]--
						number[idx+1]--
					}
				} else if idx == N-1 {
					if number[idx] != 0 && number[idx-1] != 0 { // 범위 내에 지뢰가 1개 이상 있는 경우만
						result++

						// 범위내 지뢰 개수 줄이기
						number[idx]--
						number[idx-1]--
					}
				} else {
					if number[idx-1] != 0 && number[idx] != 0 && number[idx+1] != 0 { // 범위 내에 지뢰가 1개 이상 있는 경우만
						result++

						// 범위내 지뢰 개수 줄이기
						number[idx-1]--
						number[idx]--
						number[idx+1]--
					}
				}
			}
		}
		sb.WriteString(fmt.Sprintf("%d\n", result))
		T--
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

func readLineIntSplit2() []int {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	inputSlice := strings.Split(input, "")

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

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

// 기록용
/*
2xN 의 배열이므로 하나의 지뢰의 영향력은 \|/ 3칸임

		number를 보고 지뢰가 있을 수 있는 자리 중 확률이 높은 쪽을 선택한다

		1 1 1 2 2
		0번 자리에 1이 있으니 0,1중에 지뢰가 1개 있다
		1번 자리에 1이 있으니 0,1,2중에 지뢰가 1개 있다
		2번 자리에 1이 있으니 1,2,3중에 지뢰 1개
		3번 자리에 2가 있으니 2,3,4중에 지뢰 1개
		4번 자리에 2가 있으니 3,4중에 지뢰 2개

		결과 취합:
		0번 인덱스 : 0,1에 1개.
		1번 인덱스 : 0,1,2에 1개. 0,1에 1개가 있어야 하니까 2는 지뢰 없음
		2번 인덱스 : 1,2,3에 1개. 2는 지뢰 없고 1,3중에 1개 있음
		3번 인덱스 : 2,3,4에 2개. 2는 지뢰 없으므로 3,4에 1개씩 있음
		4번 인덱스 : 3,4에 2개.

		최대 지뢰수 : 0에 1개 3,4에 1개 총 3개

		로직으로 만들면?
		inference []int // 추론용 배열
		0 : 지뢰 있을 수도
		1 : 지뢰 확실히 없음
		2 : 지뢰 있음

		<로직>
		1. 0번 인덱스 -> inference 배열에 반영한다. inference[0], inference[1] = 0
		2. 1번 인덱스 -> 0,1에 1개가 있다는 정보를 토대로    inference[2] = 1
		3. 2번 인덱스 -> inference[1], inference[3] = 0
		4. 3번 인덱스 -> inference[3], inference[4] = 2 ,  2자리가 남았는데 지뢰 2개가 있으니까.
		5. 4번 인덱스 -> 이미 4번의 과정에서 답이 나왔음.

		6. 2번 인덱스에서 [3]아니면 [1]에 지뢰가 1개 있는데, 3번 인덱스에서의 결과로 [3],[4]는 지뢰가 있다. 따라서 [1]은 지뢰가 0개

*/
