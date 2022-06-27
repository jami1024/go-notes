package method

import (
	"fmt"
	"testing"
)

func TestStructPackage(t *testing.T) {
	fmt.Println("使⽤结构体来实现封装")
	//创建⼀个对象
	h := NewPerson("雷锋")
	h.SetAge(18)                    //访问封装的⽅法设置年龄
	fmt.Println(h.Name, h.GetAge()) //使⽤对象封装的⽅法获取年龄

}

// 定义结构体实现封装
type GoodPerson struct {
	Name string
	Age  int
}

//使⽤NewPerson⽅法创建⼀个对象
func NewPerson(name string) *GoodPerson {
	return &GoodPerson{Name: name}
}

// 使⽤SetAge⽅法设置结构体成员的Age
func (h *GoodPerson) SetAge(age int) {
	h.Age = age
}

// 使⽤GetAge⽅法获取成员现在的Age
func (h *GoodPerson) GetAge() int {
	return h.Age
}
