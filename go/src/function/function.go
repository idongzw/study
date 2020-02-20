/*
* @File Name: function.go
* @Author: idongzw
* @Date:   2020-02-16 12:56:17
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-20 13:48:23
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

        // 按参数传递
        defer func(i int) {
            fmt.Println("param i =", i) // i 值
        }(i)
    }

    for _, f := range fs {
        f()
    }

    GetSum()
    GetSum(1,2,3,4,5)

    s1 := []int{1,2,3,4,5,6,7,8,9,10}
    GetSum(s1...)

    // 参数值传递
    arr1 := [...]int{1,2,3,4}
    fmt.Println("array before call test_param1", arr1) // [1 2 3 4]
    test_param1(arr1)
    fmt.Println("array after call test_param1", arr1) // [1 2 3 4]

    // 参数引用传递
    s2 := []int{1,2,3,4}
    m1 := map[int]string{1:"dzw", 2:"dada"}
    fmt.Println("slice,map before call test_param2", s2, m1) // [1 2 3 4] map[2:dada 1:dzw]
    test_param2(s2, m1)
    fmt.Println("slice,map after call test_param2", s2, m1) // [100 2 3 4] map[1:asd 2:dada]

    // 递归函数
    fmt.Println(getSum(100)) // 5050

    fmt.Println("--------------------------")
    a1 := 2
    fmt.Println("a1 =", a1)
    a1++
    fmt.Println("main a1 =", a1)
    fmt.Println("func_defer1() =", func_defer1())
    defer func_defer2(a1)
    fmt.Println("--------------------------")

    fmt.Println(oper(1, 2, add))
    fmt.Println(oper(1, 2, sub))
    fmt.Println(oper(2, 3, func(a, b int) int {
        return a * b
    }))

    fmt.Println("--------------------")
    f1 := increment()
    fmt.Printf("%T %v\n", f1, f1)
    v1 := f1()
    fmt.Println(v1) // 1
    v2 := f1()
    fmt.Println(v2) // 2
    fmt.Println("--------------------")
}

/*
defer 
1. 当外围函数中的语句正常执行完毕后，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行
2. 当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回
3. 当外围函数中的代码引发运行panic时，只有其中所有的延迟函数都执行完毕后，该运行时panic才会真正被扩展至调用函数
 */

func func_defer1() int {
    fmt.Println("call func_defer1...")
    defer func() {
        fmt.Println("defer func_defer1 1")
    }()

    defer func() {
        fmt.Println("defer func_defer1 2")
    }()

    return 0
}

func func_defer2(a int) {
    fmt.Println("call func_defer2 a =", a)
}

func A() int {
    return 3
}

//多返回值
func B() (a, b, c int) {
    a, b, c = 1, 2, 3

    return
}

/*
可变参数：
    概念：一个函数的参数的类型确定，但个数不确定，就可以使用可变参数
    语法：参数名 ...参数类型

    对于函数，可变参数相当于一个切片
    调用函数的时候，可以传入 0 个或多个参数

    注意事项：
        A: 如果一个函数的参数是可变参数，同时还有其他的参数，可变参数要放在参数列表的最后
        B: 一个函数的参数列表中最多只能有一个可变参数
 */

// 可变参数相当于一个slice
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

func GetSum(nums ...int) {
    sum := 0
    count := len(nums)
    for i := 0; i < count; i++ {
        sum += nums[i]
    }
    fmt.Println("sum:", sum)
}

/*
参数传递：
    1. 值传递：传递的是数据的副本，修改数据，对于原始的数据没有影响
        值类型的数据，默认都是值传递：基本类型, array, struct

    2. 引用传递：传递的是数据的地址，导致多个变量指向同一块内存
        引用类型的数据，默认都是引用传递：slice, map, chan
 */
// 值传递
// param array
func test_param1(arr [4]int) {
    fmt.Println("func test_param1 before mod", arr) // [1 2 3 4]
    arr[0] = 100
    fmt.Println("func test_param1 after mod", arr) // [100 2 3 4]
}

// 引用传递
// param slice map
func test_param2(s []int, m map[int]string) {
    fmt.Println("func test_param1 before mod", s, m) // [1 2 3 4] map[1:dzw 2:dada]
    s[0] = 100
    m[1] = "asd"
    fmt.Println("func test_param1 after mod", s, m) // [100 2 3 4] map[1:asd 2:dada]
}

// 闭包
// 函数作为另一个函数的返回值
func closure(x int) func(int) int {
    fmt.Printf("%p\n", &x)
    return func(y int) int {
        fmt.Printf("- %p\n", &x)
        return x + y
    }
}

/*
一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量（外层函数中的参数，或者
外层函数中直接定义的变量），并且该外层函数的返回值就是这个内层函数

这个内层函数和外层函数的局部变量，统称为闭包结构。

局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁。
但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。
 */
func increment() func() int { // 外层函数
    // 局部变量
    i := 0
    // 定义一个匿名函数
    f := func() int { //内层函数
        i++
        return i
    }

    // 返回匿名函数
    return f
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

go 没有异常机制，但有 panic/recover 模式来处理错误：
panic函数用于引发恐慌，导致程序终端执行
recover函数用于恢复程序的执行，recover()语法上要求必须在defer中执行
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

//递归函数
func getSum(n int) int {
    if n == 1 {
        return 1
    }
    return getSum(n - 1) + n
}

/*
高阶函数：
    讲一个函数作为另一个函数的参数

    func1(), func2()
    将func1函数作为func2这个函数的参数
        func2函数：就叫高阶函数
            接收了一个函数作为参数的函数，高阶函数

        func1函数：回调函数
            作为另一个函数的参数的函数，回调函数
 */

// +
func add(a, b int) int {
    return a + b
} 

func sub(a, b int) int {
    return a - b
}

// oper
func oper(a, b int, f func(int, int) int) int {
    return f(a, b)
}