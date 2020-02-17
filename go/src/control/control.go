/*
* @File Name: control.go
* @Author: idongzw
* @Date:   2020-02-15 18:22:19
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-15 20:08:27
*/
package main

import (
    "fmt"
)

func main() {
    i := 12

    if i := 1; i > 0 {
        fmt.Println(i)
    }

    j := 1
    if j > 0 {
        fmt.Println(j)
    }

    fmt.Println(i)

    // go 只有 for 一个循环语句，支持 3 种形式
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    a := 0
    for {
        if a > 3 {
            break
        }
        fmt.Println(a)
        a++
    }

    b := 0
    for b <= 3 {
        fmt.Println(b)
        b++
    }

    // switch
    // 可以使用任何类型或表达式作为条件语句
    // 不需要写 break，条件符合自动终止
    // 如果希望继续执行下一个case，需使用 fallthrough 语句
    // 支持一个初始化表达式（可以是并行方式），右侧需跟分号
    // 左大括号必须和条件语句在同一行
    c := 0
    switch c {
    case 0:
        fmt.Println("c = 0")
    case 1:
        fmt.Println("c = 1")
    default:
        fmt.Println("None")
    }
    fmt.Println(c)

    c = 1
    switch {
    case c >= 0:
        fmt.Println("c >= 0")
        fallthrough
    case c >= 1:
        fmt.Println("c >= 1")
    default:
        fmt.Println("None")
    }
    fmt.Println(c)

    switch d := 1; { // d 局部变量
    case d >= 0:
        fmt.Println("d >= 0")
        fallthrough
    case d >= 1:
        fmt.Println("d >= 1")
    default:
        fmt.Println("None")
    }

    //跳转语句 goto, break, continue
    //三个语法都可以配合标签使用
    //标签名区分大小写，若不使用会造成编译错误
    //break与continue配合标签可用于多层循环的跳出
    //goto 是调整执行位置，与其他两个语句配合标签的结果并不相同
    fmt.Println("-------------------------------------- ")
    LABEL1:
    for {
        for i := 0; i < 10; i++ {
            if i > 3 {
                break LABEL1 // 外层无限循环跳出
                //goto LABEL1 //回到 LABEL1 执行
            }
            fmt.Println(i)
        }
    }

    fmt.Println("Ok1")

    for {
        for i := 0; i < 10; i++ {
            if i > 3 {
                goto LABEL2
            }
            fmt.Println(i)
        }
    }
    LABEL2:
    fmt.Println("Ok2")

    // continue 用在有限的循环中
    LABEL3:
    for i := 0; i < 10; i++ {
        for {
            continue LABEL3
        }
        fmt.Println(i)
    }
    fmt.Println("Ok3")

    LABEL4:
    for i := 0; i < 10; i++ {
        for {
            fmt.Println(i)
            continue  LABEL4
            //goto LABEL4 //循环打印0
        }
    }
}