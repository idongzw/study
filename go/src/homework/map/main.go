/*
 * @Author: dzw
 * @Date: 2020-02-25 10:59:43
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-25 11:03:11
 */

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//统计字符串中每个单词出现的次数
	// eg. how do you do 中 how = 1 do = 2 you = 1
	{
		s := "how do you do"
		ss := strings.Split(s, " ")
		m := make(map[string]int)
		for _, v := range ss {
			if _, ok := m[v]; ok {
				m[v]++
			} else {
				m[v] = 1
			}
		}

		for k, v := range m {
			fmt.Println(k, "=", v)
		}
	}

	{
		type Map map[string][]int
		m := make(Map)
		s := []int{1, 2}
		s = append(s, 3)
		fmt.Printf("%+v, %p\n", s, &s)
		/*
			在s变量被追加1个元素3之后，就把这个切片的容量赋值给了map，当时容量是3，
			那么以后怎么搞这个变量s，对于map来说，它只会受到前3个元素的影响，3个元素之后的变化，m是不再受到影响
		*/
		m["q1mi"] = s
		sm := m["q1mi"]
		fmt.Printf("%+v, %p, len(sm) = %d\n", m["q1mi"], &sm, len(sm))
		s = append(s[:1], s[2:]...)
		fmt.Printf("%+v, %p, len(s) = %d\n", s, &s, len(s))
		sm = m["q1mi"]
		fmt.Printf("%+v, %p, len(sm) = %d\n", m["q1mi"], &sm, len(sm)) // 1 3 3
	}

	// 判断字符串中汉字的数量
	{
		s := "hello 汉字成都市城市v"
		for _, v := range s {
			fmt.Printf(string(v))
		}
		fmt.Print("\n")

		count := 0
		for _, v := range s {
			if unicode.Is(unicode.Han, v) {
				count++
			}
		}

		fmt.Println("汉字的数量是：", count)
	}

	// 回文判断
	{
		s := "s上海自来s水s来自海上s"
		// 先转成 []rune
		sr := make([]rune, 0, len(s))

		// 中文不是占用一个字节
		// num_s := len(s)
		// fmt.Println("len(s) = ", num_s)
		// for i := 0; i < num_s; i++ {
		// 	fmt.Println("s[i] =", s[i], " s[num_s-(i+1)] =", s[num_s-(i+1)])
		// 	if s[i] != s[num_s-(i+1)] {
		// 		fmt.Println("不是回文")
		// 		return
		// 	}
		// }
		// fmt.Println("是回文")

		for i, v := range s { //  i 是字符的字节位置，v 是字符的拷贝，有中文字符需要转为rune
			fmt.Println("i = ", i)
			sr = append(sr, v)
		}
		num := len(sr)
		fmt.Println("len(sr) = ", len(sr))

		for i := range sr {
			if sr[i] != sr[num-(i+1)] {
				fmt.Println("不是回文")
				return
			}
		}
		fmt.Println("是回文")
	}
}
