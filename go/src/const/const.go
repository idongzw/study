/*
* @File Name: const.go
* @Author: idongzw
* @Date:   2020-02-15 12:57:06
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-15 17:30:12
*/

package main

import "fmt"

func main() {
    /*
    常量使用关键字 const 定义，用于存储不会改变的数据。

    存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

    常量的定义格式：const identifier [type] = value
    */
    const Pi = 3.14159
    const b string = "abc"
    const c = "abc"

    fmt.Println(Pi, b, c)

   /*
   常量的值必须是能够在编译时就能够确定的；
   你可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得
    */
    const d = 2 / 3
    fmt.Println(d)

    //var sss = "asd"
    const (
        // sss 为 const 类型就可以了
        //a1 = len(sss) //const initializer len(sss) is not a constant
        a2 = 1
        a3
    )
    fmt.Println(a2, a3)

    const (
        b1 = "dzw"
        b2 = len(b1) // b1是const类型, len 为内置函数
        b3
    )
    fmt.Println(b1, b2, b3)
    // const initializer getNumber() is not a constant
    //const c2 = getNumber() // 引发构建错误: getNumber() used as value 
    //因为在编译期间自定义函数均属于未知，因此无法用于常量的赋值，但内置函数可以使用，如：len()
    
    // 在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
    const (
        c1, c2 = 1, "22"
        d1, d2
    )
    fmt.Println(c1, c2)
    fmt.Println(d1, d2)

    const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009 
    const Log2E = 1/Ln2 // this is a precise reciprocal
    const Billion = 1e9 // float constant
    const hardEight = (1 << 100) >> 97

    fmt.Println(Ln2, Log2E, Billion, hardEight)

    const (
        Unknown = 0
        Female = 1
        Male = 2
    )

    // iota 是常量的计数器，从 0 开始，组内每定义 1 个常量自动递增 1
    const (
        e = "e"
        f = iota
        g
    )

    //在每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0
    //简单地讲，每遇到一次 const 关键字，iota 就重置为 0
    //第一个常量不能省略表达式
    const (
        h = iota + 50
        i
        j
    )

    fmt.Println(Unknown, Female, Male)
    fmt.Println(e, f, g)
    fmt.Println(h, i, j)
}

func getNumber() int {
    return 3
}