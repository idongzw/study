/*
* @File Name: basic_structure.go
* @Author: idongzw
* @Date:   2020-02-15 15:34:46
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-18 20:07:04
*/

// go 程序的一般结构

// 当前程序包名
package main

// 导入其他包
//import "fmt"
//import std "fmt" //alias
//import . "fmt"

// 导入多个包

import (
    "fmt"
    "math"
    "strconv"
    "time"
)

// 常量定义
//const PI = 3.14
// 定义多个常量

const (
    PI = 3.14
    Name = "dzw"
)


// 全局变量的声明与赋值
//var name = "dzw"
// 定义多个全局变量

var (
    name = "dzw"
    age = 26
)


// 一般类型声明
//type newType int
// 声明多个一般类型

type (
    newType1 int
    newType2 float32
)

// 结构的声明
type gostruct struct {}

// 接口的声明
type gointerface interface {}

type 文本 string
var chinese 文本 = "中文"

// main 函数作为程序入口
func main() {
    fmt.Println("Hello,你好")
    //std.Println("Hello,你好") //alias
    //Println("Hello,你好")
    
    fmt.Println(chinese)

    fmt.Println(math.MaxInt8)

    a := 65
    // string 表示将数据转换为文本格式，数字65表示文本A
    b := string(a)
    fmt.Println(a, b)

    b = strconv.Itoa(a)
    fmt.Println(b)
    a, _ = strconv.Atoi(b)
    fmt.Println(a)

    fmt.Println("++++++++++++++++++++++++++++++")
    //需要注意的地方(坑)
    /*
    1. slice append 超过容量时会自动扩展新建一个slice，先拷贝原来的值再append
    */
    {
        s := make([]int, 0)
        fmt.Println(s)
        s = Append(s)
        fmt.Println(s)
    }

    /*
    2. time
     */
    {
        t := time.Now()
        fmt.Println(t)
        fmt.Println(t.Format(time.ANSIC))
        fmt.Println(time.ANSIC)
        fmt.Println(t.Format("Mon Jan _2 15:04:05 2006")) //起始时间不能修改
    }

    /*
    3. for range
     */
    {
        s := []string{"a", "b", "c"}

        // 输出 c c c
        // 不按参数传递的值，都是引用
        for _, v := range s {
            go func() {
                fmt.Println(v)
            }()
        }

        // 输出 a b c
        // 按参数传递
        for _, v := range s {
            go func(v string) {
                fmt.Println(v)
            }(v)
        }

        select {}
    }
}

// 需设置返回值
func Append(s []int) []int {
    s = append(s, 3)
    return s
}

// 可见性规则
/*
go语言中，使用 大小写 来决定该 常量、变量、类型、接口、结构或函数是否可以被外部包所调用
根据预定，函数名首字母小写为private，首字母大写为public
访问权限是 package与package 之间
 */

//alias
//byte 是 uint8 的别名
//rune 是 int32 的别名

// int
// 和操作系统位数有关 32/64

// float32/float64
//  -长度：4/8 字节
//  -小数位：精确到 7/15 小数位

// uintptr
// 32/64位整数类型 ，保存指针


// 类型零值
/* 零值并不等于空值，而是当变量被声明为某种类型后的默认值，
通常情况下值类型的默认值为 0 ，bool 为 false， string 为空字符串
*/