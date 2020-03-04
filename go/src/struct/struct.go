/*
* @File Name: struct.go
* @Author: idongzw
* @Date:   2020-02-16 15:01:56
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-25 18:20:57
*/
package main

import (
	"encoding/json"
	"fmt"
)

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
	Age  int
}

type person2 struct {
	Name    string
	Age     int
	Contact struct { // 匿名结构
		Phone, City string
	}
}

type person3 struct {
	string //匿名字段
	int    //匿名字段，默认使用数据类型作为名字
}

/*
匿名字段默认采用类型名作为字段名
结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
*/

type person4 struct {
	string //匿名字段
	int    //匿名字段，默认使用数据类型作为名字
	// 匿名字段的类型不能重复，否则会冲突
	//int //duplicate field int
}

type human struct {
	Sex int
}

type person5 struct {
	h human // 非匿名字段，模拟聚合关系 has a
	s string
}

type teacher struct {
	human // 匿名字段，模拟继承性 is a
	Name  string
	Age   int
}

//当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找

type student struct {
	human
	Sex  int
	Name string
	Age  int
}

/*
结构体标签（Tag）:
Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。
Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
	`key1:"value1" key2:"value2"`
结构体标签由一个或多个键值对组成。
键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。

注意事项：
为结构体编写Tag时，必须严格遵守键值对的规则。
结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。
例如不要在key和value之间添加空格。
*/

// Student
type Student struct {
	ID     int    `json:"id"` // 通过指定tag实现json序列化该字段时的key
	Gender string // json序列化默认使用字段名作为key
	Name   string
	//score float32 // 私有不能被json包访问
}

// Class
type Class struct {
	Title   string
	Student []*Student
}

func main() {
	p := person{}
	fmt.Println(p) // { 0}
	p.Name = "dzw"
	p.Age = 26
	fmt.Println(p) // {dzw 26}

	p2 := person{
		Name: "Joe",
		Age:  18,
	}
	fmt.Println(p2) // {Joe 18}

	A(&p2)
	fmt.Println(p2)

	p3 := &person{ // 初始化尽量使用这种方式
		Name: "asd",
		Age:  19,
	}
	p3.Name = "assas"

	fmt.Println(p3)
	A(p3)
	fmt.Println(p3)

	// 匿名结构
	p4 := struct {
		Name string
		Age  int
	}{
		Name: "wawda",
		Age:  19,
	}

	fmt.Println(p4)

	p5 := &struct {
		Name string
		Age  int
	}{
		Name: "wawda",
		Age:  19,
	}

	fmt.Println(p5)

	p6 := person2{
		Name: "qaz",
		Age:  90,
	}
	p6.Contact.Phone = "1233333333"
	p6.Contact.City = "BJ"

	fmt.Println(p6)

	p7 := person3{
		"Joe",
		18,
	}
	fmt.Println(p7)
	fmt.Println(p7.string, p7.int)

	p8 := person{
		Name: "Joe",
		Age:  18,
	}

	p9 := person{
		Name: "Joe",
		Age:  18,
	}

	fmt.Println(p8 == p9)

	fmt.Println("-----------------------------")

	p10 := teacher{
		Name:  "asd",
		Age:   27,
		human: human{Sex: 2},
	}

	fmt.Println(p10) // {{2} asd 27}
	p10.human.Sex = 100
	fmt.Println(p10) // {{100} asd 27}
	p10.Sex = 10
	fmt.Println(p10) // {{10} asd 27}

	p11 := student{
		Name: "sad",
		Age:  12,
		Sex:  23,
	}
	fmt.Println(p11) // {{0} 23 sad 12}

	p12 := student{
		Name:  "sada",
		Age:   12,
		Sex:   233,
		human: human{Sex: 1},
	}
	fmt.Println(p12) // {{1} 233 sada 12}
	p12.Sex = 4
	fmt.Println(p12) // {{1} 4 sada 12}
	p12.human.Sex = 23
	fmt.Println(p12) // {{23} 4 sada 12}

	// json序列化
	{
		c := &Class{
			Title:   "101",
			Student: make([]*Student, 0, 200),
		}

		for i := 0; i < 10; i++ {
			stu := &Student{
				Name:   fmt.Sprintf("stu%02d", i),
				Gender: "F",
				ID:     i,
			}
			c.Student = append(c.Student, stu)
		}

		// Json序列化：结构体--->Json格式的字符串
		data, err := json.Marshal(c)
		if err != nil {
			fmt.Println("json marshal failed,", err)
			return
		}
		fmt.Printf("json: %s\n", data)

		// Json反序列化：Json格式的字符串--->结构体
		str := `{"Title":"101","Student":[{"ID":0,"Gender":"F","Name":"stu00"},
        {"ID":1,"Gender":"F","Name":"stu01"},{"ID":2,"Gender":"F","Name":"stu02"},
        {"ID":3,"Gender":"F","Name":"stu03"},{"ID":4,"Gender":"F","Name":"stu04"},
        {"ID":5,"Gender":"F","Name":"stu05"},{"ID":6,"Gender":"F","Name":"stu06"},
        {"ID":7,"Gender":"F","Name":"stu07"},{"ID":8,"Gender":"F","Name":"stu08"},
        {"ID":9,"Gender":"F","Name":"stu09"}]}`

		c1 := &Class{}
		err = json.Unmarshal([]byte(str), c1)
		if err != nil {
			fmt.Println("json unmarshal failed,", err)
			return
		}
		fmt.Printf("%#v\n", c1)
		fmt.Println("Title:", c1.Title)
		fmt.Println("Student:", c1.Student)
	}

	// tag
	{
		s := Student{
			ID:     1,
			Gender: "F",
			Name:   "dzw",
		}

		data, err := json.Marshal(s)
		if err != nil {
			fmt.Println("json marshal failed,", err)
			return
		}
		fmt.Printf("json str:%s\n", data) // json str:{"id":1,"Gender":"F","Name":"dzw"}
	}
}

func A(p *person) {
	p.Age = 12
	fmt.Println(p)
}
