/*
 * @Author: dzw
 * @Date: 2020-03-27 20:41:20
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-27 20:42:58
 */

package fib

// Fib 计算第n个斐波那契数
func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}
