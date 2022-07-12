package goroutine_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGoroutineResource(t *testing.T) {
	fmt.Println("临界资源安全")
	//开启4个协程抢粮⻝
	go Fafang("灾⺠1")
	go Fafang("灾⺠2")
	go Fafang("灾⺠3")
	go Fafang("灾⺠4")

	//让程序休息5秒等待所有⼦协程执⾏完毕
	time.Sleep(5 * time.Second)
}

//定义全局变量 表示救济粮⻝总量
var food = 10

//定义⼀个发放的⽅法
func Fafang(name string) {
	for {
		if food > 0 {
			//此时有可能第⼆个goroutine访问的时候 第⼀个goroutine还未执⾏完 所以条件也成⽴
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			//随机休眠时间
			food--
			fmt.Println(name, "抢到救济粮 ，还剩下", food, "个")
		} else {
			fmt.Println(name, "别抢了 没有粮⻝了。")
			break
		}
	}
}

/*
=== RUN   TestGoroutineResource
临界资源安全
灾⺠1 抢到救济粮 ，还剩下 9 个
灾⺠4 抢到救济粮 ，还剩下 8 个
灾⺠1 抢到救济粮 ，还剩下 7 个
灾⺠4 抢到救济粮 ，还剩下 6 个
灾⺠1 抢到救济粮 ，还剩下 5 个
灾⺠3 抢到救济粮 ，还剩下 4 个
灾⺠2 抢到救济粮 ，还剩下 3 个
灾⺠4 抢到救济粮 ，还剩下 2 个
灾⺠1 抢到救济粮 ，还剩下 1 个
灾⺠3 抢到救济粮 ，还剩下 0 个
灾⺠3 别抢了 没有粮⻝了。
灾⺠1 抢到救济粮 ，还剩下 -1 个
灾⺠1 别抢了 没有粮⻝了。
灾⺠4 抢到救济粮 ，还剩下 -2 个
灾⺠4 别抢了 没有粮⻝了。
灾⺠2 抢到救济粮 ，还剩下 -3 个
灾⺠2 别抢了 没有粮⻝了。
--- PASS: TestGoroutineResource (5.00s)
PASS
ok  	gonotes/goroutine	5.468s

*/
