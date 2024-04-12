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

	// 풀이변수
)

func main() {

	/*
		3107. IPv6
		축약된 IPv6 주소가 주어졌을 때, 이를 원래로 복원하는 프로그램

		IPv6는 128비트. 16비트씩 :으로 끊어서 나타냄 (8그룹)
		-> 2001:0db8:85a3:0000:0000:8a2e:0370:7334

		1) 각 그룹의 앞자리의 0을 전체 혹은 일부를 생략가능
		-> 2001:db8:85a3:0:00:8a2e:370:7334

		2) 0으로만 이루어진 그룹이 있으면, 한 개이상 연속된 그룹을 콜론 2개로 바꿀 수 있음. 1번만 사용 가능

		-> 2001:db8:85a3::8a2e:370:7334

		      ::1   ->   0000:0000:0000:0000:0000:0000:0000:0001

	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	contractAddress := strings.Split(readLine(), ":")
	restoreAddress := ""

	length := len(contractAddress)

	// 2번 규칙이 없는 경우
	if length == 8 {
		// 0그룹복구만 하면 됨
		for _, group := range contractAddress {
			restoreAddress += FillZero(group)
		}

	} else {
		// ::이 있으므로 새로운 0그룹도 추가해야함
		var zeroGroupCount int
		zeroGroupList := make([]string, 8)
		zeroGroupList[1] = "0000:0000:" // 0그룹 2개
		zeroGroupList[2] = "0000:0000:0000:"
		zeroGroupList[3] = "0000:0000:0000:0000:"
		zeroGroupList[4] = "0000:0000:0000:0000:0000:"
		zeroGroupList[5] = "0000:0000:0000:0000:0000:0000:"
		zeroGroupList[6] = "0000:0000:0000:0000:0000:0000:0000:" // 0그룹 7개

		if length < 4 { // 이런 경우엔 2가지 분기처리 필요
			addressFlag := false
			addressLocation := 0
			for idx, group := range contractAddress {
				if group != "" {
					addressLocation = idx
					addressFlag = true
					break
				}
			}
			if addressFlag {
				// 1. 주소가 1개는 살아있는 경우 (맨 앞 or 맨 뒤가 살아있음)
				if addressLocation == 0 {
					restoreAddress += FillZero(contractAddress[addressLocation])
					restoreAddress += zeroGroupList[6]
				} else {
					restoreAddress += zeroGroupList[6]
					restoreAddress += FillZero(contractAddress[addressLocation])
				}

			} else {
				// 2. 주소가 모두 0인 경우
				restoreAddress = "0000:0000:0000:0000:0000:0000:0000:0000:"
			}

		} else {
			zeroGroupCount = 8 - length
			zeroGroupFlag := false
			for _, group := range contractAddress {
				if group == "" && !zeroGroupFlag {
					zeroGroupFlag = true
					restoreAddress += zeroGroupList[zeroGroupCount]
				} else {
					restoreAddress += FillZero(group)
				}
			}
		}
	}
	restoreAddress = restoreAddress[:len(restoreAddress)-1]

	fmt.Println(restoreAddress)

}

// 전체 혹은 일부를 잃어버린 0그룹 복구
func FillZero(group string) (newAddress string) {

	groupLength := len(group)
	if groupLength != 4 {
		switch groupLength {
		case 1:
			newAddress = "000" + group
		case 2:
			newAddress = "00" + group
		case 3:
			newAddress = "0" + group
		}

		newAddress += ":"
	} else {
		newAddress += group + ":"
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
