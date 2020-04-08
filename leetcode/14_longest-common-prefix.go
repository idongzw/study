/*
 * @Author: dzw
 * @Date: 2020-03-30 11:35:53
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-30 12:59:57
 */

/*
 14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "dg", "dog"}))
	fmt.Println(longestCommonPrefix([]string{"og", "dg", "dog"}))
}

func longestCommonPrefix(strs []string) string {
	num := len(strs)
	if num < 1 {
		return ""
	}
	if num == 1 {
		return strs[0]
	}

	s := strs[0]
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		for j := 1; j < num; j++ {
			if i >= len(strs[j]) || (i < len(strs[j]) && strs[j][i] != c) {
				return s[:i]
			}
		}
	}

	return s[:i]
}
