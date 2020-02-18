/*
* @File Name: array.go
* @Author: idongzw
* @Date:   2020-02-15 20:08:04
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-15 21:00:45
*/
package main

import (
    "fmt"
)

func main() {

    // 数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
    {
        var a [2]int
        var b [1]int

        //b = a // cannot use a (type [2]int) as type [1]int in assignment
        fmt.Println(a, b)
    }

    {
        var a [2]int
        var b [2]int

        b = a
        fmt.Println(a, b)
    }

    {
        a := [2]int{1}
        fmt.Println(a)
    }

    {
        a := [3]int{1, 2, 3}
        fmt.Println(a)
    }

    {
        a := [20]int{19:1, 18:3} //按数组下标赋值 
        fmt.Println(a)
    }

    {
        a := [...]int{1, 2, 3, 4, 5}
        fmt.Println(a)
    }

    {
        a := [...]int{19:1}
        fmt.Println(a)
    }

    {
        a := [...]int{99:1}
        var p *[100]int = &a // 数组指针

        fmt.Println(p)
    }

    {
        x, y := 1, 2
        a := [...]*int{&x, &y} //指针数组

        fmt.Println(a)
    }


    // 数组之间可以用 == 或 != 进行比较，不可用 < 或 >
    {
        a := [2]int{1, 2}
        b := [2]int{1, 2}

        fmt.Println(a == b) // true
        fmt.Println(a != b) // false

        b[1] = 1
        fmt.Println(a == b) // false

        //fmt.Println(a > b) // invalid operation: a > b (operator > not defined on array)
        
        c := [1]int{1}
        fmt.Println(c)
        //fmt.Println(a == c) // invalid operation: a == c (mismatched types [2]int and [1]int)
    }

    // 可以使用new来创建数组，此方法返回一个指向数组的指针
    {
        a := [10]int{}
        a[1] = 2
        fmt.Println(a)

        p := new([10]int)
        p[1] = 2
        fmt.Println(p)
    }

    // go 支持多维数组
    {
        a := [2][3]int {
            {1, 2},
            {4, 5}}

        fmt.Println(a)

        b := [2][3]int {
            {1:2},
            {2:5}}

        fmt.Println(b)

        c := [...][3]int {
            {},
            {}}
        fmt.Println(c)
    }

    // 冒泡排序 (BubbleSort)
    {
        a := [...]int {5, 2, 3, 1, 7, 6, 9}
        fmt.Println(a)

        num := len(a)
        for i := 0; i < num; i++ {
            for j := i + 1; j < num; j++ {   
                if a[i] > a[j] {
                    tmp := a[i]
                    a[i] = a[j]
                    a[j] = tmp
                }
            }
        }

        fmt.Println(a)
    }
}