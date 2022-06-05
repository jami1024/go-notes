package fuc_test

import (
	"fmt"
	"testing"
)

func TestFuncDefer(t *testing.T) {
	defer deferTest(1) //第⼀个被defer的，函数后执⾏
	defer deferTest(2) //第⼆个被defer的，函数先执⾏
	deferTest(3)       //没有defer的函数，第⼀次执⾏
}

func TestFuncDefer2(t *testing.T) {
	a := 10
	defer deferTest(a) //此时a已经作为10 传递出去了 只是没有执⾏
	a++                //a++ 不会影响defer函数延迟执⾏结果
	deferTest(a)
}

func deferTest(s int) {
	fmt.Println(s)
}
