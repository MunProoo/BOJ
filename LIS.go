package main

// import (
// 	"bufio"
// 	"fmt"
// 	"strconv"

// 	"os"
// )

// var writer *bufio.Writer
// var N int
// var numList []int // 입력받은 수열
// var LIS []int     // 부분 수열의 길이

// func main() {

// 	/*
// 		11053 가장 긴 증가하는 부분 수열

// 		** 문제 해석 **
// 		: 수열 A가 주어졌을 때, 가장 긴 증가하는 부분 수열의 길이를 구하여라

// 		** 문제 풀이 **
// 		참고 : https://st-lab.tistory.com/285
// 		그냥 동적계획법은 시간복잡도가 O(N^2)임.
// 		이분탐색으로 시간복잡도를 O(NlogN)으로 줄일 수 있다

// 		방법 :
// 		(1) for문에서 numList의 값을 구하고 (target)
// 		(2) target이 LIS 수열에서 어디에 들어가면 좋을지 찾는다. >> LIS의 요소보다 target이 작다면 "치환"함

// 		** 이분 탐색 알고리즘을 이용할 경우 주의점
// 		LIS의 길이는 구해지지만,, LIS가 정답에 알맞는 부분수열인 것은 아니다.!!
// 		예를 들어 seq = [10, 20, 40, 35, 30, 45, 30, 70, 85] 면
// 		LIS = [10, 20, 40, 45, 70, 85] 겠지만
// 		실제 로직을 돌릴 경우 res = [10, 20, 30, 45, 70, 85] 가 나온다. (순서 무시하고 LIS를 만듦)

// 		seq = [10, 20, 30, 15, 20, 30, 50, 40, 45] 라면
// 		로직에서 LIS = [10, 15, 20, 30, 50] 이 되었을 때, 다음 값 40이 들어오는 경우 50을 40으로 치환하지 않으면
// 		45를 LIS에 넣지 못해 길이가 5로 마무리된다.
// 		50을 40으로 치환하면 LIS = [10, 15, 20, 30, 40, 45]로 길이가 6이 된다.

// 		이미 구한 길이에 대해서만 치환하므로 실제 길이와 일치한다.

// 	*/

// 	reader := bufio.NewReader(os.Stdin)
// 	scanner := bufio.NewScanner(reader)
// 	scanner.Split(bufio.ScanWords) // 스캐너 옵션 변경
// 	writer = bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()

// 	fmt.Fscan(reader, &N)

// 	numList = make([]int, N+1)
// 	for i := 1; i < N+1; i++ {
// 		scanner.Scan()
// 		numList[i], _ = strconv.Atoi(scanner.Text())
// 	}

// 	LIS = make([]int, 0)

// 	for i := 1; i < N+1; i++ {
// 		if len(LIS) == 0 {
// 			LIS = append(LIS, numList[i])
// 		}

// 		if LIS[len(LIS)-1] < numList[i] { // LIS의 가장 큰 수보다 크다면 바로 추가
// 			LIS = append(LIS, numList[i])
// 		} else { // 가장 큰 수보다 작다면, LIS의 요소와 비교하여 알맞은 곳에 배치한다. (작으면 교체)
// 			LISidx := binarySearchLIS(0, len(LIS)-1, numList[i])
// 			LIS[LISidx] = numList[i]
// 		}
// 	}

// 	fmt.Println(len(LIS))

// }

// // 탐색은 LIS 내에서 이루어짐 (start,end는 LIS 기준)
// func binarySearchLIS(start, end, target int) int {
// 	var mid int = 0

// 	for start < end {
// 		mid = (start + end) / 2
// 		if LIS[mid] < target {
// 			start = mid + 1
// 		} else {
// 			end = mid
// 		}
// 	}

// 	return end
// }
