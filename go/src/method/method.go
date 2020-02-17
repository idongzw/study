/*
* @File Name: method.go
* @Author: idongzw
* @Date:   2020-02-16 17:38:36
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-17 12:07:43
*/
package main

import "fmt"

type method struct {
    name string // 首字母大写其他package可以访问
}

type TZ int

func main() {
    m := method {
        name: "dzw",
    }

    // 方法调用会做receiver的自动转换
    m.Print1()
    fmt.Println(m)
    m.Print2()
    fmt.Println(m)

    var tz TZ = 1
    tz.Print()
    (*TZ).Print(&tz)

    tz.Increase(100)
    fmt.Println(tz)

    method.Print1(m)
}

// 绑定方法到 method struct
func (m method) Print1() {
    m.name = "asd"
    fmt.Println("func Print1() for method struct")
}

func (m *method) Print2() {
    m.name = "asd"
    fmt.Println("func Print2() for method struct")
}

// 绑定方法到 type TZ int
func (tz *TZ) Print() {
    fmt.Println("func Print() for TZ type")
}

func (tz *TZ) Increase(num int) {
    *tz += TZ(num)
}