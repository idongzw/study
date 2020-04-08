/*
 * @Author: dzw
 * @Date: 2020-04-04 12:06:49
 * @Last Modified by: dzw
 * @Last Modified time: 2020-04-04 17:12:37
 */

/*
给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。

示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"
要求:

该字符串只包含小写的英文字母。
给定字符串的长度和 k 在[1, 10000]范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-string-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	//bacdfeg
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcdefg", 3))
	fmt.Println(reverseStr("abcdefg", 4))
	fmt.Println(reverseStr("abcdefg", 5))
	fmt.Println(reverseStr("abcdefg", 6))
	fmt.Println(reverseStr("abcdefg", 7))
	fmt.Println(reverseStr("abcdefg", 8))
}

func reverseStr(s string, k int) string {
	num := len(s)
	ss := []byte(s)

	left := 0
	for i := 0; left < num; i++ {
		left = i * 2 * k
		right := left + k - 1
		if right > num-1 { // 超过长度
			right = num - 1
		}

		// 反转
		for left < right {
			ss[left], ss[right] = ss[right], ss[left]
			left++
			right--
		}
	}

	return string(ss)
}
