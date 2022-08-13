package channel_test

import (
	"fmt"
	"testing"
)

func TestReadWrite(t *testing.T) {
	fmt.Println("读写通道")

	// 双向通道可读可写
	ch4 := make(chan int)
	// 单向通道只能写不能读
	ch5 := make(chan<- int)
	// 单向通道只能读不能写
	ch6 := make(<-chan int)

	//如果创建时候创建的就是双向通道
	//则在⼦协程内部写⼊数据，读取的时候不受影响。
	go WriteOnly(ch4)
	data1 := <-ch4
	fmt.Println("获取到只写通道中的数据是", data1)

	//如果将定向通道ch5只写通道，作为参数传递。
	//则不能读取到写回来的数据。
	go WriteOnly(ch5)
	// data2 := <-ch5 //不能读取会报错：invalid operation: <-ch5 (receive from send-only type chan<- int)

	go ReadOnly(ch6)

}

//只读
func ReadOnly(ch <-chan int) {
	data := <-ch
	fmt.Println("读取到通道的数据是：", data)
}

//只写
func WriteOnly(ch chan<- int) {
	//如果传进来的原本是双向通道
	//但是函数本身接收的是⼀个只写的通道，则在此函数内部只允许写⼊数据不允许读取数据
	//所以单向通道往往是作为参数传递
	ch <- 10
	fmt.Println("只写通道结束")
}
