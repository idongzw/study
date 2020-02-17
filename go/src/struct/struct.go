/*
* @File Name: struct.go
* @Author: idongzw
* @Date:   2020-02-16 15:01:56
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-16 17:11:39
*/
package main

import "fmt"

/*
type <Name> struct{}
支持指向自身的指针类型成员
支持匿名结构，可用作成员或定义成员变量
匿名结构也可以用于 map 的值
可以使用字面值对结构进行初始化
允许直接通过指针来读写结构成员
相同类型的成员可进行直接拷贝赋值
支持 == != 不支持 > <
支持匿名字段，本质上是定义了以某个类型名为名称的字段
嵌入结构作为匿名字段看起来像继承，但不是继承
可以使用匿名字段指针
 */

type person struct {
    Name string
    Age int
}

// 匿名结构
type person2 struct {
    Name string
    Age int
    Contact struct {
        Phone, City string
    }
}

// 匿名字段
type person3 struct {
    string 
    int
}

type human struct {
    Sex int
}

type teacher struct {
    human
    Name string
    Age int
}

type student struct {
    human
    Sex int
    Name string
    Age int
}

func main() {
    p := person{}
    fmt.Println(p) // { 0}
    p.Name = "dzw"
    p.Age = 26
    fmt.Println(p) // {dzw 26}

    p2 := person{
        Name: "Joe",
        Age: 18,
    }
    fmt.Println(p2) // {Joe 18}

    A(&p2)
    fmt.Println(p2)

    p3 := &person{ // 初始化尽量使用这种方式
        Name: "asd",
        Age: 19,
    }
    p3.Name = "assas"

    fmt.Println(p3)
    A(p3)
    fmt.Println(p3)

    // 匿名结构
    p4 := struct {
        Name string
        Age int
    }{
        Name: "wawda",
        Age: 19,
    }

    fmt.Println(p4)

    p5 := &struct {
        Name string
        Age int
    }{
        Name: "wawda",
        Age: 19,
    }

    fmt.Println(p5)

    p6 := person2{
        Name: "qaz",
        Age: 90,
    }
    p6.Contact.Phone = "1233333333"
    p6.Contact.City = "BJ"

    fmt.Println(p6)

    p7 := person3 {
        "Joe",
        18,
    }
    fmt.Println(p7)

    p8 := person{
        Name: "Joe",
        Age: 18,
    }

    p9 := person{
        Name: "Joe",
        Age: 18,
    }

    fmt.Println(p8 == p9)

    fmt.Println("-----------------------------")

    p10 := teacher {
        Name: "asd",
        Age: 27,
        human: human{Sex: 2},
    }

    fmt.Println(p10) // {{2} asd 27}
    p10.human.Sex = 100
    fmt.Println(p10) // {{100} asd 27}
    p10.Sex = 10
    fmt.Println(p10) // {{10} asd 27}

    p11 := student {
        Name: "sad",
        Age: 12,
        Sex: 23,
    }
    fmt.Println(p11) // {{0} 23 sad 12}

    p12 := student {
        Name: "sada",
        Age: 12,
        Sex: 233,
        human: human{Sex: 1},
    }
    fmt.Println(p12) // {{1} 233 sada 12}
    p12.Sex = 4
    fmt.Println(p12) // {{1} 4 sada 12}
    p12.human.Sex = 23
    fmt.Println(p12) // {{23} 4 sada 12}
}

func A(p *person) {
    p.Age = 12
    fmt.Println(p)
}