/*
* @File Name: interface.go
* @Author: idongzw
* @Date:   2020-02-16 18:33:30
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-17 12:09:24
*/
package main

import "fmt"

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

// 空借口
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
}