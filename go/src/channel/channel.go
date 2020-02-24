/*
* @File Name: channel.go
* @Author: idongzw
* @Date:   2020-02-22 11:18:07
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-22 16:43:57
*/
package main

import (
    "fmt"
    "time"
)

/*
channel 通道在使用的时候，有以下几个注意点：
    1. 用于goroutine，传递消息的
    2. 通道，每个都有相关联的数据类型
        nil chan 不能使用，类似于 nil map，不能直接存储键值对
    3. 使用通道传递数据 <-
    4. 阻塞(无缓冲区)：
        发送数据： chan <- data 阻塞的，直到另一条goroutine，读取数据来解除阻塞
        读取数据： data <- chan 阻塞的，直到另一条goroutine，写入数据来解除阻塞
    5. channel本身就是同步的，意味着同一时间，只能有一条goroutine来操作
    通道是goroutine之间的连接，所以通道的发送和接收必须处在不同的goroutine中

关闭通道：
    发送者可以关闭通道，来通知接收方不会有更多的数据发送到channel上
    close(channel)
    
    接收者可以在接收来自通道的数据时使用额外的变量来检测通道是否已经关闭
        语法结构：
            v, ok := <-ch
        如果ok的值为true，表示成功的从通道读取了一个数据value
        如果ok的值为false，这意味着我们正在从一个封闭的通道读取数据。
    从关闭的通道读取的值将是通道类型的零值
 */

func main() {
    {
        var c chan int
        fmt.Printf("channel %T, %v\n", c, c) // channel chan int, <nil>

        if c == nil {
            fmt.Println("channel is nil, need make")
            c = make(chan int)
        }
        fmt.Printf("make channel %T, %v\n", c, c)
    }

    fmt.Println("------------------------------")

    {
        c := make(chan bool)

        go func () {
            for i := 0; i < 10; i++ {
                fmt.Println("son goroutine, i:", i)
            }
            fmt.Println("son over...")
            c <- true // 写数据
        }()

        data := <-c // 读数据，阻塞

        fmt.Println("main read data:", data)
        fmt.Println("1. main over...")
    }

    fmt.Println("------------------------------")

    {
        //time.Sleep(2 * time.Second)
        c1 := make(chan int)

        go func () {
            //time.Sleep(2 * time.Second)
            fmt.Println("son read data before")
            data := <-c1
            fmt.Println("son read data:", data) 
        }()

        fmt.Println("write data before")
        //time.Sleep(3 * time.Second)
        c1 <- 10
        fmt.Println("2. main over...")
    }

    fmt.Println("------------------------------")

    {
        c := make(chan int)

        // send data
        go func () {
            for i := 0; i < 10; i++ {
                c <- i
            }
            close(c)
        }()

        // read data
        for {
            if v, ok := <-c; !ok { //close之后，ok为false
                fmt.Println("read data over", ok)
                break
            } else {
                fmt.Println("read data:", v, ok)
            }
            //time.Sleep(1 * time.Second)
        }

        fmt.Println("3. main over...")
    }

    fmt.Println("------------------------------")

    // range 访问通道
    {
        c := make(chan int)

        go func () {
            for i := 0; i < 10; i++ {
                c <- i
                //time.Sleep(1 * time.Second)
            }
            close(c)
        }()

        for v := range c {
            fmt.Println("read data:", v)
        }

        fmt.Println("4. main over...")
    }

    fmt.Println("------------------------------")

    /*
    非缓冲通道：make(chan T)
        一次发送，一次接收，都是阻塞的
    缓冲通道：make(chan T, capacity)
        发送：缓冲区的数据满了，才会阻塞
        接收：缓冲区的数据空了，才会阻塞
     */
    {
        c1 := make(chan int) // 非缓冲通道
        fmt.Println(len(c1), cap(c1)) // 0 0
        //阻塞式的，需要有其他的goroutine解除阻塞，否则deadlock
        //c1 <- 100 // fatal error: all goroutines are asleep - deadlock!

        c2 := make(chan int, 5) // 缓冲通道，缓冲区大小是5
        fmt.Println(len(c2), cap(c2)) // 0 5

        c2 <- 100
        fmt.Println(len(c2), cap(c2)) // 1 5
        c2 <- 200
        c2 <- 300
        c2 <- 400
        c2 <- 500
        fmt.Println(len(c2), cap(c2)) // 5 5
        // 缓冲通道满，阻塞
        // c2 <- 600 // fatal error: all goroutines are asleep - deadlock!
    }

    fmt.Println("------------------------------")

    {
        c := make(chan int, 4)

        go sendData(c)

        for v := range c {
            fmt.Println("read data:", v)
        }
    }

    fmt.Println("------------------------------")

    /*
    双向通道：
        chan T
            chan <- data: 发送数据，写出
            data <- chan: 获取数据，读取

    单向通道：
        chan <- T :只支持写
        <- chan T :只支持读
     */
    
    // 双向通道
    {
        c := make(chan string)
        done := make(chan bool)

        go func () {
            c <- "son goroutine data"
            data := <-c
            fmt.Println("from main goroutine data:", data)
            done <- true
        }()

        data := <-c
        fmt.Println("from son goroutine data:", data)
        c <- "main goroutine data"

        <-done
        fmt.Println("5. main over...")
    }

    fmt.Println("------------------------------")

    // 单向通道
    {
        //c1 := make(chan <- int) //单向通道，只能写，不能读 
        //c1 <- 100
        //<-c1 // invalid operation: <-c1 (receive from send-only type chan<- int)
        
        //c2 := make(<- chan int) // 单向通道，只能读，不能写
        //<-c2
        //c2 <- 100 // invalid operation: c2 <- 100 (send to receive-only type <-chan int) 
        
        c3 := make(chan int)
        done := make(chan bool)

        // 单向通道主要用于函数参数，限定参数内部只读或只写
        go writeOnly(c3, done)
        go readOnly(c3, done)

        for i := 0; i < 2; i++ {
            fmt.Println("done...", i, <-done)
        }

        fmt.Println("6. main over...")
    }

    fmt.Println("------------------------------")

    // time channel
    {
        timer := time.NewTimer(3 * time.Second)
        fmt.Printf("%T\n", timer) // *time.Timer
        fmt.Println(time.Now()) // 2020-02-22 16:33:44.863651961 +0800 CST m=+0.002492743
        c := timer.C
        // 此处等待channel中的值，会阻塞3s
        fmt.Println(<-c) // 2020-02-22 16:33:47.86371565 +0800 CST m=+3.002556483
    }

    fmt.Println("------------------------------")

    // stop timer
    {
        t1 := time.NewTimer(5 * time.Second)
        // goroutine
        go func () {
            <- t1.C
            fmt.Println("timer is over...")
        }()

        time.Sleep(3 * time.Second)
        flag := t1.Stop()
        if flag {
            fmt.Println("timer is stop")
        }
    }

    fmt.Println("------------------------------")

    // After
    // func After(d Duration) <-chan Time
    // 返回一个通道：chan，存储的是 d 时间间隔之后的当前时间
    // 相当于：return NewTimer(d).C
    {
        ch := time.After(3 * time.Second)
        fmt.Printf("%T\n", ch) // <-chan time.Time
        fmt.Println(time.Now()) // 2020-02-22 16:42:35.333657128 +0800 CST m=+6.003071254

        t := <-ch
        fmt.Println(t) // 2020-02-22 16:42:38.333689947 +0800 CST m=+9.003104021
    }    
}

func sendData(c chan int) {
    for i := 0; i < 10; i++ {
        c <- i
        fmt.Println("send data:", i)
    }
    close(c)
}

// 通道c只写
func writeOnly(c chan <- int, done chan <- bool) {
    c <- 100
    fmt.Println("write data")
    done <- true
    //<-c // invalid operation: <-c (receive from send-only type chan<- int)
}

// 通道c只读
func readOnly(c <- chan int, done chan <- bool) {
    data := <-c
    fmt.Println("read data:", data)
    done <- true
    //c <- 100 // invalid operation: c <- 100 (send to receive-only type <-chan int)
}