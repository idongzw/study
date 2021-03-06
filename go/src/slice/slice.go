/*
* @File Name: slice.go
* @Author: idongzw
* @Date:   2020-02-16 10:42:10
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-24 20:38:08
*/
package main

import "fmt"

// slice
/*
本身并不是数组，它指向底层的数组
作为变长数组的替代方案，可以关联底层数组的局部和全部
为引用类型
可以直接创建或从底层数组获取生成
使用len()获取元素个数，cap()获取容量
一般使用make()创建
如果多个slice指向相同底层数组，其中一个的值改变会影响全部


make([]T, len, cap)
其中cap可以省略，则和len的值相同
len表示存数的元素个数，cap表示容量
*/

func main() {
	{
		var s1 []int
		fmt.Println(s1)
	}

	fmt.Println("1-----------------------------------------")

	{
		a := [10]int{5: 1, 9: 2} // [0 0 0 0 0 1 0 0 0 2]
		fmt.Println(a)

		s1 := a[5:10] // [1 0 0 0 2] //包含下标5 不包含下标10
		s2 := a[5:]   // [1 0 0 0 2]
		s3 := a[:]    // [0 0 0 0 0 1 0 0 0 2]
		fmt.Println(s1, ",s1 len =", len(s1), ",s1 cap =", cap(s1))
		fmt.Println(s2, ",s2 len =", len(s2), ",s2 cap =", cap(s2))
		fmt.Println(s3, ",s3 len =", len(s3), ",s3 cap =", cap(s3))

		s4 := a[:6]
		fmt.Println(s4) // [0 0 0 0 0 1]

		// slice 引用类型
		s3[0] = 3
		fmt.Println(s3, a) // [3 0 0 0 0 1 0 0 0 2] [3 0 0 0 0 1 0 0 0 2]

		s5 := append(s3, s3...)
		fmt.Printf("s3 %p %v\n", s3, s3)
		fmt.Printf("s5 %p %v\n", s5, s5)
		fmt.Printf("a %p %v\n", &a, a)
	}

	fmt.Println("2-----------------------------------------")

	// slice 引用类型
	{
		s1 := []int{1, 2, 3, 4}
		s2 := s1
		fmt.Println(s1, s2) // [1 2 3 4] [1 2 3 4]
		s2[1] = 9
		fmt.Println(s1, s2) // [1 9 3 4] [1 9 3 4]

		s3 := s1[2:]
		fmt.Printf("s1    = %p %v\n", s1, s1)
		fmt.Printf("s1[2] = %p %v\n", &s1[2], s1[2])
		fmt.Printf("s3    = %p %v\n", s3, s3)
	}

	fmt.Println("3-----------------------------------------")

	{
		s1 := make([]int, 3, 10)

		fmt.Println(len(s1), cap(s1), s1) // 3 10 [0 0 0]
	}

	fmt.Println("4-----------------------------------------")

	{
		a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}

		sa := a[2:5]
		sb := a[3:5]
		fmt.Println(sa, sb)                 // [99 100 101] [100 101]
		fmt.Println(string(sa), string(sb)) // cde de

		// reslice时索引以被slice的切片为准
		// cde 索引是 0 1 2
		sas1 := sa[0:]
		fmt.Println(string(sas1)) // cde
		sas2 := sa[3:5]
		fmt.Println(string(sas2), cap(sa)) // fg 9

		// 索引不可以超过被slice的切片的容量cap()值
		//sas3 := sa[9:11] //panic: runtime error: slice bounds out of range
		//fmt.Println(sas3)
		// 索引越界不会导致底层数组的重新分配而是引发错误
	}

	fmt.Println("5-----------------------------------------")

	// append
	/*
	   可以在slice尾部追加元素
	   可以将一个slice追加在另一个slice尾部
	   如果最终长度未超过追加到slice的容量则返回原始slice
	   如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
	*/
	{
		s1 := make([]int, 3, 6)
		fmt.Printf("%p\n", s1)

		s1 = append(s1, 1, 2, 3)
		fmt.Printf("%p %v\n", s1, s1)

		s1 = append(s1, 4, 5, 6)
		fmt.Printf("%p %v\n", s1, s1)

		s2 := make([]int, 2, 3)
		fmt.Println(s2)

		// slice 追加 slice
		s1 = append(s1, s2...)
		fmt.Println(s1)

		s3 := append(s1, s2...)
		fmt.Printf("s1 = %p %v\n", s1, s1)
		fmt.Printf("s3 = %p %v\n", s3, s3)

		// append操作会对该切片做自动初始化扩容操作
		var s4 []int
		s4 = append(s4, 1)
		fmt.Printf("s4 = %p %v\n", s4, s4)
		s4 = append(s4, 2, 3, 4)
		fmt.Printf("s4 = %p %v\n", s4, s4)
	}

	fmt.Println("6-----------------------------------------")

	{
		a := []int{1, 2, 3, 4, 5}
		s1 := a[2:5]
		s2 := a[1:3]
		fmt.Println(s1, s2) // [3 4 5] [2 3]

		s1[0] = 9           // s1 s2 重复元素 3
		fmt.Println(s1, s2) // [9 4 5] [2 9]
	}

	fmt.Println("7-----------------------------------------")

	{
		a := []int{1, 2, 3, 4, 5}
		s1 := a[2:5]
		s2 := a[1:3]
		fmt.Println(s1, s2) // [3 4 5] [2 3]

		// 如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
		// s2 超过自身容量，重新分配，所以不指向 a 数组了
		s2 = append(s2, 6, 6, 6, 6, 6, 6, 6, 6)
		s1[0] = 9           // s1 s2 重复元素 3
		fmt.Println(s1, s2) // [9 4 5] [2 3 6 6 6 6 6 6 6 6]
	}

	fmt.Println("8-----------------------------------------")

	// copy
	{
		s1 := []int{1, 2, 3, 4, 5, 6}
		s2 := []int{7, 8, 9}

		fmt.Println(s1, s2) // [1 2 3 4 5 6] [7 8 9]
		copy(s1, s2)
		fmt.Println(s1, s2) // [7 8 9 4 5 6] [7 8 9]
	}

	fmt.Println("9-----------------------------------------")

	{
		s1 := []int{1, 2, 3, 4, 5, 6}
		s2 := []int{7, 8, 9}

		fmt.Println(s1, s2) // [1 2 3 4 5 6] [7 8 9]
		copy(s2, s1)
		fmt.Println(s1, s2) // [1 2 3 4 5 6] [1 2 3]
	}

	fmt.Println("10----------------------------------------")

	{
		s1 := []int{1, 2, 3, 4, 5, 6}
		s2 := []int{7, 8, 9, 10, 1, 1, 1, 1, 1, 1}

		fmt.Println(s1, s2) // [1 2 3 4 5 6] [7 8 9 10 1 1 1 1 1 1]
		copy(s2[2:4], s1[1:3])
		fmt.Println(s1, s2) // [1 2 3 4 5 6] [7 8 2 3 1 1 1 1 1 1]
	}

	fmt.Println("11----------------------------------------")

	{
		s1 := []int{1, 2, 3, 4}
		s2 := make([]int, 0)

		for i := 0; i < 4; i++ {
			s2 = append(s2, s1[i])
		}
		fmt.Printf("s1 = %p %v\n", s1, s1)
		fmt.Printf("s1 = %p %v\n", s2, s2)

		s1[0] = 9
		fmt.Printf("s1 = %p %v\n", s1, s1)
		fmt.Printf("s1 = %p %v\n", s2, s2)
	}

	fmt.Println("12----------------------------------------")

	{
		array := [...]int{2, 4, 1, 9, 7, 8, 6}
		BubbleSort(array[:])
		fmt.Println(array) // [1 2 4 6 7 8 9]
	}

	fmt.Println("13----------------------------------------")

	// 从切片中删除元素
	/*
	   总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	*/
	{
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Println(s) // [1 2 3 4 5 6 7 8 9]

		//删除索引是2的元素
		s = append(s[:2], s[3:]...)
		fmt.Println(s) // [1 2 4 5 6 7 8 9]
	}
}

func BubbleSort(array []int) {
	num := len(array)

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if array[i] > array[j] {
				tmp := array[i]
				array[i] = array[j]
				array[j] = tmp
			}
		}
	}
}
