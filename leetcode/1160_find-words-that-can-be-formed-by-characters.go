/*
 * @Author: dzw
 * @Date: 2020-03-17 15:04:38
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-17 15:39:12
 */

/*
给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。

假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。

注意：每次拼写时，chars 中的每个字母都只能用一次。

返回词汇表 words 中你掌握的所有单词的 长度之和。



示例 1：

输入：words = ["cat","bt","hat","tree"], chars = "atach"
输出：6
解释：
可以形成字符串 "cat" 和 "hat"，所以答案是 3 + 3 = 6。
示例 2：

输入：words = ["hello","world","leetcode"], chars = "welldonehoneyr"
输出：10
解释：
可以形成字符串 "hello" 和 "world"，所以答案是 5 + 5 = 10。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countCharacters1([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
	fmt.Println(countCharacters2([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}

func countCharacters1(words []string, chars string) int {
	count := 0
	wordLen := 0

	for _, word := range words {
		wordLen = len(word)
		tmp := chars
		flag := true
		for i := 0; i < wordLen; i++ {
			idx := strings.IndexByte(tmp, word[i])
			if idx == -1 {
				flag = false
				break
			} else {
				tmp = tmp[:idx] + tmp[idx+1:]
			}
		}
		if flag {
			count += wordLen
		}

	}

	return count
}

func countCharacters2(words []string, chars string) int {
	charsCount := [26]int{}
	for _, c := range chars {
		charsCount[c-'a']++ // 储存每个字母出现次数
	}

	count := 0

	for _, word := range words {
		match := true
		bcc := charsCount
		for _, c := range word {
			if bcc[c-'a'] <= 0 {
				match = false
				break
			}
			bcc[c-'a']--
		}

		if match {
			count += len(word)
		}
	}

	return count
}
