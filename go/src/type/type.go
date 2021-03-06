/*

bool

数字类型
	int8, int16, int32, int64, int
	uint8, uint16, uint32, uint64, uint
	float32, float64
	complex64, complex128
	byte(uint8)
	rune(int32)

string
*/

package main

import (
	"fmt"
	"unsafe" //Sizeof
)

func main() {
	//bool 类型表示一个布尔值，值为 true 或者 false  
	{
		a := true
		b := false
		fmt.Println("a:", a, "b:", b)

		c := a && b
		fmt.Println("c:", c)

		d := a || b
		fmt.Println("d:", d)
	}

	fmt.Println("------------------------------------------------")

	/*
	有符号整型
	int8：表示 8 位有符号整型
	大小：8 位
	范围：-128～127

	int16：表示 16 位有符号整型
	大小：16 位
	范围：-32768～32767

	int32：表示 32 位有符号整型
	大小：32 位
	范围：-2147483648～2147483647

	int64：表示 64 位有符号整型
	大小：64 位
	范围：-9223372036854775808～9223372036854775807

	int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。
	大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
	范围：在 32 位系统下是 -2147483648～2147483647，而在 64 位系统是 -9223372036854775808～9223372036854775807。
	*/
	{
		var a int = 90
		b := 80
		fmt.Println("value of 'a' is:", a, "and 'b' is", b)
		fmt.Printf("type of 'a' is %T, size of 'a' is %d\n", a, unsafe.Sizeof(a)) //a type and value
		fmt.Printf("type of 'b' is %T, size of 'b' is %d\n", b, unsafe.Sizeof(b)) //a type and value
	}

	fmt.Println("------------------------------------------------")

	/*
	无符号整型
	uint8：表示 8 位无符号整型
	大小：8 位
	范围：0～255

	uint16：表示 16 位无符号整型
	大小：16 位
	范围：0～65535

	uint32：表示 32 位无符号整型
	大小：32 位
	范围：0～4294967295

	uint64：表示 64 位无符号整型
	大小：64 位
	范围：0～18446744073709551615

	uint：根据不同的底层平台，表示 32 或 64 位无符号整型。
	大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
	范围：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615。
	*/

	/*
	浮点型
	float32：32 位浮点数
	float64：64 位浮点数
	*/
	{
		a, b := 1.2, 3.5
		fmt.Printf("type of 'a' %T, 'b' %T\n", a, b)

		sum := a + b
		diff := a - b
		fmt.Println("sum:", sum, "diff:", diff)

		c, d := 12, 34
		fmt.Println("sum:", c + d, "diff:", c - d)
	}

	fmt.Println("------------------------------------------------")

	/*
	复数类型
	complex64：实部和虚部都是 float32 类型的的复数。
	complex128：实部和虚部都是 float64 类型的的复数。
	
	内建函数 complex 用于创建一个包含实部和虚部的复数。complex 函数的定义如下：

		func complex(r, i FloatType) ComplexType

	该函数的参数分别是实部和虚部，并返回一个复数类型。实部和虚部应该是相同类型，也就是 float32 或 float64。如果实部和虚部都是 float32 类型，则函数会返回一个 complex64 类型的复数。如果实部和虚部都是 float64 类型，则函数会返回一个 complex128 类型的复数。
	
	还可以使用简短语法来创建复数：
		c := 6 + 7i
	*/
	{
		a := complex(1, 2)
		fmt.Println("a:", a)
		fmt.Printf("type of 'a' %T, size of 'a' %d\n", a, unsafe.Sizeof(a))

		b := 8 + 1i
		c := a + b
		d := a * b
		fmt.Println("b:", b, "c:", c, "d:", d)
	}

	fmt.Println("------------------------------------------------")

	/*
	其他数字类型
	byte 是 uint8 的别名。
	rune 是 int32 的别名。
	*/
	{
		var a byte = 6
		var b rune = 8
		fmt.Printf("type of 'a' %T, b %T, size of 'a' %d, b %d\n", a, b, unsafe.Sizeof(a), unsafe.Sizeof(b))
		fmt.Println("value of a:", a, "b:", b)
	}

	fmt.Println("------------------------------------------------")

	// string 类型
	{
		a := "dzw"
		b := "space"
		c := `Hello`
		fmt.Println(a + " " + b)
		fmt.Println(c)
	}

	fmt.Println("------------------------------------------------")

	/*
	类型转换
	Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换。
	*/
	{
		a := 2
		b := 3.1
		//c := a + b // invalid operation: a + b (mismatched types int and float64)
		// 显式类型装换
		c := a + int(b)
		fmt.Println("a:", a, "b:", b, "c:", c)

		d := 8
		// 显式类型装换
		var e float64 = float64(d)
		fmt.Println("e:", e)

		// 常量在需要的时候，会自动转型
		f1 := 3.14
		sum := f1 + 100
		fmt.Printf("sum: %T,%f\n", sum, sum)

		const c1 = 12
		fmt.Println("c1:", c1)
		const c2 float32 = c1
		fmt.Println("c2:", c2)
		fmt.Printf("c1 = %T,%d\n", c1, c1)
		fmt.Printf("c2 = %T,%f\n", c2, c2)

		// 兼容类型才可以转换
		b1 := false
		i1 := 12
		//i1 = int(b1) // cannot convert b1 (type bool) to type int
		fmt.Println("b1:", b1)
		fmt.Println("i1:", i1)
	}

	fmt.Println("------------------------------------------------")

	/*
	数据类型：
		1. 按数据类型来分：
			基本数据类型：
				int, float, string bool, ...
			复合数据类型：
				array, slice, map, struct, interface, ...
		2. 按照数据的存储特点来分：
			值类型的数据：操作的是数值本身
				int, float, string, bool, array
			引用数据类型：操作的是数据的地址
				slice, map, chan
	 */
	
	{
		var i1 myint = 12
		var i2 int = 13

		fmt.Println("i1 =", i1)
		fmt.Println("i2 =", i2)

		//i1 = i2 // cannot use i2 (type int) as type myint in assignment
		i1 = myint(i2)

		fmt.Printf("i1 type %T, i2 type %T\n", i1, i2) //i1 type main.myint, i2 type int

		var i3 myint2 = 14
		var i4 int = i3
		fmt.Printf("i3 type %T, i4 type %T\n", i3, i4) // i3 type int, i4 type int
	}
}

//定义一个新的类型
type myint int

//类型别名 go version 1.9后
type myint2 = int // 不是重新定义新的数据类型，只是给int起的别名，和int可以通用

//定义函数类型
type myfun func(int, int) string
