/*
 * @Author: dzw
 * @Date: 2020-03-06 14:00:14
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-06 15:40:04
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f(i int, sw *sync.WaitGroup) {
	defer sw.Done()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println("f i =", i)
}

// GOMAXPROCS
func a(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("A:", i)
	}
}

func b(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("B:", i)
	}
}

// channel 练习
// 1. 启动一个goroutine，生成100个数发送到ch1
// 2. 启动一个goroutine，从ch1中取值，计算其平方放到ch2中
// 3. 在main中从ch2取值打印

func main() {
	// sw := &sync.WaitGroup{}
	// for i := 0; i < 10; i++ {
	// 	sw.Add(1)
	// 	go f(i, sw)
	// }
	// sw.Wait()

	// runtime.GOMAXPROCS(2) // 默认是cpu的逻辑核心数，默认跑满cpu
	// wg := &sync.WaitGroup{}
	// wg.Add(2)
	// go a(wg)
	// go b(wg)
	// wg.Wait()

	// fmt.Println("runtime num cpu:", runtime.NumCPU())

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go func() {
		for i := 1; i < 101; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for v := range ch1 {
			ch2 <- v * v
		}
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println("result =", v)
	}
}
