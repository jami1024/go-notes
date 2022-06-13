package struct_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

//结构体
type Person1 struct {
	Name string
	Addr *Addr1
}

// 结构体2
type Person2 struct {
	Name string `json:"name"`
	Addr *Addr1 `json:"addr,omitempty"`
}

//地址结构体
type Addr1 struct {
	Addr string
}

func TestStructToJson(t *testing.T) {
	fmt.Println("结构体转json")
	p := Person1{}
	p2 := Person2{}
	p.Name = "张三"
	p2.Name = "李四"
	a := Addr1{}
	a.Addr = "beijing"
	p.Addr = &a
	buf, err := json.Marshal(p) //转换为json返回两个结果
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("json = ", string(buf))
	/*
		结果 json =  {"Name":"张三","Addr":{"Addr":"beijing"}}
		从结果可以看出其中json字符中每⼀个key的⾸字母也是⼤写，
		如果没有设置数据的字段的结果为null。
		如何强制将他变为⼩写的。并且将不需要显⽰的字段隐藏掉。就需要在结构体上添加标记。
	*/

	buf2, err := json.Marshal(p2) //转换为json返回两个结果
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("json = ", string(buf2))
	/*
		结果：json =  {"name":"李四"}
	*/
}
