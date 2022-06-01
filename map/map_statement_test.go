package map_test

import (
	"fmt"
	"testing"
)

func TestMapStatement(t *testing.T) {
	fmt.Println("map的声明")
	// 1 var 变量名称 map[key的数据类型]value的数据类型 默认值是nil
	var m1 map[string]int
	// 2 使⽤make声明 make(map[key_data_type]value_data_type)
	var m2 = make(map[string]int)
	//3,直接声明并初始化赋值map⽅法
	m3 := map[string]int{"张三": 89, "李四": 23, "王二": 90} // map[] map[] map[张三:89 李四:23 王二:90]
	fmt.Println(m1, m2, m3)
	fmt.Println(m1 == nil) //true
	fmt.Println(m2 == nil) //false
	fmt.Println(m3 == nil) //false
}
