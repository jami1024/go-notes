package method

import (
	"fmt"
	"testing"
)

func TestMethodStatement(t *testing.T) {
	fmt.Println("结构体类型调⽤⽅法")
	//使⽤User结构体创建⼀个⻆⾊代表张三

	zhangsan := User{"张三", 10}

	zhangsan.jineng() //调⽤这个结构体的⽅法。
}

//创建⼀个结构体代表⼈物⻆⾊--张三
type User struct {
	Name  string // 姓名
	Level int    //级别
}

//创建⼀个⽅法,只要是User结构体就能调⽤。

func (u User) jineng() {
	kill := 200.0
	ability := "泰拳"
	fmt.Printf("我是:%s，我的等级:%d，我的武功:%s，杀伤⼒%.1f\n", u.Name, u.Level, ability, kill)
}
