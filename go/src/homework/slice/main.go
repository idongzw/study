/*
 * @Author: dzw
 * @Date: 2020-02-24 19:55:45
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-24 20:46:23
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	{
		var a = make([]string, 5, 10)
		for i := 0; i < 10; i++ {
			a = append(a, fmt.Sprintf("%v", i))
		}
		fmt.Println(a) // [     0 1 2 3 4 5 6 7 8 9]
	}

	fmt.Println("----------------------------------")

	{
		var a = [...]int{3, 7, 8, 9, 1}
		fmt.Println(a) // [3 7 8 9 1]

		sort.Ints(a[:])
		fmt.Println(a) // [1 3 7 8 9]
	}
}
