/*
 * @Author: dzw
 * @Date: 2020-03-06 18:28:48
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-06 21:27:46
 */

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	x       = 0
	y       = 0
	wg      sync.WaitGroup
	mutex   sync.Mutex   // 互斥锁
	rwMutex sync.RWMutex // 读写锁
)

// 互斥锁
// 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
// 当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
func add() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		mutex.Lock() // 临界操作加锁
		x++
		mutex.Unlock()
	}
}

// 读写锁
/*
读写锁分为两种：读锁和写锁。
当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待
*/
func read() {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Println("read start...")
	fmt.Println("read data =", y)
	time.Sleep(time.Second * 2)
	rwMutex.RUnlock()
	fmt.Println("read end...")
}

func write() {
	defer wg.Done()
	rwMutex.Lock()
	fmt.Println("write start...")
	y++
	time.Sleep(time.Second * 1)
	rwMutex.Unlock()
	fmt.Println("write end...")
}

func set(m map[int]int, k, v int) {
	m[k] = v
}

func get(m map[int]int, k int) int {
	return m[k]
}

// Test sync.Once
func fOnce(once *sync.Once, i int) {
	defer wg.Done()
	fmt.Println("------------test once-------- ", i)
	once.Do(func() {
		fmt.Println("111111111111  ", i)
	}) // 只能是无参函数
}

func main() {
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println("x =", x)

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go write()
	// }

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go read()
	// }

	// wg.Wait()

	// 内置map
	// {
	// 	wg1 := sync.WaitGroup{}
	// 	m := make(map[int]int)
	// 	for i := 0; i < 20; i++ {
	// 		wg1.Add(1)
	// 		go func(n int) {
	// 			set(m, n, n) // concurrent map writes
	// 			fmt.Println("k =", n, ",v =", get(m, n))
	// 			wg1.Done()
	// 		}(i)
	// 	}

	// 	wg1.Wait()
	// }

	// sync.Map 并发安全的map
	m := sync.Map{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			v, _ := m.Load(key)
			fmt.Println("k =", key, ",v =", v)
			wg.Done()
		}(i)
	}

	wg.Wait()

	wg.Add(1)
	go func(sm *sync.Map) {
		sm.Range(func(key, vaule interface{}) bool {
			fmt.Println("key =", key, ",value =", vaule)
			return true
		})
		wg.Done()
	}(&m)
	wg.Wait()

	// test sync.Once
	once := &sync.Once{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fOnce(once, i)
	}
	wg.Wait()
}
