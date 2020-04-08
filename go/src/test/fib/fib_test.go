/*
 * @Author: dzw
 * @Date: 2020-03-27 20:43:08
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-27 20:58:45
 */

package fib

import "testing"

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}

// go test -bench=.
/*
goos: linux
goarch: amd64
BenchmarkFib1-4         403121470                2.87 ns/op
BenchmarkFib2-4         154821326                7.62 ns/op
BenchmarkFib3-4         96675228                12.4 ns/op
BenchmarkFib10-4         2816910               421 ns/op
BenchmarkFib20-4           21072             63019 ns/op
BenchmarkFib40-4               2         815533926 ns/op
PASS
ok      _/home/dzw/golang/study/go/src/test/fib 11.492s
*/
/*
这里需要注意的是，默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。

最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果。
*/

// go test -bench=Fib40 -benchtime=20s
/*
goos: linux
goarch: amd64
BenchmarkFib40-4              28         827141750 ns/op
PASS
ok      _/home/dzw/golang/study/go/src/test/fib 23.979s
*/

func BenchmarkFibParallel(b *testing.B) {
	b.SetParallelism(2) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fib(40)
			// benchmarkFib(b, 40)
		}
	})
}

// go test -bench=Parallel -cpu=1
/*
goos: linux
goarch: amd64
BenchmarkFibParallel           2         791877922 ns/op
PASS
ok      _/home/dzw/golang/study/go/src/test/fib 2.390s
*/

// go test -bench=Parallel -cpu=2
/*
goos: linux
goarch: amd64
BenchmarkFibParallel-2                 3         421923029 ns/op
PASS
ok      _/home/dzw/golang/study/go/src/test/fib 2.963s
*/
