/*
* @File Name: basic_structure.go
* @Author: idongzw
* @Date:   2020-02-15 15:34:46
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-21 12:15:05
*/

// go 程序的一般结构

// 当前程序包名
package main

// 导入其他包
//import "fmt"
//import std "fmt" //alias
//import . "fmt"

// 导入多个包

import (
    "fmt"
    "math"
    "strconv"
    "time"
    "errors"
    "os"
    "net"
)

// 常量定义
//const PI = 3.14
// 定义多个常量

const (
    PI = 3.14
    Name = "dzw"
)


// 全局变量的声明与赋值
//var name = "dzw"
// 定义多个全局变量

var (
    name = "dzw"
    age = 26
)


// 一般类型声明
//type newType int
// 声明多个一般类型

type (
    newType1 int
    newType2 float32
)

// 结构的声明
type gostruct struct {}

// 接口的声明
type gointerface interface {}

type 文本 string
var chinese 文本 = "中文"

// main 函数作为程序入口
func main() {
    fmt.Println("Hello,你好")
    //std.Println("Hello,你好") //alias
    //Println("Hello,你好")
    
    fmt.Println(chinese)

    fmt.Println(math.MaxInt8)

    a := 65
    // string 表示将数据转换为文本格式，数字65表示文本A
    b := string(a)
    fmt.Println(a, b)

    b = strconv.Itoa(a)
    fmt.Println(b)
    a, _ = strconv.Atoi(b)
    fmt.Println(a)

    fmt.Println("++++++++++++++++++++++++++++++")
    //需要注意的地方(坑)
    /*
    1. slice append 超过容量时会自动扩展新建一个slice，先拷贝原来的值再append
    */
    {
        s := make([]int, 0)
        fmt.Println(s)
        s = Append(s)
        fmt.Println(s)
    }

    /*
    2. time
     */
    {
        t := time.Now()
        fmt.Println(t)
        fmt.Println(t.Format(time.ANSIC))
        fmt.Println(time.ANSIC)
        fmt.Println(t.Format("Mon Jan _2 15:04:05 2006")) //起始时间不能修改
    }

    /*
    3. for range
     */
    {
        /*
        s := []string{"a", "b", "c"}

        // 输出 c c c
        // 不按参数传递的值，都是引用
        for _, v := range s {
            go func() {
                fmt.Println(v)
            }()
        }

        // 输出 a b c
        // 按参数传递
        for _, v := range s {
            go func(v string) {
                fmt.Println(v)
            }(v)
        }
        */

//        select {}
    }

    // 创建error信息两种方式
    {
        err1 := errors.New("my errors") // *errors.errorString
        fmt.Printf("err1 type (%T), err1 = (%v)\n", err1, err1)

        err2 := fmt.Errorf("error code = %d", 100) // *errors.errorString
        fmt.Printf("err2 type (%T), err2 = (%v)\n", err2, err2)
    }

    {
        err := checkAge(-20)
        if err != nil {
            fmt.Println(err)
            //return
        }
        fmt.Println("run...")
    }

    // open file
    {
        f, err := os.Open("test.txt")

        if err != nil {
            fmt.Println(err) // open test.txt: no such file or directory
            if ins, ok := err.(*os.PathError); ok {
                fmt.Println("1.Op:", ins.Op) // 1.Op: open
                fmt.Println("2.Path:", ins.Path) // 2.Path: test.txt
                fmt.Println("3.Err:", ins.Err) // 3.Err: no such file or directory
            }
            return
        }
        fmt.Println(f.Name(),"open success")
    }

    {
        addr, err := net.LookupHost("45.76.240.193")
        fmt.Println(addr, err)
    }

    {
        age, err := testMyError(-30)
        if err != nil {
            fmt.Println(err) // error code: 100, error msg: age is illegal
        }
        fmt.Println(age)
    }
}

/*
error: 内置的数据类型，内置的接口
    定义方法：Error() string

使用go语言提供好的包：
    errors包下提供的函数：New()，创建一个error对象
    fmt包下的Errorf()函数：
        func Errorf(format string, a ...interface{}) error
 */

// error
func checkAge(age int) error {
    if age < 0 {
        //return errors.New("Illegal age")
        return fmt.Errorf("Age %d is illegal", age)
    }

    fmt.Println("Age is:", age)
    return nil
}

// 自定义 error
type AgeError struct {
    errCode int
    errString string
}

func (e *AgeError) Error() string {
    return fmt.Sprintf("error code: %d, error msg: %s", e.errCode, e.errString)
}

func testMyError(age int) (int, error) {
    if age < 0 {
        return -1, &AgeError{100, "age is illegal"}
    }

    return age, nil
}

// 需设置返回值
func Append(s []int) []int {
    s = append(s, 3)
    return s
}

// 可见性规则
/*
go语言中，使用 大小写 来决定该 常量、变量、类型、接口、结构或函数是否可以被外部包所调用
根据预定，函数名首字母小写为private，首字母大写为public
访问权限是 package与package 之间
 */

//alias
//byte 是 uint8 的别名
//rune 是 int32 的别名

// int
// 和操作系统位数有关 32/64

// float32/float64
//  -长度：4/8 字节
//  -小数位：精确到 7/15 小数位

// uintptr
// 32/64位整数类型 ，保存指针


// 类型零值
/* 零值并不等于空值，而是当变量被声明为某种类型后的默认值，
通常情况下值类型的默认值为 0 ，bool 为 false， string 为空字符串
*/


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

/*
关于包的使用：
    1. 一个目录下的同级文件归属一个包，package的声明要一致
    2. package声明的包和对应的目录名可以不一致，但习惯上还是要写一致
    3. 包可以嵌套
    4. 同包下的函数不需要导入包，可以直接使用
    5. main包，main()函数所在的包，其他包不能使用
    6. 导入包的时候，路径要从 src 下开始写


    import 包时:
    1. 点操作：
        import (
            . "fmt"
        )
    这个点操作的含义是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名
    eg. fmt.Println("hello world") 可以省略写成 Println("hello world")

    2. 起别名
        import (
            f "fmt"
        )
        eg. f.Println("hello world")

    3. _ 操作
    如果仅仅需要导入包时执行初始化操作，并不需要使用包内的其它函数，常量等资源。则可以在导入包时，匿名导入。
        import (
            _ "fmt"
        )
        _ 操作其实是引入该包，而不直接使用包里边的函数，而是调用该包里边的init函数。也就是说，使用 _ 作为包
        的别名，会仅仅执行 init()

    导入的包的路径名，可以是相对路径也可以是绝对路径，推荐使用绝对路径（起始于工程根目录）
 */

