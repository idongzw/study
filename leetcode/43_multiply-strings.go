/*
 * @Author: dzw
 * @Date: 2020-04-06 13:39:08
 * @Last Modified by: dzw
 * @Last Modified time: 2020-04-06 14:38:32
 */

/*
 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/multiply-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(multiply("123", "24"))
	fmt.Println(multiply("0", "24"))
	fmt.Println(multiply("", "24"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)
	if len1 == 0 || len2 == 0 {
		return ""
	}

	rst := make([]byte, len1+len2)
	for i := len1; i > 0; i-- {
		idx := len2 + i - 1
		for j := len2; j > 0; j-- {
			tmp := (num1[i-1]-'0')*(num2[j-1]-'0') + rst[idx]
			rst[idx] = tmp % 10
			idx--
			rst[idx] += tmp / 10
		}
	}

	idx := -1
	for i := 0; i < len1+len2; i++ {
		if rst[i] != 0 && idx == -1 {
			idx = i
		}
		rst[i] += '0'
	}

	return string(rst[idx:])
}
