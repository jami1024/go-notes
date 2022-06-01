package map_test

import (
	"fmt"
	"testing"
)

func TestMapOperate(t *testing.T) {
	fmt.Println("map的使用")
	var m1 map[int]string         //只是声明 nil
	var m2 = make(map[int]string) //创建
	m3 := map[string]int{"张三": 89, "李四": 23, "王二": 90}

	fmt.Println(m1, m2, m3) // map[] map[] map[张三:89 李四:23 王二:90]

	//map 为nil的时候不能使⽤ 所以m1操作之前要初始化值
	m1 = make(map[int]string)

	//1存储键值对到map中 语法:map[key]=value
	m1[1] = "狗"
	m1[2] = "猫"
	m1[0] = "鸡"

	//2获取map中的键值对 语法:map[key]
	val := m1[2]
	fmt.Println(val) // 猫

	//3判断key是否存在 语法：value,ok:=map[key]
	val, ok := m1[1]
	fmt.Println(val, ok) //结果返回两个值，⼀个是当前获取的key对应的val值。⼆是当前值否存在，会返回⼀个true或false。

	//4修改map如果不存在则添加， 如果存在直接修改原有数据。
	m1[1] = "⼩狗狗"
	fmt.Println(m1) //map[0:鸡 1:⼩狗狗 2:猫]

	//5删除map中key对应的键值对数据 语法: delete(map, key)
	delete(m1, 1)
	fmt.Println(m1) //map[0:鸡 2:猫]

	//6 获取map中的总⻓度 len(map)
	fmt.Println(len(m1)) //2

	//7 map和切片的配合使用
	s1 := make([]map[string]string, 0, 2)
	map1 := make(map[string]string)
	map1["name"] = "zhangsan"
	map1["age"] = "18"

	map2 := make(map[string]string)
	map2["name"] = "lisi"
	map2["age"] = "20"

	s1 = append(s1, map1, map2)

	fmt.Println(s1) // [map[age:18 name:zhangsan] map[age:20 name:lisi]]

	for key, val := range s1 {
		fmt.Println(key, val)
		// 0 map[age:18 name:zhangsan]
		// 1 map[age:20 name:lisi]
	}
}
