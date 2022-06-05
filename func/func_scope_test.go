package fuc_test

import (
	"fmt"
	"testing"
)

func TestFuncScope(t *testing.T) {
	fmt.Println("变量在函数中的作用域")
	x := 2
	test()
	fmt.Printf("x的值:%d\n", x)
	if i := 2; i <= 30 {
		x := 20
		fmt.Println(x) //if语句内部定义的x,当没if内没有定义时就会往上找
	}
	// fmt.Println("index", i) //if内部定义的i if之外就不能访问了
}

//函数外部定义全局变量
//全局变量不⽀持简短定义⽅法 n:=0
var n = 3 //全局变量可以随意被修改 如果不想被修改可以定义为常量
func test() {
	// fmt.Print(x)   //undefined: x 未定义 x是main函数中定义的变量 所以不能访问
	fmt.Println(n) //n是全局变量 任何地⽅都可以访问
}
