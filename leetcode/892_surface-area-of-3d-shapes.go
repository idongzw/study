/*
 * @Author: dzw
 * @Date: 2020-03-25 10:08:16
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-25 11:03:40
 */

/*
 在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。

请你返回最终形体的表面积。



示例 1：

输入：[[2]]
输出：10
示例 2：

输入：[[1,2],[3,4]]
输出：34
示例 3：

输入：[[1,0],[0,2]]
输出：16
示例 4：

输入：[[1,1,1],[1,0,1],[1,1,1]]
输出：32
示例 5：

输入：[[2,2,2],[2,1,2],[2,2,2]]
输出：46


提示：

1 <= N <= 50
0 <= grid[i][j] <= 50

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surface-area-of-3d-shapes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(surfaceArea([][]int{{2}}))
	fmt.Println(surfaceArea([][]int{{1, 2}, {3, 4}}))
	fmt.Println(surfaceArea([][]int{{1, 0}, {0, 2}}))
	fmt.Println(surfaceArea([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	fmt.Println(surfaceArea([][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}))
}

func surfaceArea(grid [][]int) int {
	iNum := len(grid)
	if iNum == 0 {
		return 0
	}
	jNum := len(grid[0])
	// fmt.Println(iNum, jNum)

	area := 0
	for i := 0; i < iNum; i++ {
		for j := 0; j < jNum; j++ {
			if grid[i][j] == 0 { // == 0 跳过
				continue
			}
			area += grid[i][j]*4 + 2

			// 去除掉相邻覆盖掉的
			if i > 0 && grid[i-1][j] > 0 {
				area -= min(grid[i-1][j], grid[i][j]) * 2
			}
			if j > 0 && grid[i][j-1] > 0 {
				area -= min(grid[i][j-1], grid[i][j]) * 2
			}
		}
	}

	return area
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
