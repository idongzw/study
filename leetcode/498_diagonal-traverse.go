/*
 * @Author: dzw
 * @Date: 2020-04-07 14:26:19
 * @Last Modified by: dzw
 * @Last Modified time: 2020-04-07 15:53:03
 */

/*
 给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。



示例:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

输出:  [1,2,4,7,5,3,6,8,9]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diagonal-traverse
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	fmt.Println(findDiagonalOrder([][]int{
		{1, 2},
		{4, 5},
		{7, 8},
	}))

	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))

	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3},
	}))

	fmt.Println(findDiagonalOrder([][]int{
		{1},
		{4},
		{7},
	}))
}

func findDiagonalOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	rst := make([]int, m*n)

	rst[0] = matrix[0][0] // start
	count := func(x, y int) int {
		if m < n {
			return m
		}
		return n
	}(m, n)

	j, k := 0, 1
	flag := true
	num := 1
	for i := 0; i < count; i++ {
		if flag {
			for ; j >= 0 && j < m && k >= 0 && k < n; j, k = j+1, k-1 {
				rst[num] = matrix[j][k]
				num++
			}
			if k < 0 {
				if m >= n {
					k = 0
				} else {
					k = 1
				}
			}
			if j >= m {
				j = m - 1
			}
			flag = false
		} else {
			for ; j >= 0 && j < m && k >= 0 && k < n; j, k = j-1, k+1 {
				rst[num] = matrix[j][k]
				num++
			}
			if j < 0 {
				if m >= n {
					j = 1
				} else {
					j = 0
				}
			}
			if k >= n {
				k = n - 1
			}
			flag = true
		}
		// j++
	}
	rst[m*n-1] = matrix[m-1][n-1] // end

	return rst
}
