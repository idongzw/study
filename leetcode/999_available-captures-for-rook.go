/*
 * @Author: dzw
 * @Date: 2020-03-26 17:05:05
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-26 17:35:19
 */

/*

 */

package main

import "fmt"

func main() {
	board := [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'}}
	fmt.Println(numRookCaptures(board))
}

func numRookCaptures(board [][]byte) int {
	iNum := len(board)
	if iNum == 0 {
		return 0
	}
	jNum := len(board[0])

	// 上下左右四个方向
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	count := 0
	for i := 0; i < iNum; i++ {
		for j := 0; j < jNum; j++ {
			if board[i][j] == 'R' { //找到车
				// 上下左右四个方向判断
				for k := 0; k < 4; k++ {
					x, y := i, j
					for {
						x += dx[k]
						y += dy[k]
						if x < 0 || x >= iNum || y < 0 || y >= jNum || board[x][y] == 'B' {
							break
						}
						if board[x][y] == 'p' {
							count++
							break
						}
					}
				}
				return count
			}
		}
	}

	return count
}
