/*
* @File Name: control.go
* @Author: idongzw
* @Date:   2020-02-15 18:22:19
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-19 17:38:24
*/
package main

import (
    "fmt"
    "math"
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

    // 1-100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("1-100 sum:", sum)

    //1-100 能被 3 整除，不能被 5 整除的数字，
    //统计被打印数字的个数，每行打印 5 个
    num := 0
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 && i % 5 != 0 {
            num++
            fmt.Printf("%d\t", i)
            if num % 5 == 0 {
                fmt.Printf("\n")
            }
        }
    }
    fmt.Println("\nnum = :", num)

    for i := 0; i < 5; i++ {
        for i := 0; i < 5; i++ {
            fmt.Print("*")
        }
        fmt.Print("\n")
    }

    // 九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d\t", j, i, i * j)
        }
        fmt.Print("\n")
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

    /**
     * fallthrough:用于穿透switch
     *     当switch中某个case匹配成功后，就执行该case语句
     *     如果遇到fallthrough，那么后边紧邻的case，无需匹配，穿透执行
     *     fallthrough 应该位于某个case的最后一行
     */

    // 省略 switch 后边的变量，相当于直接作用在 true 上
    c = 0
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

    switch d := 1; { // d 局部变量，只作用于 switch
    case d >= 0:
        fmt.Println("d >= 0")
        fallthrough
    case d >= 1:
        fmt.Println("d >= 1")
    default:
        fmt.Println("None")
    }

    // case 可以同时测试多个可能符合条件的值，使用逗号分割
    /*
    一个月的天数：
    1,3,5,7,8,10,12
        31
    4,6,9,11
        30
    2 : 28/29
     */
    month := 5
    day := 0
    year := 2020
    switch month {
    case 1,3,5,7,8,10,12:
        day = 31
    case 4,6,9,11:
        day = 30
    case 2:
        if year % 400 == 0 || 
        (year % 4 == 0 && year % 100 != 0) {
            day = 29
        } else {
            day = 28
        }
    default:
        fmt.Println("month error")
    }
    fmt.Println("month", month, "have", day, "days")

    //switch break
    switch n := 2; n {
    case 1:
        fmt.Println("11111")
        fmt.Println("11111")
        fmt.Println("11111")
    case 2:
        fmt.Println("22222")
        fmt.Println("22222")
        break // 用于强制结束case，意味着switch被强制结束
        fmt.Println("22222")
    case 3:
        fmt.Println("33333")
        fmt.Println("33333")
        fmt.Println("33333")
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

    // 水仙花数 三位数 [100 999]
    // 每位数上的数字的立方和，刚好等于该数字本身，那么就叫水仙花数
    for i := 100; i <= 999; i++ {
        x := i / 100     // 百位
        y := i / 10 % 10 // 十位
        z := i % 10      //个位

        if math.Pow(float64(x), 3) + 
            math.Pow(float64(y), 3) + 
            math.Pow(float64(z), 3) == float64(i) {
            fmt.Printf("%d\t", i)
        }
    }
    fmt.Print("\n")

    fmt.Println("---------------------")

    /**
     * 百位：1-9
     * 十位：0-9
     * 个位：0-9
     */
    for i := 1; i < 10; i++ {
        for j := 0; j < 10; j++ {
            for k := 0; k < 10; k++ {
                n := i * 100 + j * 10 + k
                if i*i*i + j*j*j + k*k*k == n {
                    fmt.Printf("%d\t", n)
                }
            }
        }
    }
    fmt.Print("\n")

    fmt.Println("---------------------")

    //素数
    //只能被1和自身整除的数
    for i := 2; i <= 100; i++ {
        flag := true
        for j := 2; j <= int(math.Sqrt(float64(i))) ; j++ {
            if i % j == 0 {
                flag = false
                break
            }
        }

        if flag {
            fmt.Printf("%d\t", i)
        }
    }
    fmt.Print("\n")
}
