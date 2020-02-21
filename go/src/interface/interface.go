/*
* @File Name: interface.go
* @Author: idongzw
* @Date:   2020-02-16 18:33:30
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-20 17:59:12
*/
package main

import "fmt"

/*
当某个类型为这个接口中的所有方法提供了方法的实现，它被称为实现接口
Go语言中，接口和类型的实现关系，是非侵入式
//其他语言
class Mouse implements USB {}

1. 当需要接口类型的对象时，可以使用任意实现类对象代替
2. 接口对象不能访问实现类中的属性

接口的用法：
    1. 一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的任意实现类型对象作为参数
    2. 定义一个类型为接口类型，实际上可以赋值为任意实现类的对象
 */

// 空接口
// 不包含任何的方法，所以所有的类型都实现了空接口，因此空接口可以存储任意类型的数值
type Empty interface {

}

// 接口嵌套
type USB interface {
    Name() string
    Connect
}

type Connect interface {
    Connected()
}

type PhoneConnecter struct {
    name string
}

func (pc PhoneConnecter) Name() string {
    return pc.name
}

func (pc PhoneConnecter) Connected() {
    fmt.Println("Connected:", pc.name)
}

func DisConnected(usb USB) {
    //fmt.Println("DisConnected:", usb.Name())

    if pc, ok := usb.(PhoneConnecter); ok {
        fmt.Println("DisConnected:", pc.name)
        return
    }
    fmt.Println("Unknown device")
}

// 空接口
func DisConnected2(usb interface{}) {
    switch v := usb.(type) {
    case PhoneConnecter:
        fmt.Println("DisConnected:", v.name)
    default:
        fmt.Println("Unknown device")
    }
}

type TVConnecter struct {
    name string
}

func (tv TVConnecter) Connected() {
    fmt.Println("Connected:", tv.name)
}

func main() {
    pc := PhoneConnecter{"PhoneConnecter"}

    fmt.Println(pc.Name())
    pc.Connected()
    DisConnected(pc)
    DisConnected2(pc)

    var usb USB
    usb = USB(pc)

    fmt.Println(usb.Name())
    usb.Connected()
    DisConnected(usb)
    DisConnected2(usb)

    var c Connect
    c = Connect(pc)
    c.Connected()
    //c.Name() // type Connect has no field or method Name
    //DisConnected(c) // Connect does not implement USB (missing Name method)
    DisConnected2(c)

    tv := TVConnecter{"TVConnecter"}
    var c2 Connect 
    c2 = Connect(tv)
    c2.Connected()

    // TVConnecter 没有实现 interface USB
    //var usb2 USB
    //usb2 = USB(tv) // cannot convert tv (type TVConnecter) to type USB:
    //TVConnecter does not implement USB (missing Name method)
    //usb2.Connected()
    
    /*
    将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针
    既无法修改复制品的状态，也无法获取指针
     */
    pc.name = "sadsadsad"
    usb.Connected()
    c.Connected()
    pc.Connected()

    // 只有当接口存储的类型和对象都为 nil 时，接口才等于 nil
    var i1 interface{}
    fmt.Println(i1 == nil) // true

    var p *int = nil
    i1 = p
    fmt.Println(i1 == nil) // false

    // 接口调用不会做receiver的自动转换
    // 接口同样支持匿名字段方法
    // 接口也可实现类似OOP的多态
    // 空接口可以作为任何类型数据的容器
    
    var e1 Empty = "sadsa"
    var e2 Empty = PhoneConnecter{"iPhone"}
    var e3 Empty = 12

    fmt.Println(e1)
    fmt.Println(e2)
    fmt.Println(e3)

    // map value 任意类型
    m := make(map[string]interface{})
    m["sad"] = 1
    m["sa"] = PhoneConnecter{"iPhone"}
    m["wq"] = false

    fmt.Println(m)

    // slice 存储任意类型
    s := make([]interface{}, 0, 10)
    s = append(s, "dzw", 1, false, PhoneConnecter{"iPhone"})

    fmt.Println(s)
    Display(s)
}

func Display(s []interface{}) {
    num := len(s)

    for i := 0; i < num; i++ {
        fmt.Println(s[i])
    }
}