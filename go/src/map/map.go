/*
* @File Name: map.go
* @Author: idongzw
* @Date:   2020-02-16 12:04:36
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-16 12:55:09
*/
package main

import (
    "fmt"
    "sort"
)

func main() {
    // key-value 形式存储数据
    // key 必须是支持 == 或者 != 比较运算的类型，不可以是函数、map或slice
    // make(map[keyType]valueType, cap), cap 表示容量，可省略
    // 超出容量时会自动扩容，但尽量提供一个合理的初始值
    // len()获取元素个数
    // 键值对不存在时自动添加，使用 delete() 删除某键值对
    // 使用 for range 对 map 和 slice 进行迭代操作
    {
        var m map[int]string
        m = map[int]string{}

        fmt.Println(m)

        m = make(map[int]string)
        fmt.Println(m)

        m1 := make(map[int]string)
        fmt.Println(m1)

        m2 := map[int]string{}
        fmt.Println(m2)

        m[1] = "Ok"
        fmt.Println(m) // map[1:Ok]
        value1 := m[1]
        fmt.Println(value1) // Ok
        value2 := m[2]
        fmt.Println(value2) //
        m[2] = "Sa"
        fmt.Println(m) // map[1:Ok 2:Sa]

        delete(m, 1)
        fmt.Println(m) // map[2:Sa]
    }

    // 每一级都要初始化
    {
        m := make(map[int]map[int]string)
        fmt.Println(m)

        m[1] = make(map[int]string)
        m[1][1] = "Ok"
        m[1][2] = "SS"
        m[2] = make(map[int]string)
        m[2][1] = "Ds"
        fmt.Println(m) // map[1:map[1:Ok 2:SS] 2:map[1:Ds]]

        v, ok := m[3][1]
        if !ok {
            fmt.Println("make m[3]")
            m[3] = make(map[int]string)
        }
        m[3][1] = "AA"
        v, ok = m[3][1]
        fmt.Println(v, ok)
    }

    // 迭代
    {
        sm := make([]map[int]string, 5)
        for _, v := range sm { // v 是拷贝
            v = make(map[int]string)
            v[1] = "Ok"
            fmt.Println(v) // map[1:Ok]
        }
        fmt.Println(sm) // [map[] map[] map[] map[] map[]]
    }

    {
        sm := make([]map[int]string, 5)
        for i := range sm { 
            sm[i] = make(map[int]string)
            sm[i][1] = "Ok"
            fmt.Println(sm[i]) // map[1:Ok]
        }
        fmt.Println(sm) // [map[1:Ok] map[1:Ok] map[1:Ok] map[1:Ok] map[1:Ok]]
    }

    // map 间接排序
    {
        m := map[int]string{1:"dzw", 2:"zaq", 3:"sad", 4:"sdf"}
        s := make([]int, len(m))

        fmt.Println(m)
        i := 0
        for k := range m {
            s[i] = k
            i++
        }
        sort.Ints(s)

        fmt.Println(s)
    }

    // key value 交换
    {
        m1 := map[int]string{1:"dzw", 2:"zaq", 3:"sad", 4:"sdf"}
        fmt.Println(m1)

        m2 := make(map[string]int, len(m1))
        //m2 := map[string]int{}
        for k, v := range m1 {
            m2[v] = k
        }
        fmt.Println(m2)
    }
}