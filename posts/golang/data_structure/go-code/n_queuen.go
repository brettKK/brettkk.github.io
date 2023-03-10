package main

import "fmt"

var res [][][]byte
func sovle_N_queen(N int) {
	res = make([][][]byte, 0)
	var queen = [][]byte{}
	for i := 0; i < N; i++ {
		queen[i] = make([]byte, N)
		for j := 0; j < N; j++ {
			queen[i][j] = '.'
		}
	}
	var attack = [][]int{}
	dfs(0, N, queen, attack)
	fmt.Printf("%v\n", res)
}

func dfs(row int, N int, queen [][]byte, attack [][]int) {
	if row == N {
		// append queen to res
		res = append(res, queen)
		return 
	}
	// 选择第i列
	for i := 0; i < N; i++ {
		if attack[row][i] == 0 {
			tmp := attack
			queen[row][i] = 'Q'
			dfs(row+1, N, queen, attack)
			queen[row][i] = ' '
			attack = tmp
		}
	}
}

func add_queen_attack(x, y int, attack [][]int) {
	var dx []int = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	var dy []int = []int{-1,  0,  1,-1, 1,-1, 0, 1}
	size := len(attack[0])
	for i := 0; i < 8; i++ {
		// inner to outer add attack positions
		for j := 1; j < size; j++ {
			nx := x + j * dx[i]
			ny := y + j * dy[i]
			if nx >= 0 && nx < size && ny >= 0 && ny < size {
				attack[nx][ny] = 1
			}
		}	
	}
}