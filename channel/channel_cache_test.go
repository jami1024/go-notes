package channel_test

import (
	"fmt"
	"testing"
)

func TestChannelCache(t *testing.T) {
	fmt.Println("缓冲通道测试案例")
	//定义⼀个缓冲区⼤⼩为5的通道
	ch3 := make(chan int, 5)
	//开启⼦协程写⼊数据
	go func() {
		for i := 0; i < 10; i++ {
			ch3 <- i
			fmt.Println("⼦协程写⼊数据：", i)
		}
		close(ch3) //关闭通道
	}()
	//主协程读取数据
	for {
		v, ok := <-ch3
		if !ok {
			fmt.Println("读取结束", ok)
			break
		}
		fmt.Println("主协程读取到的数据为：", v)
	}
	fmt.Println("主协程结束")
}
/*
=== RUN   TestChannelCache
缓冲通道测试案例
⼦协程写⼊数据： 0
⼦协程写⼊数据： 1
⼦协程写⼊数据： 2
⼦协程写⼊数据： 3
⼦协程写⼊数据： 4
⼦协程写⼊数据： 5
主协程读取到的数据为： 0
主协程读取到的数据为： 1
主协程读取到的数据为： 2
主协程读取到的数据为： 3
主协程读取到的数据为： 4
主协程读取到的数据为： 5
主协程读取到的数据为： 6
⼦协程写⼊数据： 6
⼦协程写⼊数据： 7
⼦协程写⼊数据： 8
⼦协程写⼊数据： 9
主协程读取到的数据为： 7
主协程读取到的数据为： 8
主协程读取到的数据为： 9
读取结束 false
主协程结束
--- PASS: TestChannelCache (0.00s)
PASS
ok  	gonotes/channel	(cached)
*/
