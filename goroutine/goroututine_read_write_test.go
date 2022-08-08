package goroutine_test

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

//创建⼀把读写锁 可以是他的指针类型
var rwmutex sync.RWMutex

var wg3 sync.WaitGroup

func TestReadWrite(t *testing.T) {
    fmt.Println("读写锁")
    wg3.Add(3)
    go Read(1)
    go Write(2)
    go Read(3)
    wg3.Wait()
    fmt.Println("======main结束======")
}

func Read(i int) {
    defer wg3.Done()
    fmt.Println("准备读取数据")
    // 读上锁
    rwmutex.RLock()
    fmt.Println("正在读取..", i)
    rwmutex.RUnlock() //读取操作解锁
    fmt.Println("======读取结束======")

}

func Write(i int) {
    defer wg3.Done()
    fmt.Println("======开始读写数据======")
    rwmutex.Lock() //写操作上锁
    fmt.Println("======正写数据...", i)
    time.Sleep(3 * time.Second)
    rwmutex.Unlock() //写操作解锁
    fmt.Println("======写操作结束======")
}

/*
=== RUN   TestReadWrite
读写锁
准备读取数据
正在读取.. 3
======读取结束======
准备读取数据
正在读取.. 1
======读取结束======
======开始读写数据======
======正写数据... 2
======写操作结束======
======main结束======
--- PASS: TestReadWrite (3.00s)
PASS
ok  	gonotes/goroutine	3.331s

*/
