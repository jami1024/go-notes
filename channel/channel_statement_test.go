package channel_test

import (
    "fmt"
    "testing"
)

var channel1 chan int

func TestChanelStatement(t *testing.T) {
    fmt.Println("通道的定义")
    fmt.Printf("通道的数据类型:%T,通道的值:%v\n", channel1, channel1)
    if channel1 == nil {
        channel1 = make(chan int)
        fmt.Printf("通过make创建的通道数据类型:%T,通道的值:%v,\n", channel1, channel1)
    }
}
/*
=== RUN   TestChanelStatement
通道的定义
通道的数据类型:chan int,通道的值:<nil>
通过make创建的通道数据类型:chan int,通道的值:0xc0000244e0,
--- PASS: TestChanelStatement (0.00s)
PASS
ok  	gonotes/channel	0.394s
*/