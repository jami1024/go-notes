package struct_test

import (
	"fmt"
	"testing"
)

// 结构体1

type P1 struct {
	name string
	age  int
	addr Addr
}

// 结构体2
type Addr struct {
	addr string
}

// 也可以嵌套结构体指针
type P2 struct {
	name string
	age  int
	addr *Addr
}

func TestStructNested(t *testing.T) {
	fmt.Println("结构体嵌套")
	p := P1{}
	p.name = "章三"
	p.age = 18
	p.addr = Addr{addr: "beijing"}
	fmt.Println(p) //{章三 18 {beijing}}
	//结构体初始化可以使⽤上⾯两种格式将字段名和对应的值写在括号内，使⽤(字段名:值,)的格式填充
	//第⼆种初始化的⽅式，定义好结构体之后使⽤重新赋值的⽅式:使⽤(变量.字段名=值)的格式

	//嵌套结构体指针
	pr := P2{}
	pr.name = "lisi"
	pr.age = 22
	pre := Addr{}
	pre.addr = "chaoyang"
	pr.addr = &pre
	fmt.Println(pr) //{lisi 22 0xc000096520}
}
