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
空接口作为函数的参数:
使用空接口实现可以接收任意类型的函数参数
空接口作为map的值:
使用空接口实现可以保存任意值的字典
*/

// 空接口
// 不包含任何的方法，所以所有的类型都实现了空接口，因此空接口可以存储任意类型的数值

// 类型断言
/*
 接口值:
一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值

判断空接口中的值，可以使用类型断言：
	x.(T)
x : 表示类型为 interface{}的变量
T : 表示断言 x 可能是的类型
该语法返回两个参数，第一个参数是 x 转化为 T 类型后的变量，第二个值是一个bool值，true表示断言成功，false失败
*/
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

/*
 值接收者和指针接收者实现接口的区别
*/
// Mover ...
type Mover interface {
	move()
}

type dog struct{}

// 1. 值接收者实现接口
func (d dog) move() {
	fmt.Println("dog can move...")
}

type cat struct{}

// 2. 指针接收者实现接口
func (c *cat) move() {
	fmt.Println("cat can move...")
}

// 类型与接口的关系

// 1. 一个类型实现多个接口
// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现
// Sayer1 接口
type Sayer1 interface {
	say()
}

// Mover1 接口
type Mover1 interface {
	move()
}

type dog1 struct {
	name string
}

// 实现 Sayer1 接口
func (d dog1) say() {
	fmt.Printf("%s can say, wangwangwang~\n", d.name)
}

// 实现 Mover1 接口
func (d dog1) move() {
	fmt.Printf("%s can move\n", d.name)
}

/*
 一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现
*/
// WashingMachine ...
// 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("dry")
}

// haier 洗衣机
type haier struct {
	dryer // 嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("haier wash")
}

// 接口嵌套
// 接口与接口间可以通过嵌套创造出新的接口
type animal interface {
	Sayer1
	Mover1
}

type cat1 struct{}

func (c cat1) say() {
	fmt.Println("miaomiaomiao~")
}

func (c cat1) move() {
	fmt.Println("cat can move...")
}

func main() {
	{
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

		i1 = 1
		fmt.Println("i1 = ", i1)

		i1 = (*int)(nil)
		fmt.Println("i1 = ", i1 == nil)

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

	// 值接收者实现接口测试
	{
		var m Mover
		d1 := dog{}  // dog 类型
		d2 := &dog{} // *dog 类型

		m = d1
		m.move()
		// 使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
		// 因为Go语言中有对指针类型变量求值的语法糖，dog指针 d2 内部会自动求值 *d2
		m = d2
		m.move()
	}

	// 指针接收者实现接口测试
	{
		var m Mover
		//c1 := cat{}  // cat 类型
		c2 := &cat{} // *cat 类型

		/*
		   此时实现Mover接口的是*cat类型，所以不能给m传入cat类型的c1，此时m只能存储*cat类型的值
		*/
		//m = c1 // cannot use c1 (type cat) as type Mover in assignment:
		//cat does not implement Mover (move method has pointer receiver)
		m = c2
		m.move()
	}

	// 一个类型实现多个接口
	{
		var s Sayer1
		var m Mover1

		d := dog1{"asd"}

		s = d
		m = d

		s.say()
		m.move()

		d.say()
		d.move()
	}

	// 多个类型实现同一接口
	{
		var m Mover

		d := dog{}
		c := cat{}

		m = d
		m.move()

		m = &c
		c.move()
	}

	{
		var w WashingMachine

		w = haier{}
		w.dry()
		w.wash()
	}

	{
		var a animal

		a = cat1{}
		a.move()
		a.say()

		var s Sayer1 = a
		s.say()

		var m Mover1 = a
		m.move()
	}

	// 空接口类型的变量可以存储任意类型的变量
	{
		var empty interface{}

		s := "string~~"
		empty = s
		fmt.Printf("type:%T vaule:%v\n", empty, empty)

		i := 12
		empty = i
		fmt.Printf("type:%T vaule:%v\n", empty, empty)
	}

	// 断言
	{
		var x interface{}
		x = "ssssss"

		v, ok := x.(string)
		if ok {
			fmt.Println(v)
		} else {
			fmt.Println("type assert failed")
		}

		justifyType("sss")
		justifyType(int64(12))
		justifyType(12)
		justifyType(false)
	}
}

// Display ...
func Display(s []interface{}) {
	num := len(s)

	for i := 0; i < num; i++ {
		fmt.Println(s[i])
	}
}

func justifyType(in interface{}) {
	switch v := in.(type) {
	case string:
		fmt.Println("in is a string,", v)
	case int:
		fmt.Println("in is a int,", v)
	case bool:
		fmt.Println("in is a bool,", v)
	default:
		fmt.Println("unsupport type!")
	}
}
