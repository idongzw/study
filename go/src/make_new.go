/*
* @File Name: make_new.go
* @Author: idongzw
* @Date:   2020-02-22 11:49:55
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-22 12:12:47
*/
package main

import "fmt"

//make new
/*
make 用于内建类型(map、slice和channel)的内存分配

内建函数 new 本质上说跟其他语言中的同名函数功能一样；
new(T)分配了零值填充的T类型的内存空间，并且返回其地址，
即一个 *T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值
new返回的是指针

内建函数make(T, args) 与 new(T) 有着不同的功能，make只能创建 slice，map，channel，
并且返回一个有初始值(非零)的T类型，而不是 *T。本质来讲，导致这三个类型有所不同的原因
是指向数据结构的引用在使用之前必须被初始化。例如，一个slice，是一个包含指向数据（内部array）的指针、
长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化
了内部的数据结构，填充适当的值。
make返回初始化后的（非零）值
 */

func main() {
    {
        // slice
        s1 := make([]int, 0, 3)
        fmt.Printf("s1 type %T, s1=%v\n", s1, s1) // s1 type []int, s1=[]
        if s1 == nil {
            fmt.Println("s1 == nil")
        } else {
            fmt.Println("s1 != nil") // s1 != nil
        }

        s2 := new([]int)
        fmt.Printf("s2 type %T, s2=%v\n", s2, s2) // s2 type *[]int, s2=&[]
        fmt.Printf("*s2=%v\n", *s2) // *s2=[]
        if *s2 == nil {
            fmt.Println("*s2 == nil") // *s2 == nil
        } else {
            fmt.Println("*s2 != nil")
        }

        // int
        i := new(int)
        fmt.Println("new *i =", *i) // new *i = 0

        // bool
        b := new(bool)
        fmt.Println("new *b =", *b) // new *b = false

        // map
        m1 := make(map[int]string, 0)
        fmt.Printf("m1 type %T, m1=%v\n", m1, m1) // m1 type map[int]string, m1=map[]
        if m1 == nil {
            fmt.Println("m1 == nil")
        } else {
            fmt.Println("m1 != nil") // m1 != nil
        }

        m2 := new(map[int]string)
        fmt.Printf("m2 type %T, m2=%v\n", m2, m2) // m2 type *map[int]string, m2=&map[]
        fmt.Printf("*m2=%v\n", *m2) // *m2=map[]
        if *m2 == nil {
            fmt.Println("*m2 == nil") // *m2 == nil
        } else {
            fmt.Println("*m2 != nil")
        }

        // chan
        c1 := make(chan int)
        fmt.Printf("c1 type %T, c1=%v\n", c1, c1) // c1 type chan int, c1=0xc1e100
        if c1 == nil {
            fmt.Println("c1 == nil")
        } else {
            fmt.Println("c1 != nil") // c1 != nil
        }

        c2 := new(chan int)
        fmt.Printf("c2 type %T, c2=%v\n", c2, c2) // m2 type *map[int]string, m2=&map[]
        fmt.Printf("*c2=%v\n", *c2) // *c2=<nil>
        if *c2 == nil {
            fmt.Println("*c2 == nil") // *c2 == nil
        } else {
            fmt.Println("*c2 != nil")
        }
    }
}