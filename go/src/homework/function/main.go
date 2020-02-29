/*
 * @Author: dzw
 * @Date: 2020-02-25 15:00:46
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-25 16:33:42
 */

package main

import (
	"fmt"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 6

	/*
		defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
			A 1 2 3
			B 10 2 12
			BB 10 12 22
			AA 1 3 4
	*/
	{
		x := 1
		y := 2
		defer calc("AA", x, calc("A", x, y))
		x = 10
		defer calc("BB", x, calc("B", x, y))
		y = 20
	}

	left := dispatchCoin()
	for k, v := range distribution {
		fmt.Println(k, "get coins:", v)
	}
	fmt.Println("left:", left)
}

// defer 执行时机
/*
在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
而defer语句执行的时机就在返回值赋值操作后，RET指令执行前

return x
1. 返回值=x
2. RET指令

defer 语句执行时机
return x
1. 返回值=x
2. 运行defer
3. RET指令
*/

// defer 经典案例
func f1() int {
	x := 5
	defer func() {
		x++
	}()

	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()

	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()

	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)

	return 5
}

func f5() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)

	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

func dispatchCoin() int {
	for _, v := range users {
		if _, ok := distribution[v]; !ok {
			distribution[v] = 0
		}
		for _, value := range v {
			switch value {
			case 'e', 'E':
				distribution[v]++
				coins--
			case 'i', 'I':
				coins -= 2
				distribution[v] += 2
			case 'o', 'O':
				coins -= 3
				distribution[v] += 3
			case 'u', 'U':
				coins -= 4
				distribution[v] += 4
			}
		}
	}

	return coins
}
