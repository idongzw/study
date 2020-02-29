/*
 * @Author: dzw
 * @Date: 2020-02-24 18:33:06
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-24 19:29:12
 */

package main

import "fmt"

func main() {
	//1. 求数组{1, 3, 5, 7, 8}所有元素的值
	{
		a := [...]int{1, 3, 5, 7, 8}
		sum := 0
		for _, v := range a {
			sum += v
		}

		fmt.Println("sum =", sum)
	}

	//2. 找出数组中和为指定值的两个元素的下标，
	//比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	{
		a := [...]int{1, 3, 5, 7, 8}
		sum := 8
		num := len(a)
		for i := 0; i < num; i++ {
			for j := i + 1; j < num; j++ {
				if a[i]+a[j] == sum {
					fmt.Println("i =", i, ",j =", j)
				}
			}
		}
	}
}
