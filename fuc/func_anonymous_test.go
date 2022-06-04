package fuc_test

import (
	"fmt"
	"testing"
)

func TestAnonymousFunc(t *testing.T) {
	// 定义匿名函数
	func() {
		fmt.Println("匿名函数")
	}()

	func(a int, b int) {
		fmt.Println(a, b)
	}(1, 2) //括号内为匿名函数的实参

	res := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(res) //打印匿名函数返回值

	res2 := oper(20, 12, add)
	fmt.Printf("res2的值:%v\n", res2)

	//匿名函数作为回调函数直接写⼊参数中
	res3 := oper(2, 4, func(a, b int) int {
		return a + b
	})
	fmt.Printf("res3的值:%v\n", res3)
}

// 定义用于相加的函数
func add(a, b int) int {
	return a + b
}

//oper就叫做⾼阶函数
//fun 函数作为参数传递则fun在这⾥叫做回调函数
func oper(a, b int, fun func(int, int) int) int {
	// fmt.Println(a, b, fun) //20 12 0x49a810A 第三个打印的是传⼊的函数体内存地址
	fmt.Printf("第一个参数%v\t第二个参数%v\n", a, b)
	res := fun(a, b) //fun 在这⾥作为回调函数 程序执⾏到此之后才完成调⽤
	return res
}
