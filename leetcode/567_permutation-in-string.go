/*
 * @Author: dzw
 * @Date: 2020-04-04 19:00:32
 * @Last Modified by: dzw
 * @Last Modified time: 2020-04-04 22:28:51
 */

/*
 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").


示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False


注意：

输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

func checkInclusion(s1 string, s2 string) bool {
	num1 := len(s1)
	num2 := len(s2)

	if num1 == 0 {
		return true
	}

	if num1 > num2 {
		return false
	}

	a1 := [26]int{}
	a2 := [26]int{}

	for i := 0; i < num1; i++ {
		a1[s1[i]-'a']++
		a2[s2[i]-'a']++
	}

	if a1 == a2 {
		return true
	}

	// 滑动窗口
	for i := num1; i < num2; i++ {
		a2[s2[i]-'a']++
		a2[s2[i-num1]-'a']--
		if a1 == a2 {
			return true
		}
	}

	return false
}
