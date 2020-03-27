/*
 * @Author: dzw
 * @Date: 2020-03-18 14:48:20
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-18 15:06:41
 */

/*
 矩形以列表 [x1, y1, x2, y2] 的形式表示，其中 (x1, y1) 为左下角的坐标，(x2, y2) 是右上角的坐标。

如果相交的面积为正，则称两矩形重叠。需要明确的是，只在角或边接触的两个矩形不构成重叠。

给出两个矩形，判断它们是否重叠并返回结果。



示例 1：

输入：rec1 = [0,0,2,2], rec2 = [1,1,3,3]
输出：true
示例 2：

输入：rec1 = [0,0,1,1], rec2 = [1,0,2,1]
输出：false
*/

package main

import "fmt"

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
}

// 两个同时满足左下角不大于对方的右上角
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if len(rec1) < 4 || len(rec2) < 4 {
		return false
	}

	if (rec1[0] < rec2[2]) && (rec1[1] < rec2[3]) && (rec2[0] < rec1[2]) && (rec2[1] < rec1[3]) {
		return true
	}

	return false
}
