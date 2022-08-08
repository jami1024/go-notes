package channel_test

import (
    "fmt"
    "testing"
)

func TestChannelUse(t *testing.T) {
    fmt.Println("通道的使用")
    ch1 := make(chan int)

    go func() {

        fmt.Println("======⼦协程执⾏======")
        data := <-ch1 //从通道中读取数据
        fmt.Println("读取到通道中的数据是:", data)
    }()

    ch1 <- 10 //往通道⾥放数据

    fmt.Println("======主协程结束======")
}

/*
=== RUN   TestChannelUse
通道的使用
======⼦协程执⾏======
读取到通道中的数据是: 10
======主协程结束======
--- PASS: TestChannelUse (0.00s)
PASS
ok  	gonotes/channel	0.286s
*/
