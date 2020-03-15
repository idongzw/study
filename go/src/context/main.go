/*
 * @Author: dzw
 * @Date: 2020-03-07 19:42:37
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-07 21:51:22
 */

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 子goroutine优雅退出
func worker(ctx context.Context) {
	defer wg.Done()
	// 当子goroutine又开启另外一个goroutine时，只需要将ctx传入即可
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待父goroutine通知
			break LOOP // 收到通知结束循环
		default:
		}
	}
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待父goroutine通知
			break LOOP // 收到通知结束循环
		default:
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	fmt.Printf("n address = %p, n = %v\n", &n, n)
	fmt.Printf("dst address = %p, n = %#v\n", dst, dst)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
				fmt.Printf("anonymous func n address = %p, n = %v\n", &n, n)
			}
		}
	}()
	return dst
}

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// wg.Add(1)
	// go worker(ctx)
	// time.Sleep(time.Second * 3)
	// cancel() // 通知子goroutine退出
	// wg.Wait()
	// fmt.Println("end...")

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel() // 当我们取完需要的整数后调用cancel

	// for n := range gen(ctx) {
	// 	fmt.Println(n)
	// 	if n == 5 {
	// 		break
	// 	}
	// }

	// context.WithDeadline 到时间自动cancel,但是还是尽量调用cancel函数
	d := time.Now().Add(200 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second): // 延迟1s执行
		fmt.Println("1s timeout")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
