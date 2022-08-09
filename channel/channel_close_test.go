package channel_test

import (
	"fmt"
	"testing"
)

func TestChannelClose(t *testing.T) {
	fmt.Println("这是通道的关闭测试案例")
	//创建⼀个通道⽤来传递数据
	ch2 := make(chan int)
	//通过⼦协程往通道中放数据
	go func() {
		fmt.Println("======放数据的⼦协程执⾏======")
		for i := 0; i < 10; i++ {
			ch2 <- i //往通道中放数据
		}
		close(ch2) //关闭通道
	}()
	//主协程通过for循环来获取通道中的所有数据
	for {
		v, ok := <-ch2 //获取通道的状态以及数据
		if !ok {
			fmt.Println("⼦协程已将通道关闭")
			break
		}
		fmt.Println("获取到的⼦协程数据为", v)
	}
	fmt.Println("主协程结束")
}

/*

=== RUN   TestChannelClose
这是通道的关闭测试案例
======放数据的⼦协程执⾏======
获取到的⼦协程数据为 0
获取到的⼦协程数据为 1
获取到的⼦协程数据为 2
获取到的⼦协程数据为 3
获取到的⼦协程数据为 4
获取到的⼦协程数据为 5
获取到的⼦协程数据为 6
获取到的⼦协程数据为 7
获取到的⼦协程数据为 8
获取到的⼦协程数据为 9
⼦协程已将通道关闭
主协程结束
--- PASS: TestChannelClose (0.00s)
PASS
ok  	gonotes/channel	(cached)

*/
