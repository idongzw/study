/*
* @File Name: pointer.go
* @Author: idongzw
* @Date:   2020-02-20 13:53:27
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-20 16:06:18
 */

package main

import "fmt"

/*
数组指针：首先是一个指针，一个数组的地址
    *[num]Type
    *
指针数组：首先是一个数组，存储的数据类型是指针
    [num]*Type


*[5]float64:指针，存储5个float64类型数据的数组的指针
*[3]string:指针，存储3个string类型数据的数组的指针
[3]*string:数组，存储3个string的指针地址的数组
[5]*float64:数组，存储5个float64的指针地址的数组
*[5]*float64:指针
*[3]*string:指针
**[4]string:指针
**[4]*string:指针

函数指针：一个指针，指向了一个函数的指针
    go语言中，function，默认看作一个指针，没有 *

    slice,map,function 存储的就是数据的地址

指针函数：一个函数，该函数的返回值是一个指针


指针作为参数：
    主要用于值传递类型

参数的传递：值传递，引用传递
*/
func main() {
	{
		arr1 := [4]int{1, 2, 3, 4}
		fmt.Println(arr1)

		// 数组指针
		var p1 *[4]int
		p1 = &arr1
		fmt.Println(p1) // &[1 2 3 4]
		fmt.Println(*p1)

		(*p1)[0] = 1000
		fmt.Println(arr1)

		p1[0] = 2000 // 简化写法
		fmt.Println(arr1)
	}

	fmt.Println("++++++++++++++++++++++++")

	{
		a := 1
		b := 2
		c := 3
		d := 4

		arr1 := [4]int{a, b, c, d} // 值传递
		// 指针数组
		arr2 := [4]*int{&a, &b, &c, &d}

		fmt.Println(arr1)
		fmt.Println(arr2)

		*arr2[0] = 100
		fmt.Println(arr1)
		fmt.Println(arr2)
		fmt.Println(a, b, c, d)

		b = 200
		fmt.Println(arr1)
		fmt.Println(arr2)
		fmt.Println(*arr2[1])
		fmt.Println(a, b, c, d)
	}

	fmt.Println("++++++++++++++++++++++++")

	{
		var f1 func()
		f1 = func1
		f1()

		arr1 := func2()
		fmt.Printf("arr1 type: %T, address: %p, value: %v\n", arr1, &arr1, arr1)

		arr2 := func3()
		fmt.Printf("arr2 type: %T, address: %p, value: %v\n", arr2, &arr2, arr2)
		fmt.Printf("arr2 value %p\n", arr2)
	}

	fmt.Println("++++++++++++++++++++++++")

	{
		//make返回初始化后的（非零）值
		s := make([]int, 3)
		if s == nil {
			fmt.Println("s == nil")
		}
		fmt.Println(s)

		i := new(int)
		fmt.Println(i, *i)

		//new(T)分配了零值填充的T类型的内存空间，并且返回其地址
		s_new := new([]int)
		fmt.Println(s_new, *s_new)
		if *s_new == nil {
			fmt.Println("*s_new == nil")
		}

		*s_new = append(*s_new, s...)
		fmt.Println(s_new)
	}

	fmt.Println("++++++++++++++++++++++++")

	{
		//指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值
		var pi *int
		pi = new(int)
		*pi = 100 // invalid memory address or nil pointer dereference
		fmt.Println(*pi, pi)

		//对于引用类型的变量，要分配内存空间
		var m map[string]int
		m = make(map[string]int)
		m["dzw"] = 1 // assignment to entry in nil map
		fmt.Println(m)
	}

}

func func1() {
	fmt.Println("func1...")
}

func func2() [4]int {
	return [4]int{1, 2, 3, 4}
}

// 指针函数
func func3() *[4]int {
	//return &[4]int{5,6,7,8}
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("func3 arr address %p\n", &arr)

	return &arr
}
