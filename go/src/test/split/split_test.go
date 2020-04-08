/*
 * @Author: dzw
 * @Date: 2020-03-27 18:29:39
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-27 21:29:49
 */

package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误信息
	}
}

func TestMoreSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("abcd", "bc")
	want := []string{"a", "d"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误信息
	}
}

// 测试组
func TestSplit2(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "上海自来水来自海上", sep: "来", want: []string{"上海自", "水", "自海上"}},
	}

	// 遍历，逐一进行测试
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
		}
	}
}

func TestSplit3(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义一个存储测试用例的map
	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "上海自来水来自海上", sep: "来", want: []string{"上海自", "水", "自海上"}},
	}

	// 遍历，逐一进行测试
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s excepted:%#v, got:%#v", name, tc.want, got)
		}
	}
}

// 子测试 Go1.7+
// t.Run
func TestSplit4(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义一个存储测试用例的map
	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "上海自来水来自海上", sep: "来", want: []string{"上海自", "水", "自海上"}},
	}

	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	// 遍历，逐一进行测试
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行testdoen操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

// go test -bench=Split
/*
goos: linux
goarch: amd64
BenchmarkSplit-4         3509728               340 ns/op
PASS
ok      _/home/dzw/golang/study/go/src/test/split       1.549s

BenchmarkSplit-4 表示对Split函数进行基准测试，数字4表示GOMAXPROCS的值
3509728和340ns/op表示每次调用Split函数耗时304ns，这个结果是3509728次调用的平均值
*/

// 为基准测试添加-benchmem参数，来获得内存分配的统计数据
// go test -bench=Split -benchmem
/*
goos: linux
goarch: amd64
BenchmarkSplit-4         3154197               333 ns/op             112 B/op          3 allocs/op
PASS
ok      _/home/dzw/golang/study/go/src/test/split       1.445s

112 B/op表示每次操作内存分配了112字节，3 allocs/op则表示每次操作进行了3次内存分配
*/

// 优化
/*
使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加
*/
func BenchmarkSplit1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split1("a:b:c", ":")
	}
}

// 这个使用make函数提前分配内存的改动，减少了2/3的内存分配次数，并且减少了一半的内存分配
// go test -bench=Split -benchmem
/*
goos: linux
goarch: amd64
BenchmarkSplit-4         3335353               349 ns/op             112 B/op          3 allocs/op
BenchmarkSplit1-4        7697913               152 ns/op              48 B/op          1 allocs/op
PASS
ok      _/home/dzw/golang/study/go/src/test/split       2.866s
*/

// 有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown
// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}

func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}
