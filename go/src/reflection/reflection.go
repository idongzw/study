/*
* @File Name: reflection.go
* @Author: idongzw
* @Date:   2020-02-17 12:11:09
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-17 21:05:09
*/
package main

import (
    "fmt"
    "reflect"
)

/*
反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
反射会将匿名字段作为独立字段（匿名字段本质）
想要利用反射修改对象状态，前提是 interface.data 是 settable，即 pointer-interface
通过反射可以“动态”调用方法
 */

type user struct {
    Id int
    Name string
    Age int
}

func (u user) Hello(name string) {
    fmt.Println("Hello", name, ", my name is", u.Name)
}

func (u *user) SetId(id int) {
    u.Id = id
}

func (u *user) SetName(name string) {
    u.Name = name
}

func (u *user) SetAge(age int) {
    u.Age = age
}

func (u user) GetId() int {
    return u.Id
}

func (u user) GetName() string {
    return u.Name
}

func (u user) GetAge() int {
    return u.Age
}

func (u *user) DisplayInfo() {
    fmt.Println("  Id:", u.Id)
    fmt.Println("Name:", u.Name)
    fmt.Println(" Age:", u.Age)
}

func Info(o interface{}) {
    // type
    t := reflect.TypeOf(o)
    fmt.Println(t)
    fmt.Println("Type:", t.Name()) // user

    if k := t.Kind(); k != reflect.Struct {
        fmt.Println("Type != reflect.Struct")

        return
    }

    // value
    v := reflect.ValueOf(o)
    fmt.Println("Fields:")

    vf := v.NumField()
    fmt.Println("vf = ", vf)

    // type field
    for i := 0; i < t.NumField(); i++ {
        // type field
        f := t.Field(i)

        // value field
        val := v.Field(i).Interface()

        fmt.Printf("%6s: %v = %v\n", f.Type, f.Name, val)
    }

    // type method
    for i := 0; i < t.NumMethod(); i++ {
        m := t.Method(i)

        fmt.Printf("%s: %v\n", m.Type, m.Name)
    }
}

type Manager struct {
    user // 匿名字段
    title string
}

// 修改 struct 值
func Set(o interface{}) {
    v := reflect.ValueOf(o)

    // Kind 判断类型
    // 只有 o 是指针，才可以通过 reflect修改实际值
    if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
        fmt.Println("XXX")

        return
    } else {
        v = v.Elem()
    }

    /*
    if f := v.FieldByName("Name"); f.Kind() == reflect.String {
        f.SetString("setName")
    }
    */
   
    f := v.FieldByName("Name")
    if !f.IsValid() {
        fmt.Println("not exists")

        return 
    }

    if f.Kind() == reflect.String {
        f.SetString("SetName")
    }
}

func main() {
    // 已知原有类型【进行“强制转换”】
    // realValue := value.Interface().(已知的类型)
    {
        var num float64 = 1.23

        pointer := reflect.ValueOf(&num)
        value := reflect.ValueOf(num)
        fmt.Printf("%#v\n%#v\n", pointer, value)

        convertPointer := pointer.Interface()
        //convertPointer := pointer.Interface().(*float64)
        fmt.Println(convertPointer)
        convertValue := value.Interface()
        //convertValue := value.Interface().(float64)
        fmt.Println(convertValue)
    }

    fmt.Println("------------------------------")

    {
        u := user{1, "dzw", 26}
        fmt.Println(u)

        Info(&u) // &u == reflect.Ptr
        Info(u)  // u == reflect.Struct
    }

    fmt.Println("----------------------------------------")

    {
        m := Manager{user: user{1, "Ok", 2}, title: "123"}
        fmt.Println(m)

        t := reflect.TypeOf(m)
        v := reflect.ValueOf(m)

        for i := 0; i < t.NumField(); i++ {
            fmt.Printf("%#v\n", t.Field(i))
        }

        for i := 0; i < v.NumField(); i++ {
            fmt.Println("++++++")
            fmt.Printf("%#v\n", v.Field(i))
            // title 首字母小写 不支持 Interface() ,
            // 首字母大写支持 Interface()
            fmt.Println(v.Field(i).CanInterface())
            //fmt.Println(v.Field(i).Interface())
        }

        // user Id
        fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))

        // user Name
        fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

        // user Age
        fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 2}))

        // title
        fmt.Printf("%#v\n", t.FieldByIndex([]int{1}))


        fmt.Printf("%#v\n", v)
        //val := v.FieldByName("title")
        val := v.FieldByIndex([]int{1})
        fmt.Println(val)
    }

    // 修改变量值
    {
        x := 123
        v := reflect.ValueOf(&x)
        v.Elem().SetInt(999)

        fmt.Println(x)
    }

    {
        u := user{1, "dzw", 11}
        fmt.Println(u)
        Set(&u)
        fmt.Println(u)
    }

    // 调用函数
    {
        u := user{1, "dzw", 123}
        v := reflect.ValueOf(u)

        mv_hello := v.MethodByName("Hello")
        args := []reflect.Value{reflect.ValueOf("Joe")}
        mv_hello.Call(args)

        mv_getname := v.MethodByName("GetName")
        args = []reflect.Value{}
        name := mv_getname.Call(args)
        fmt.Println(name, reflect.ValueOf(name).Kind())

        v_r := reflect.ValueOf(&u)
        mv_setid := v_r.MethodByName("SetId")
        args = []reflect.Value{reflect.ValueOf(8)}
        mv_setid.Call(args)
        fmt.Println(u)
        fmt.Println(v_r.Kind())
    }
}