package fuc_test

import (
	"fmt"
	"testing"
)

func TestFuncClosure(t *testing.T) {
	fmt.Println("闭包")
	res := closure()
	r1 := res()               //执⾏closure函数返回的匿名函数
	fmt.Printf("r1:%v\n", r1) //1
	r2 := res()
	//普通的函数应该返回1，⽽这⾥存在闭包结构所以返回2 。
	fmt.Printf("r2:%v\n", r2) //2
	//⼀个外层函数当中有内层函数，这个内层函数会操作外层函数的局部变量,
	//并且外层函数把内层函数作为返回值,则这⾥内层函数和外层函数的局部变量,统称为闭包结构
	//所以上⾯打印的r2 是累计到2 。

	res2 := closure() //再次调⽤则产⽣新的闭包结构 局部变量则新定义的
	r3 := res2()
	fmt.Printf("r3:%v\n", r3) //1
}

//定义⼀个闭包结构的函数 返回⼀个匿名函数
func closure() func() int {
	//外层函数
	//定义局部变量a
	a := 0 //外层函数的局部变量 //定义内层匿名函数 并直接返回
	return func() int {
		//内层函数
		a++ //在匿名函数中将变量⾃增。内层函数⽤到了外层函数的局部变量，此变量不会随着外层函数的结束销毁
		return a
	}
}
