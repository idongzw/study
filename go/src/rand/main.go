/*
 * @Author: dzw
 * @Date: 2020-03-06 14:10:29
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-06 14:13:07
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 保证每次的随机数都不一样
	for i := 0; i < 10; i++ {
		r1 := rand.Int()    // int64
		r2 := rand.Intn(10) // 0 <= r2 < 10
		fmt.Println("r1 =", r1, ",r2 =", r2)
	}
}
