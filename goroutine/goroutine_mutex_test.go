package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
)

//定义全局变量 表示救济粮⻝总量
var food2 = 10

//同步等到组对象
var wg2 sync.WaitGroup

//创建⼀把锁
var mutex sync.Mutex

func TestMutex(t *testing.T) {
	fmt.Println("互斥锁")
	wg2.Add(4)
	// 开启4个携程抢粮食
	go Relief("A")
	go Relief("B")
	go Relief("C")
	go Relief("D")
	// 阻塞主协程，等待⼦协程执⾏结束
	wg2.Wait()
}

//定义⼀个发放的⽅法
func Relief(name string) {
	defer wg2.Done()
	for {
		//上锁
		mutex.Lock()
		if food2 > 0 {
			//加锁控制之后每次只允许⼀个协程进来，就会避免争抢
			food2--
			// time.Sleep(1 * time.Second)
			fmt.Println(name, "抢到救济粮 ，还剩下", food2, "个")
		} else {
			mutex.Unlock()
			//条件不满⾜也需要解锁 否则就会造成死锁其他不能执⾏
			fmt.Println(name, "别抢了 没有粮⻝了。")
			break
		}
		//执⾏结束解锁，让其他协程也能够进来执⾏
		mutex.Unlock()
	}
}

/*

=== RUN   TestMutex
互斥锁
D 抢到救济粮 ，还剩下 9 个
D 抢到救济粮 ，还剩下 8 个
D 抢到救济粮 ，还剩下 7 个
D 抢到救济粮 ，还剩下 6 个
D 抢到救济粮 ，还剩下 5 个
D 抢到救济粮 ，还剩下 4 个
D 抢到救济粮 ，还剩下 3 个
D 抢到救济粮 ，还剩下 2 个
D 抢到救济粮 ，还剩下 1 个
D 抢到救济粮 ，还剩下 0 个
D 别抢了 没有粮⻝了。
A 别抢了 没有粮⻝了。
B 别抢了 没有粮⻝了。
C 别抢了 没有粮⻝了。
--- PASS: TestMutex (0.00s)
PASS
ok  	gonotes/goroutine	(cached)
*/
