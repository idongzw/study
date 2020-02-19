/*
* @File Name: map.go
* @Author: idongzw
* @Date:   2020-02-16 12:04:36
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-19 21:22:23
*/
package main

import (
    "fmt"
    "sort"
)

/*
    1. 创建map:
        var map1 map[keyType]valueType
        nil map 无法直接使用

        var map2 = make(map[keyType]valueType)

        var map3 = map[keyType]valueType{key:value,key:value,key:value...}
    2. 添加/修改
        map[key] = value
            如果key不存在，就是添加数据
            如果key存在，就是修改数据
    3. 获取
        map[key] -----> value
        value, ok := map[key]
            根据key获取对应的value：
            如果key存在，value就是对应的数据，ok为 true
            如果key不存在，value就是值类型的默认值，ok为false
    4. 删除数据
        delete(map, key)
        如果key存在，就可以直接删除
        如果key不存在，删除失败
    5. 长度
        len()
*/

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

    {
        var m1 map[int]string //没有初始化 nil
        //m1[1] = "dzw" // panic: runtime error: assignment to entry in nil map
        var m2 = make(map[int]string)
        var m3 = map[string]int{"Go": 1, "dzw": 2}

        fmt.Println(m1)
        fmt.Println(m2)
        fmt.Println(m3)

        fmt.Println(m1 == nil) // true
        fmt.Println(m2 == nil) // false
        fmt.Println(m3 == nil) // false

        if m1 == nil {
            m1 = make(map[int]string) // 初始化 
        }
        m1[1] = "dzw"

        fmt.Println(m1)
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

        v, ok := m[3]
        fmt.Println(v, ok)
        if !ok {
            fmt.Println("make m[3]")
            m[3] = make(map[int]string)
        }
        m[3][1] = "AA"
        v, ok = m[3]
        fmt.Println(v, ok)

        v2, ok2 := m[3][2]
        if !ok2 {
            fmt.Println("m[3][2] not exist")
        }
        fmt.Println(v2, ok2)
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

        m := make(map[int]string)
        m[1] = "sad"
        m[2] = "sadsa"
        m[3] = "wadac"
        m[4] = "sadwa"

        for k, v := range m {
            fmt.Println(k, v)
        }
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

        for _, key := range s {
            fmt.Println(key, m[key])
        }
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

    {
        // slice string 排序
        s := []string{"Apple", "Windows", "Orange", "abc", "刹车", "acd", "acc"}
        fmt.Println(s)
        sort.Strings(s)
        fmt.Println(s)
    }

    {
        m1 := map[string]string{"name":"dzw", "age":"26", "sex":"Male"}
        m2 := map[string]string{"name":"qaz", "age":"24", "sex":"Female"}
        m3 := map[string]string{"name":"asd", "age":"23", "sex":"Male"}

        ms := make([]map[string]string, 0, 3)

        fmt.Println(m1)
        fmt.Println(m2)
        fmt.Println(m3)

        ms = append(ms, m1)
        ms = append(ms, m2)
        ms = append(ms, m3)

        fmt.Println(ms)
    }

    // map 引用类型
    {
        m1 := map[int]string{1:"dzw", 2:"qqq"}
        m2 := m1

        fmt.Println(m1) // map[1:dzw 2:qqq]
        fmt.Println(m2) // map[1:dzw 2:qqq]

        m2[1] = "aaa"
        fmt.Println(m1) // map[1:aaa 2:qqq]
        fmt.Println(m2) // map[1:aaa 2:qqq]
    }
}