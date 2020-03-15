/*
 * @Author: dzw
 * @Date: 2020-03-06 15:47:11
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-06 16:29:22
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序。
	1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	3. 主goroutine从resultChan取出结果并打印到终端输出
*/

func createNum(c chan<- int64) {
	for i := 0; i < 100; i++ {
		c <- int64(rand.Intn(900) + 100)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func getNum(cr <-chan int64, cw chan<- int64, wg *sync.WaitGroup, workid int) {
	defer wg.Done()

	for v := range cr {
		var rst int64 = 0
		for v > 0 {
			rst += v % 10
			v /= 10
		}
		cw <- rst
		fmt.Println("worker id:", workid, ",rst:", rst)
	}
}

func main() {
	jobChan := make(chan int64)
	rstChan := make(chan int64)

	// 开启1个goroutine
	go createNum(jobChan)

	// 开启24个goroutine
	wg := &sync.WaitGroup{}
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go getNum(jobChan, rstChan, wg, i)
	}
	go func() {
		wg.Wait()
		close(rstChan)
	}()

	i := 0
	for v := range rstChan {
		fmt.Println("rst:", v)
		i++
	}
	fmt.Println(i)
}
