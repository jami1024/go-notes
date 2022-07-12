package goroutine_test

import (
	"fmt"
	"testing"
)

func TestGoroutineStatement(t *testing.T) {

	fmt.Println("如何创建一个Goroutine")
	go testgo()
	for i := 0; i <= 10; i++ {
		// time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	fmt.Println("函数结束")
}

//⾃定义函数
func testgo() {
	for i := 0; i <= 1000; i++ {
		// time.Sleep(10 * time.Second)
		fmt.Println("测试goroutine", i)
	}
}

/*
=== RUN   TestGoroutineStatement
如何创建一个Goroutine
0
1
2
3
4
5
6
7
8
9
10
函数结束
--- PASS: TestGoroutineStatement (0.00s)
测试goroutine 0
测试goroutine 1
测试goroutine 2
测试goroutine 3
测试goroutine 4
测试goroutine 5
测试goroutine 6
测试goroutine 7
测试goroutine 8
测试goroutine 9
测试goroutine 10
PASS
ok  	gonotes/goroutine	0.285s
*/
