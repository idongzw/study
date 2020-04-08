/*
 * @Author: dzw
 * @Date: 2020-04-07 09:09:50
 * @Last Modified by: dzw
 * @Last Modified time: 2020-04-07 10:45:03
 */

/*
 给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？



示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-matrix-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	count := len(matrix)

	/*
	   [1,2,3]       [1,4,7]
	   [4,5,6]  ---> [2,5,8]
	   [7,8,9]       [3,6,9]
	*/

	for i := count - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	/*
	   [1,4,7]       [7,4,1]
	   [2,5,8]  ---> [8,5,2]
	   [3,6,9]       [9,6,3]
	*/

	for i := 0; i < count; i++ {
		for j, k := 0, count-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
