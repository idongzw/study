/*
* @File Name: concurrency.go
* @Author: idongzw
* @Date:   2020-02-17 21:09:08
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-18 18:26:50
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	/*
	   {
	       go Go()
	       time.Sleep(2 * time.Second)
	   }
	*/

	/*
	   channel
	   channel 是 goroutine 沟通的桥梁，大都是阻塞同步的
	   通过make创建，close关闭
	   channel是引用类型
	   可以使用for range 来迭代不断操作 channel
	   可以设置单向或双向通道
	   可以设置缓存大小，在未被填满前不会发生阻塞
	*/
	{
		c := make(chan bool) //无缓存，阻塞
		go func() {
			fmt.Println("chan test1")
			c <- true // 写
		}()

		<-c // 读
	}

	fmt.Println("--------------------")

	{
		c := make(chan bool) //无缓存，阻塞
		go func() {
			fmt.Println("chan test2")
			<-c // 读
		}()

		c <- true // 写
	}

	fmt.Println("--------------------")

	{
		c := make(chan bool)
		go func() {
			fmt.Println("chan test3")
			c <- true // 写
			close(c)
		}()

		// 读
		for v := range c {
			fmt.Println(v)
		}
	}

	fmt.Println("--------------------")

	{
		c := make(chan bool, 1) // 有缓存，异步
		go func() {
			fmt.Println("chan test4")
			c <- true // 写不阻塞
			close(c)
		}()

		<-c // 读阻塞
	}

	fmt.Println("--------------------")

	{
		c := make(chan bool, 1) // 有缓存，异步
		go func() {
			fmt.Println("chan test5")
			<-c // 读
		}()

		// 缓存未被填满之前不会阻塞
		c <- true // 写不阻塞
	}

	//time.Sleep(2 * time.Second)

	fmt.Println("--------------------")

	{
		fmt.Println("NumCPU:", runtime.NumCPU())
		runtime.GOMAXPROCS(runtime.NumCPU()) // 设置使用cpu核数
		c := make(chan bool, 10)
		for i := 0; i < 10; i++ {
			go GoChannel(c, i)
		}

		for i := 0; i < 10; i++ {
			<-c
		}
	}

	fmt.Println("--------------------")

	{
		fmt.Println("NumCPU:", runtime.NumCPU())
		runtime.GOMAXPROCS(runtime.NumCPU()) // 设置使用cpu核数

		wg := sync.WaitGroup{}
		wg.Add(10)

		for i := 0; i < 10; i++ {
			go GoSync(&wg, i)
		}

		wg.Wait()
	}

	fmt.Println("++++++++++++++++++++++")

	// select
	// 可以处理一个或者多个channel的发送和接受
	// 同时有多个可用的channel时，按随机顺序处理
	// 可用空的 select 来阻塞main函数
	// 可设置超时
	{
		c1, c2 := make(chan int), make(chan string)
		o := make(chan bool, 2)

		go func() {
			for {
				select {
				case v, ok := <-c1:
					if !ok {
						//没办法判断两个都关闭了，只能一个关闭就退出
						o <- true
						fmt.Println("c11111111111", v)
						break
					}
					fmt.Println("c1 =", v, ok)
				case v, ok := <-c2:
					if !ok {
						//没办法判断两个都关闭了，只能一个关闭就退出
						o <- true
						fmt.Println("c22222222222", v)
						break
					}
					fmt.Println("c2 =", v, ok)
				}
			}
		}()

		c1 <- 1
		c2 <- "hi"
		c1 <- 2
		c2 <- "hello"

		// 只要有一个 channel 被关闭时(select还会读channel的值)，select 就应该退出，
		// 不然被关闭那个按照上述程序会是死循环状态，影响另一个channel
		close(c1)
		close(c2)

		for i := 0; i < 2; i++ {
			<-o
		}

		//for {}
	}

	fmt.Println("===============================")

	{
		c1, c2 := make(chan int), make(chan string)
		o := make(chan bool, 2)

		go func() {
			a, b := false, false
			for {
				select {
				case v, ok := <-c1:
					if !ok {
						if !a { // 保证只写 channel o一次
							o <- true
							a = true
							fmt.Println("c!!!!!!!!!!!!!!!!!!", v)
						}
						break
					}
					fmt.Println("c1 =", v, ok)
				case v, ok := <-c2:
					if !ok {
						if !b { // 保证只写 channel o一次
							o <- true
							b = true
							fmt.Println("c@@@@@@@@@@@@@@@@", v)
						}
						break
					}
					fmt.Println("c2 =", v, ok)
				}
			}
		}()

		c1 <- 1
		c2 <- "hi"
		c1 <- 2
		c2 <- "hello"

		// 只要有一个 channel 被关闭时(select还会读channel的值)，select 就应该退出，
		// 不然被关闭那个按照上述程序会是死循环状态，影响另一个channel
		// 没办法判断两个都关闭了
		close(c1)
		close(c2)

		for i := 0; i < 2; i++ {
			<-o
		}

		//for {}
	}

	fmt.Println("--------------------------------")

    /*
	// 随机输出 0 1
	{
		c := make(chan int)

		go func() {
			for v := range c {
				fmt.Println(v)
			}
		}()

		for {
			select {
			case c <- 0:
			case c <- 1:
			}
		}
	}
    */
    
    fmt.Println("--------------------------------")

    // timeout
    {
        c := make(chan bool)
        select {
        case v := <-c:
            fmt.Println(v)
        case <-time.After(3 * time.Second):
            fmt.Println("Timeout")
        }
    }

    {
        c := make(chan string)
        go GoPingpang(c)
        for i := 0; i < 10; i++ {
            c <- fmt.Sprintf("From main: Hello, #%d", i)
            fmt.Println(<-c)
        }
    }
}

func Go() {
	fmt.Println("Go Go Go!!!")
}

func GoChannel(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}

func GoSync(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}

func GoSelect() {

}

func GoPingpang(c chan string) {
    i := 0
    for {
        fmt.Println(<-c)
        c <- fmt.Sprintf("From GoPingpang: Hi, #%d", i)
        i++
    }
}