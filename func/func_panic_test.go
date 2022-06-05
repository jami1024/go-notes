package fuc_test

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	fmt.Println("panic和recover")
	func_panic()
}

func func_panic() {

	defer func() {
		ms := recover()            //这⾥执⾏恢复操作
		fmt.Println(ms, "恢复执⾏了..") //恢复程序执⾏,且必须在defer函数中执⾏
	}()
	defer fmt.Println("第1个被defer执⾏")
	defer fmt.Println("第2个被defer执⾏")
	for i := 0; i <= 6; i++ {
		if i == 4 {
			panic("中断操作") //让程序进⼊恐慌 终端程序操作
		}

	}

	defer fmt.Println("第3个被defer执⾏") //恐慌之后的代码是不会被执⾏的

	/*
		=== RUN   TestPanic
		panic和recover
		第2个被defer执⾏
		第1个被defer执⾏
		中断操作 恢复执⾏了..
		--- PASS: TestPanic (0.00s)
		PASS
		ok  	gonotes/func	0.463s

	*/

}
