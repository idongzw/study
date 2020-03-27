/*
 * @Author: dzw
 * @Date: 2020-03-16 15:52:02
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-16 15:58:02
 */

/*
编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。
*/

package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(4))
}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}

	return count
}
