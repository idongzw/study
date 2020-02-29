/*
 * @Author: dzw
 * @Date: 2020-02-29 19:40:05
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-29 20:42:23
 */

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("1.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}

	// close file
	defer file.Close()

	// read file
	// seekLen, seekErr := file.Seek(0, 1)
	// if seekErr != nil {
	// 	fmt.Println("seek1 failed,", seekLen)
	// 	return
	// }
	// fmt.Println("seek1 len=", seekLen)
	// reader := bufio.NewReader(file)
	// dataLenBytes, err := reader.Peek(8) // Peek 8 之后，文件指针当前位置不是在8
	// if err != nil {
	// 	fmt.Println("data1 len")
	// 	return
	// }
	// fmt.Println(string(dataLenBytes))

	// seekLen, seekErr = file.Seek(0, 0) // 回到开始位置
	// fmt.Println("cur seek ", seekLen)

	// seekLen, seekErr = file.Seek(8, 1)
	// if seekErr != nil {
	// 	fmt.Println("seek2 failed,", seekLen)
	// 	return
	// }
	// seekLen, seekErr = file.Seek(0, 1) // 当前位置
	// fmt.Println("cur seek ", seekLen)

	// reader = bufio.NewReader(file)
	// dataLenBytes, err = reader.Peek(1)
	// if err != nil {
	// 	fmt.Println("data2 len")
	// 	return
	// }
	// fmt.Println(string(dataLenBytes))
	// seekLen, seekErr = file.Seek(0, 1)
	// fmt.Println("cur seek ", seekLen)

	// 一次读取正好4个字节
	bytesSlice := make([]byte, 4)
	for {
		n, err := io.ReadFull(file, bytesSlice)
		if err == io.EOF {
			fmt.Println("n =", n, err)
			fmt.Println(string(bytesSlice[:n]))
			return
		}
		fmt.Println(string(bytesSlice[:n]))
	}
}
