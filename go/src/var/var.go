package main

import "fmt"
import "math"

func main() {
	{
		// var name type
		var i int //default 0
		// var name type = initialvalue
		var j int = 32
		fmt.Println("i =", i, "j =", j)
	}

	{
		// var name type
		var age int
		age = 25
		fmt.Println("my age is", age)
	}

	{
		// var name1, name2 type = initialvalue1, initialvalue2
		var width, height int = 50, 100
		//var width, height = 50, 100
		fmt.Println("width is", width, ", height is", height)
	}

	{
		// var name = initialvalue
		var name = "dzw"
		var age = 25
		fmt.Println("name =", name, "age =", age)
	}

	{
		// must init
		//var name, age, hight = "dzw", 25, 0

		var (
			name = "dzw"
			age = 25
			hight int //default 0
		)

		fmt.Println("name =", name, "age =", age, "hight =", hight)
	}

	{
		// name := initialvalue
		name, age := "dzw", 25
		fmt.Println("name =", name, "age =", age)

		// name, age = "dzw" //error

		// := 操作符左边至少有一个变量是未声明的
		age, hight := 30, 170
		fmt.Println("name =", name, "age =", age, "hight =", hight)

		// no new variables on left side of :=
		//name, age := "dzw", 25
	}

	{
		a, b := 12.5, 23.8
		c := math.Max(a, b)

		fmt.Println("maximum is", c)
	}

	{
		//a := 12
		// cannot use "zwd" (type string) as type int in assignment 
		//a = "zwd"
	}
}

