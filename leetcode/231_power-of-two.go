/*
 * @Author: dzw
 * @Date: 2020-03-16 16:35:01
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-16 16:43:24
 */

/*
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:

输入: 1
输出: true
解释: 20 = 1
示例 2:

输入: 16
输出: true
解释: 24 = 16
示例 3:

输入: 218
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-two
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(2))
}

func isPowerOfTwo(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}
