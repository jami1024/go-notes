package method

import (
	"fmt"
	"testing"
)

func TestInherit(t *testing.T) {
	fmt.Println("方法的继承")
	o := &Ouyangcrazy{
		Name: "欧阳疯",
		Age:  70,
	} //创建⽗类
	o.ToadKongfu() //⽗类对象访问⽗类⽅法

	y := &YangGuo{
		Ouyangcrazy{
			Name: "杨过",
			Age:  18,
		},
		"古墓",
	} //创建⼦类

	fmt.Println(y.Name)    //⼦类对象访问⽗类中有的字段
	fmt.Println(y.Address) //⼦类访问⾃⼰的字段

	y.ToadKongfu() //⼦类对象访问⽗类⽅法
	y.NewKongfu()  //⼦类访问⾃⼰的⽅法
	y.ToadKongfu() //如果存在⾃⼰的⽅法 访问⾃⼰重写的⽅法

	/*
		=== RUN   TestInherit
		方法的继承
		欧阳疯 的蛤蟆功！
		杨过
		古墓
		杨过 的新蛤蟆功！
		杨过 ⼦类⾃⼰的新功夫！
		杨过 的新蛤蟆功！
		--- PASS: TestInherit (0.00s)
		PASS
		ok  	gonotes/method	0.108s
	*/
}

// 创建⼀个结构体起名 Ouyangcrazy 代表⽗类
type Ouyangcrazy struct {
	Name    string
	Age     int
	Ability string
}

//创建⼀个结构体代表⼦类
type YangGuo struct {
	Ouyangcrazy        //包含⽗类所有属性
	Address     string //单独⼦类有的字段
}

// ⽗类的⽅法
func (o *Ouyangcrazy) ToadKongfu() {
	fmt.Println(o.Name, "的蛤蟆功！")
}

//⼦类的⽅法
func (y *YangGuo) NewKongfu() {
	fmt.Println(y.Name, "⼦类⾃⼰的新功夫！")
}

//⼦类重写⽗类的⽅法
func (y *YangGuo) ToadKongfu() {
	fmt.Println(y.Name, "的新蛤蟆功！")
}
