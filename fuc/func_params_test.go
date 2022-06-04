package fuc_test

import (
	"fmt"
	"testing"
)

func TestFuncParams(t *testing.T) {
	//函数的调⽤
	sum := getSum2(30)      //这⾥30 作为实参传递给⽅法
	fmt.Printf("%d\n", sum) //465

	//声明⼀个变量f
	var f func()

	//将⾃定义函数myfunc 赋给变量f
	f = myfunc
	//调⽤变量f 相当于调⽤函数myfunc()
	f()

}

//定义函数 num为形参⽤于接收调⽤⽅传递过来的参数
func getSum2(num int) int {
	sum := 0
	for i := 0; i <= num; i++ {
		sum += i
	}
	return sum
}

//⾃定义函数
func myfunc() {
	fmt.Println("myfunc...")
}
