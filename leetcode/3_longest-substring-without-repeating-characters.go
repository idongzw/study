/*
 * @Author: dzw
 * @Date: 2020-03-16 11:08:19
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-30 11:31:18
 */

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("adsadadsdsfafafrgeqgteabgbthrynwyrhnjnjutmt"))

}

func lengthOfLongestSubstring(s string) int {
	num := len(s)
	if num == 0 {
		return 0
	}
	max := 0
	left := 0
	right := 0
	s1 := s[left:right] // 不断扩大s1

	for ; right < num; right++ {
		if idx := strings.IndexByte(s1, s[right]); idx != -1 { //查找s1中是否有当前字符
			left += idx + 1 // 遇见重复字符时，取第一次出现字符之后的字符串
		}
		s1 = s[left : right+1] // 取当前最长的字符串

		if len(s1) > max {
			max = len(s1)
		}
	}
	return max
}

func maxForInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring1(s string) int {
	i := 0
	max := 0

	//abcabcbb
	sr := []rune(s)
	for idx, c := range sr {
		for j := i; j < idx; j++ {
			if sr[j] == c { // 检测当前字符和之前的有没有相同，相同则后移
				i = j + 1
			}
		}
		// 当前最长
		if idx-i+1 > max {
			max = idx - i + 1
		}
	}

	return max
}
