/*
 * @Author: dzw
 * @Date: 2020-03-29 14:33:11
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-29 22:45:41
 */

/*
 你现在手里有一份大小为 N x N 的『地图』（网格） grid，上面的每个『区域』（单元格）都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地，你知道距离陆地区域最远的海洋区域是是哪一个吗？请返回该海洋区域到离它最近的陆地区域的距离。

我们这里说的距离是『曼哈顿距离』（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个区域之间的距离是 |x0 - x1| + |y0 - y1| 。

如果我们的地图上只有陆地或者海洋，请返回 -1。



示例 1：

1	0	1
0 	0 	0
1 	0 	1

输入：[[1,0,1],[0,0,0],[1,0,1]]
输出：2
解释：
海洋区域 (1, 1) 和所有陆地区域之间的距离都达到最大，最大距离为 2。

示例 2：

1	0	0
0	0	0
0	0	0

输入：[[1,0,0],[0,0,0],[0,0,0]]
输出：4
解释：
海洋区域 (2, 2) 和所有陆地区域之间的距离都达到最大，最大距离为 4。


提示：

1 <= grid.length == grid[0].length <= 100
grid[i][j] 不是 0 就是 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/as-far-from-land-as-possible
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(maxDistance([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}))
	fmt.Println(maxDistance([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}

func maxDistance(grid [][]int) int {
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	queueX := make([]int, 0, 10)
	queueY := make([]int, 0, 10)
	iNum := len(grid)
	jNum := len(grid[0])
	// 保存陆地
	for i := 0; i < iNum; i++ {
		for j := 0; j < jNum; j++ {
			if grid[i][j] == 1 {
				queueX = append(queueX, i)
				queueY = append(queueY, j)
			}
		}
	}

	hasOcean := false
	var point []int
	for len(queueX) > 0 {
		point = []int{queueX[0], queueY[0]}
		x := point[0]
		y := point[1]
		for i := 0; i < 4; i++ {
			newX := x + dx[i]
			newY := y + dy[i]
			if newX < 0 || newX >= iNum || newY < 0 || newY >= jNum || grid[newX][newY] != 0 {
				continue
			}
			grid[newX][newY] = grid[x][y] + 1
			hasOcean = true
			queueX = append(queueX, newX)
			queueY = append(queueY, newY)
		}
		queueX = queueX[1:]
		queueY = queueY[1:]
	}

	// 没有陆地或者海洋
	if point == nil || !hasOcean {
		return -1
	}

	return grid[point[0]][point[1]] - 1
}
