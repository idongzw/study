/*
 * @Author: dzw
 * @Date: 2020-04-07 10:53:09
 * @Last Modified by: dzw
 * @Last Modified time: 2020-04-07 11:08:58
 */

/*
 给定一个矩阵 A， 返回 A 的转置矩阵。

矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。



示例 1：

输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：[[1,4,7],[2,5,8],[3,6,9]]
示例 2：

输入：[[1,2,3],[4,5,6]]
输出：[[1,4],[2,5],[3,6]]


提示：

1 <= A.length <= 1000
1 <= A[0].length <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/transpose-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

func transpose(A [][]int) [][]int {
	x, y := len(A), len(A[0])

	rst := make([][]int, y)
	for i := 0; i < y; i++ {
		rst[i] = make([]int, x)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			rst[j][i] = A[i][j]
		}
	}

	return rst
}

// 优化 x==y
func transpose1(A [][]int) [][]int {
	x, y := len(A), len(A[0])

	if x == y {
		for i := x - 1; i >= 0; i-- {
			for j := 0; j < i; j++ {
				A[i][j], A[j][i] = A[j][i], A[i][j]
			}
		}

		return A
	}

	rst := make([][]int, y)
	for i := 0; i < y; i++ {
		rst[i] = make([]int, x)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			rst[j][i] = A[i][j]
		}
	}

	return rst
}
