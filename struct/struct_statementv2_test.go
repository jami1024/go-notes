package struct_test

import (
	"fmt"
	"testing"
)

//定义结构体
type Personv2 struct {
	name string
	age  int
	sex  string
}

func TestStructStatementV2(t *testing.T) {
	p := newPerson("wagner", 18, "男")
	fmt.Println(p.name, p.age, p.sex)
}

//使⽤函数来实例化结构体
func newPerson(name string, age int, sex string) *Personv2 {
	return &Personv2{
		name: name,
		age:  age,
		sex:  sex}
}

/*
=== RUN   TestStructStatementV2
wagner 18 男
--- PASS: TestStructStatementV2 (0.00s)
PASS
ok  	gonotes/struct	(cached)
*/
