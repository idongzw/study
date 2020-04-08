/*
 * @Author: dzw
 * @Date: 2020-03-28 10:12:22
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-28 10:14:15
 */

package main

import "fmt"

func main() {
	x, y := 1, 2
	fmt.Println("x =", x, ",y =", y)
	x, y = y, x // swap
	fmt.Println("x =", x, ",y =", y)
}
