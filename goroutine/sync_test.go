package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
)

//创建⼀个同步等待组的对象
var wg sync.WaitGroup

func TestSync(t *testing.T) {
	fmt.Println("sync包测试")
	wg.Add(3) //设置同步等待组的数量
	go Relief1()
	go Relief2()
	go Relief3()
	wg.Wait() //主goroutine进⼊阻塞状态
	fmt.Println("main end...")

}

func Relief1() {
	fmt.Println("func1...")
	wg.Done() //执⾏完成 同步等待数量减1
}
func Relief2() {
	defer wg.Done()
	fmt.Println("func2...")
}
func Relief3() {
	defer wg.Done() //推荐使⽤延时执⾏的⽅法来减去执⾏组的数量
	fmt.Println("func3...")
}

/*
=== RUN   TestSync
sync包测试
func3...
func2...
func1...
main end...
--- PASS: TestSync (0.00s)
PASS
ok  	gonotes/goroutine	(cached)
*/
