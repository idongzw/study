/*
 * @Author: dzw
 * @Date: 2020-03-19 10:42:22
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-19 11:28:33
 */

/*
 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7
*/

package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("ad"))
}

func longestPalindrome(s string) int {
	num := len(s)
	m := make(map[byte]int)
	for i := 0; i < num; i++ {
		m[s[i]]++
	}

	// 计算字符出现奇数的次数，最后构成回文只能出现一次个数为奇数的字符
	count := 0
	for _, v := range m {
		count += v % 2
	}

	if count == 0 { // 没有出现奇数次数的字符
		return num
	}

	return num - count + 1
}
