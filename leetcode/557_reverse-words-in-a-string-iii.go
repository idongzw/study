/*
 * @Author: dzw
 * @Date: 2020-04-04 17:09:18
 * @Last Modified by: dzw
 * @Last Modified time: 2020-04-04 17:29:11
 */

/*
 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

// 输入: "Let's take LeetCode contest"
// 输出: "s'teL ekat edoCteeL tsetnoc"
func reverseWords(s string) string {
	ss := []byte(s)
	ss = append(ss, ' ')

	start := 0
	end := 0
	for i, v := range ss {
		if v == ' ' {
			end = i - 1
			for start < end {
				ss[start], ss[end] = ss[end], ss[start]
				start++
				end--
			}
			start = i + 1
		}
	}

	return string(ss[:len(ss)-1])
}
