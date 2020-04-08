/*
 * @Author: dzw
 * @Date: 2020-04-07 11:11:44
 * @Last Modified by: dzw
 * @Last Modified time: 2020-04-07 14:22:19
 */

/*
 给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

示例 1:
输入:

0 0 0
0 1 0
0 0 0
输出:

0 0 0
0 1 0
0 0 0
示例 2:
输入:

0 0 0
0 1 0
1 1 1
输出:

0 0 0
0 1 0
1 2 1
注意:

给定矩阵的元素个数不超过 10000。
给定矩阵中至少有一个元素是 0。
矩阵中的元素只在四个方向上相邻: 上、下、左、右。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/01-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))

	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}))

	fmt.Println(updateMatrix([][]int{
		{1, 0, 1},
		{1, 1, 1},
		{1, 1, 1},
	}))

	fmt.Println(updateMatrix([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}))

	fmt.Println(updateMatrix([][]int{
		{0, 1, 1, 0},
		{1, 0, 1, 1},
		{1, 1, 1, 1},
	}))

	fmt.Println(updateMatrix([][]int{
		{1, 1, 1, 0},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}))
}

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	dist := make([][]int, m)
	int_max := int(^uint(0) >> 1)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = int_max - 1 // int_max-1
		}
	}

	// left and top
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				dist[i][j] = 0
			} else {
				if i > 0 {
					dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
				}
				if j > 0 {
					dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
				}
			}
		}
	}

	// right and bottom
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < m-1 {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j < n-1 {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}

	return dist
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
