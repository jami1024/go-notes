package method

import (
	"fmt"
	"testing"
)

//测试⽅法.需要参数类型为Kongfu接⼝类型的参数
func testInterface(k KongFu) {
	k.Toad()
	k.SixSwords()
}

func TestPolymorphism(t *testing.T) {
	fmt.Println("利用接口实现多态")
	h := Haojiahuo{name: "好家伙"}
	fmt.Println(h.name)
	l := Laolitou{name: "⽼李头"}
	fmt.Println(l.name)
	testInterface(h) // 好家伙 实现了蛤蟆功.. 好家伙 实现了六脉神剑.
	testInterface(l) // ⽼李头 也实现了蛤蟆功.. ⽼李头 也实现了六脉神剑.
	kf := h
	kf.Toad()
	kf1 := l
	kf1.SixSwords()

}

// 定义接口
type KongFu interface {
	Toad()      //蛤蟆功
	SixSwords() //六脉神剑
}

// 定义类
type Haojiahuo struct {
	name string
}

type Laolitou struct {
	name string
}

//实现⽅法
func (o Haojiahuo) Toad() {
	fmt.Println(o.name, "实现了蛤蟆功..")
}

//实现⽅法
func (o Haojiahuo) SixSwords() {
	fmt.Println(o.name, "实现了六脉神剑..")
}

//实现⽅法
func (f Laolitou) Toad() {
	fmt.Println(f.name, "也实现了蛤蟆功..")
}

//实现⽅法
func (f Laolitou) SixSwords() {
	fmt.Println(f.name, "也实现了六脉神剑.")
}

//实现⾃⼰的⽅法
func (f Laolitou) PlayGame() {
	fmt.Println(f.name, "玩游戏..")
}
