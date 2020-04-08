/*
 * @Author: dzw
 * @Date: 2020-03-30 09:20:14
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-30 09:58:02
 */

/*
 0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。

例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。



示例 1：

输入: n = 5, m = 3
输出: 3
示例 2：

输入: n = 10, m = 17
输出: 2


限制：

1 <= n <= 10^5
1 <= m <= 10^6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./", os.Args[0], "5 3")
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	m, _ := strconv.Atoi(os.Args[2])
	fmt.Println(lastRemaining(n, m))
}

func lastRemaining(n int, m int) int {
	s := 0
	for i := 2; i <= n; i++ {
		s = (s + m) % i
	}
	return s
}
