/*
* @File Name: function.go
* @Author: idongzw
* @Date:   2020-02-16 12:56:17
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-16 15:18:04
*/
package main

import (
    "fmt"
)

/*
go 函数不支持嵌套、重载和默认参数
支持以下特性：
无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包

定义函数使用关键字func，且左大括号不能另起一行
函数也可以作为一种类型使用
 */
func main() {
    fmt.Println(A())
    fmt.Println(B())

    C(1, 2, 3)
    C(1, 2, 3, 4, 5)

    D("dzw", 1, 2)
    D("dzw", 1, 2, 3)

    a, b, c := 1, 2, 3
    E(a, b, c);
    fmt.Println(a, b, c)

    f := A
    fmt.Println(f())

    // 匿名函数
    f2 := func() {
        fmt.Println("Anonymous Function")
    }

    f2()

    fmt.Println((closure(2))(3))
    f3 := closure(3)
    fmt.Println(f3(5))
    fmt.Println(f3(6))

    //defer
    defer fmt.Println("defer1")
    defer fmt.Println("defer2")

    for i := 0; i < 3; i++ {
        defer fmt.Println(i) // i 值
    }

    for i := 0; i < 3; i++ {
        defer func() {
            fmt.Println(i) // i 引用
        }()
    }

    func1()
    func2()
    func3()

    fs := [4]func(){}

    for i := 0; i < 4; i++ {
        defer fmt.Println("defer i = ", i)
        defer func() {
            fmt.Println("defer_closure i = ", i) // i 引用
        }()

        fs[i] = func() {
            fmt.Println("closure i = ", i) // i 引用
        }
    }

    for _, f := range fs {
        f()
    }
}

func A() int {
    return 3
}

func B() (a, b, c int) {
    a, b, c = 1, 2, 3

    return
}

// 可变参数
func C(a ...int) {
    fmt.Println(a)
}

// 可变参数只能放在最后一个参数
func D(a string, b ...int) {
    fmt.Println(a, b)
}

func E(a ...int) {
    a[0] = 2
    fmt.Println(a)
}

// 闭包
func closure(x int) func(int) int {
    fmt.Printf("%p\n", &x)
    return func(y int) int {
        fmt.Printf("- %p\n", &x)
        return x + y
    }
}

// defer
/*
defer 的执行方式类似其他语言中的析构函数，在函数体执行结束后
按照调用顺序的相反顺序逐个执行

即使函数发生严重错误也会执行
支持匿名函数的调用
常用于资源清理、文件关闭、解锁以及记录时间等操作
通过与匿名函数配合可在return之后修改函数计算结果
如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时即已经获得了拷贝，否则则是引用某个变量的地址

go 没有异常机制，但有 panic/recover 模式来处理错误
panic可以在任何地方引发，但recover只有在defer调用的函数中有效
 */

func func1() {
    fmt.Println("func1")
}

func func2() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
            fmt.Println("Recover in B")
        }
    }()
    panic("Panic in func2")
}

func func3() {
    fmt.Println("func3")
}