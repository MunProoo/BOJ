package main

var (
	visitedDFS []int
	cnt        int
	nodes      [][]int
)

func dfs(start int) {
	visitedDFS[start] = cnt

	for _, val := range nodes[start] {
		if val != 1 && visitedDFS[val] == 0 {
			cnt++
			dfs(val)
		}
	}
}
