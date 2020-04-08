/*
 * @Author: dzw
 * @Date: 2020-04-06 09:32:10
 * @Last Modified by: dzw
 * @Last Modified time: 2020-04-06 13:16:51
 */

/*
 给定一个字符串，逐个翻转字符串中的每个单词。



示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


说明：

无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


进阶：

请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world! "))
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	ss := []byte(s)
	count := len(ss)
	if count == 0 {
		return ""
	}
	// s revers
	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	ss = append(ss, ' ')
	count++

	// c revers
	start := 0
	for i := 0; i < count; i++ {
		if ss[i] == ' ' {
			for j, k := start, i-1; j < k; j, k = j+1, k-1 {
				ss[j], ss[k] = ss[k], ss[j]
			}
			start = i + 1
		}
	}

	// rm space
	sss := []byte{}
	num := 0
	flag := true // 开头空格
	for i := 0; i < count; i++ {
		if ss[i] == ' ' && !flag {
			num++
		}

		if ss[i] != ' ' {
			flag = false
			if num > 0 {
				num = 0
				sss = append(sss, ' ', ss[i])
			} else {
				sss = append(sss, ss[i])
			}
		}
	}

	return string(sss)
}
