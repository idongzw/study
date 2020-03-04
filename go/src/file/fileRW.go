/*
 * @Author: dzw
 * @Date: 2020-03-02 15:55:11
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-02 16:02:55
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if err := CopyFile("2.txt", "file.go"); err != nil {
		fmt.Println("copy file failed,", err)
	} else {
		fmt.Println("copy file.go to 2.txt success")
	}

	// 0666
	perm := 0666 // 8进制
	fmt.Printf("0666 type:%T, vaule:%v\n", perm, perm)

	// 1. Write和WriteString
	if err := fileWrite("1.txt"); err != nil {
		fmt.Println("write file 1.txt failed,", err)
	} else {
		fmt.Println("write file 1.txt success")
	}

	// 2. bufio.NewWriter
	if err := fileWriteBufio("3.txt"); err != nil {
		fmt.Println("write file 3.txt failed,", err)
	} else {
		fmt.Println("write file 3.txt success")
	}

	// 3. ioutil.WriteFile
	if err := fileWriteIoutil("4.txt"); err != nil {
		fmt.Println("write file 4.txt failed,", err)
	} else {
		fmt.Println("write file 4.txt success")
	}

	// 1. file.Read() 读取最多N个字节
	if err := fileRead("test.txt"); err != nil {
		fmt.Println("1. read file test.txt failed,", err)
	}

	// 2. 读取正好N个字节
	// io.ReadFull()
	if err := fileReadIoReadFull("test.txt"); err != nil {
		fmt.Println("2. read file test.txt failed,", err)
	}

	// 3. 读取至少N个字节
	// io.ReadAtLeast
	if err := fileReadIoReadAtLeast("test.txt"); err != nil {
		fmt.Println("3. read file test.txt failed,", err)
	}

	// 4. 读取全部字节
	// ioutil.ReadAll
	if err := fileReadIoutilReadAll("test.txt"); err != nil {
		fmt.Println("4. read file test.txt failed,", err)
	}

	// 5. 快读到内存
	// ioutil.ReadFile
	if err := fileReadIoutilReadFile("test.txt"); err != nil {
		fmt.Println("5. read file test.txt failed,", err)
	}

	// 6. 使用缓存读
	// bufio.NewReader
	if err := fileReadBufio("test.txt"); err != nil {
		fmt.Println("6. read file test.txt failed,", err)
	}

	// Scanner
	if err := fielReadScanner("test.txt"); err != nil {
		fmt.Println("read file test.txt failed,", err)
	}

	// remove file
	os.Remove("1.txt")
	fmt.Println("remove 1.txt")
	os.Remove("2.txt")
	fmt.Println("remove 2.txt")
	os.Remove("3.txt")
	fmt.Println("remove 3.txt")
	os.Remove("4.txt")
	fmt.Println("remove 4.txt")
}

// CopyFile ...
func CopyFile(dst string, src string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	n, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	fmt.Printf("Copy %d bytes.", n)

	// err = dstFile.Sync()
	// if err != nil {
	// 	return err
	// }

	return nil
}

// Write file
// 1. Write和WriteString
func fileWrite(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte("abcdefg中文"))
	if err != nil {
		return err
	}
	_, err = file.WriteString("123456789")
	if err != nil {
		return err
	}

	return nil
}

// 2. bufio.NewWriter
// 带缓存的writer，处理很多的数据时，可以节省操作硬盘I/O的时间
func fileWriteBufio(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	bufferWriter := bufio.NewWriter(file)
	n, err := bufferWriter.Write([]byte("abcdefghijk"))
	if err != nil {
		return err
	}
	fmt.Printf("write %d bytes.\n", n)

	err = bufferWriter.WriteByte('1')
	if err != nil {
		return err
	}

	n, err = bufferWriter.WriteString("aaaaaaaaaaaaaaaa\n")
	if err != nil {
		return err
	}
	fmt.Printf("write %d bytes.\n", n)

	// 检查缓存中的字节数
	unflushBufferSize := bufferWriter.Buffered()
	fmt.Printf("Bytes buffered: %d\n", unflushBufferSize)

	// 还有多少字节可用
	bytesAvailable := bufferWriter.Available()
	fmt.Printf("1. Available buffer size: %d\n", bytesAvailable)

	// 写内存buffer到disk
	bufferWriter.Flush()

	// 丢弃还没有flush的缓存的内容
	// 你想将缓存传给另外一个writer时有用
	bufferWriter.Reset(bufferWriter)
	bytesAvailable = bufferWriter.Available()
	fmt.Printf("2. Available buffer size: %d\n", bytesAvailable)

	// 重新设置缓存大小
	// 默认缓存大小 4096
	bufferWriter = bufio.NewWriterSize(bufferWriter, 8000)

	// resize 之后检查缓存大小
	bytesAvailable = bufferWriter.Available()
	fmt.Printf("3. Available buffer size: %d\n", bytesAvailable)

	return nil
}

// 3. ioutil.WriteFile
func fileWriteIoutil(filename string) error {
	err := ioutil.WriteFile(filename, []byte("1234567890"), 0666)
	if err != nil {
		return err
	}

	return nil
}

// Read file
// 1. file.Read() 读取最多N个字节
func fileRead(filename string) error {
	file, err := os.Open("test.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	bufSlice := make([]byte, 8) // 一次最多读取 8 bytes
	data := make([]byte, 0, 128)
	for {
		n, err := file.Read(bufSlice)
		if err == io.EOF {
			fmt.Println("read file done. n =", n, err)
			break
		}

		if err != nil {
			return err
		}

		fmt.Printf("read %d bytes from file.\n", n)
		data = append(data, bufSlice[:n]...)
	}
	fmt.Println(string(data))

	return nil
}

// 2. 读取正好N个字节
// io.ReadFull()
func fileReadIoReadFull(filename string) error {
	file, err := os.Open("test.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	bufSlice := make([]byte, 8) // 如果从文件中读取的字节数小于 8 bytes，则返回错误
	data := make([]byte, 0, 128)
	for {
		n, err := io.ReadFull(file, bufSlice)

		if n != len(bufSlice) {
			fmt.Printf("read bytes %d < %d\n", n, len(bufSlice))
		}

		if err != nil {
			fmt.Println("read file failed. n =", n, err)
			return err
		}

		fmt.Printf("read %d bytes from file.\n", n)
		data = append(data, bufSlice[:n]...)
	}
	fmt.Println(string(data))

	return nil
}

// 3. 读取至少N个字节
// io.ReadAtLeast
func fileReadIoReadAtLeast(filename string) error {
	file, err := os.Open("test.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	bufSlice := make([]byte, 128)
	data := make([]byte, 0, 128)
	minBytes := 50 // 一次至少读取50bytes

	for {
		// io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的数据保存
		n, err := io.ReadAtLeast(file, bufSlice, minBytes)
		if err == io.EOF {
			fmt.Println("read file done. n =", n, err)
			break
		}

		if err != nil {
			// break
			return err
		}

		fmt.Printf("read %d bytes from file.\n", n)
		data = append(data, bufSlice[:n]...)
	}
	fmt.Println(string(data))
	fmt.Println("read data:", string(bufSlice))

	return nil
}

// 4. 读取全部字节
// ioutil.ReadAll
func fileReadIoutilReadAll(filename string) error {
	file, err := os.Open("test.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Printf("read %d bytes from file.\n", len(data))
	fmt.Println("read data:", string(data))
	return nil
}

// 5. 快读到内存
// ioutil.ReadFile
func fileReadIoutilReadFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Printf("read %d bytes from file.\n", len(data))
	fmt.Println("read data:", string(data))

	return nil
}

// 6. 使用缓存读
// bufio.NewReader
func fileReadBufio(filename string) error {
	file, err := os.Open("test.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	bufferReader := bufio.NewReader(file)
	// bufSlice := make([]byte, 5)
	// bufSlice, err = bufferReader.Peek(5)
	// n, _ := file.Seek(0, 1)
	// fmt.Println(string(bufSlice), ",current file point:", n)
	// file.Seek(5, 0)
	for {
		line, err := bufferReader.ReadString('\n') // 读取到 '\n'，包含分隔符，返回字符串
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(line)
	}

	return nil
}

// Scanner
// Scanner是bufio包下的类型,在处理文件中以分隔符分隔的文本时很有用
func fielReadScanner(filename string) error {
	file, err := os.Open("test.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // 参数为函数

	for {
		ok := scanner.Scan()
		if !ok {
			err = scanner.Err()
			if err != nil {
				return err
			}
			fmt.Println("scan completed and reached EOF")
			break
		}

		fmt.Println("scanner word found:", scanner.Text())
	}

	return nil
}
