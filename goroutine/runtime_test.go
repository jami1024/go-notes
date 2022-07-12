package goroutine_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestRuntime(t *testing.T) {
	//获取当前GOROOT⽬录
	fmt.Println("GOROOT:", runtime.GOROOT())
	//获取当前操作系统
	fmt.Println("操作系统:", runtime.GOOS)
	//获取当前逻辑CPU数量
	fmt.Println("逻辑CPU数量：", runtime.NumCPU())
	//设置最⼤的可同时使⽤的CPU核数 取逻辑cpu数量
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n) //⼀般在使⽤之前就将cpu数量设置好 所以最好放在init函数内执⾏
	//goexit 终⽌当前goroutine //创建⼀个goroutine
	go func() {
		fmt.Println("start...")
		defer func() {
			fmt.Println("defer end")
		}()
		runtime.Goexit()      //终⽌当前goroutine
		fmt.Println("end...") //不会执行
	}()
	time.Sleep(1 * time.Second) //主goroutine休眠1秒让⼦goroutine执⾏完
	fmt.Println("main_end...")
}

/*
=== RUN   TestRuntime
GOROOT: /usr/local/go
操作系统: darwin
逻辑CPU数量： 8
8
start...
defer end
main_end...
--- PASS: TestRuntime (1.00s)
PASS
ok  	gonotes/goroutine	1.285s
*/
