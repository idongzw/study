/*
* @File Name: method.go
* @Author: idongzw
* @Date:   2020-02-16 17:38:36
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-25 17:54:58
*/
package main

import "fmt"

type method struct {
	name string // 首字母大写其他package可以访问
}

type TZ int

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	school string
}

type Worker struct {
	Person
	salary float32
}

func main() {
	m := method{
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

	fmt.Println("------------------------------")

	p := Person{"dzw", 26}
	p.eat()

	s := Student{Person{"dzw", 26}, "xuchang"}
	s.eat()
	s.study()

	w := Worker{Person{"dzw", 23}, 7000.0}
	w.eat()
	w.work()
}

/*
什么时候应该使用指针类型接收者:
1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/

// 注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法

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

func (p Person) eat() {
	fmt.Println("Person method eat()")
}

func (s Student) eat() {
	fmt.Println("Student method eat()")
}

func (s Student) study() {
	fmt.Println("Student method study()")
}

func (w Worker) work() {
	fmt.Println("Worker method work()")
}
