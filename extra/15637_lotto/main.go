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
		15637. 로또

		로또 1~700회까지 숫자중 N이 몇번 나왔는지 출력
		excel 파일로 다운받은 로또 숫자를 파싱하여 풀이했음

		1:103 2:88 3:91 4:99 5:101

		6:86 7:96 8:100 9:71 10:91

		11:96 12:91 13:98 14:102 15:96

		16:85 17:102 18:97 19:91 20:112

		21:85 22:79 23:86 24:91 25:97

		26:98 27:110 28:80 29:88 30:80

		31:95 32:83 33:93 34:109 35:88

		36:91 37:103 38:90 39:95 40:113

		41:81 42:86 43:97 44:93 45:93

		로또 API를 이용해서도 가능
	*/

	reader = bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// // 엑셀 파일 파싱
	// data := make(map[int]int)
	// file, err := excelize.OpenFile("data.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// cells := []string{"N", "O", "P", "Q", "R", "S"}
	// for i := 4; i < 704; i++ {
	// 	// 4~703
	// 	for _, cell := range cells {
	// 		cell += strconv.Itoa(i)
	// 		cellValue, err := file.GetCellValue("Sheet1", cell)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		number, _ := strconv.Atoi(cellValue)
	// 		data[number]++
	// 	}
	// }
	// fmt.Println(data)

	data := make(map[int]int)
	data[1] = 103
	data[2] = 88
	data[3] = 91
	data[4] = 99
	data[5] = 101
	data[6] = 86
	data[7] = 96
	data[8] = 100
	data[9] = 71
	data[10] = 91
	data[11] = 96
	data[12] = 91
	data[13] = 98
	data[14] = 102
	data[15] = 96
	data[16] = 85
	data[17] = 102
	data[18] = 97
	data[19] = 91
	data[20] = 112
	data[21] = 85
	data[22] = 79
	data[23] = 86
	data[24] = 91
	data[25] = 97
	data[26] = 98
	data[27] = 110
	data[28] = 80
	data[29] = 88
	data[30] = 80
	data[31] = 95
	data[32] = 83
	data[33] = 93
	data[34] = 109
	data[35] = 88
	data[36] = 91
	data[37] = 103
	data[38] = 90
	data[39] = 95
	data[40] = 113
	data[41] = 81
	data[42] = 86
	data[43] = 97
	data[44] = 93
	data[45] = 93

	N := readLineInt()[0]
	fmt.Println(data[N])
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
