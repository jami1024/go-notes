package struct_test

import (
	"fmt"
	"testing"
)

//定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func TestStructStatement(t *testing.T) {
	fmt.Println("结构体声明")
	//实例化后并使⽤结构体
	p := Person{} //使⽤简短声明⽅式，后⾯加上{}代表这是结构体
	p.age = 18    //给结构体内成员变量赋值
	p.address = "北京"
	p.name = "张三"
	p.sex = "男"
	fmt.Println(p.age, p.address, p.name, p.sex) //使⽤点.来访问结构体内成员的变量的值。

	/*
		=== RUN   TestStructStatement
		结构体声明
		18 北京 张三 男
		--- PASS: TestStructStatement (0.00s)
		PASS
		ok  	gonotes/struct	0.453s
	*/
	//直接给成员变量赋值
	p2 := Person{age: 18, address: "北京", name: "lisi", sex: "⼥"}
	fmt.Println(p2.age, p2.address, p2.name, p2.sex)

	/*
		=== RUN   TestStructStatement
		结构体声明
		18 北京 张三 男
		18 北京 lisi ⼥
		--- PASS: TestStructStatement (0.00s)
		PASS
		ok  	gonotes/struct	0.434s
	*/

}
