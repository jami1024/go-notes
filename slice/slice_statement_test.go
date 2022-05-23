package slice

import (
	"fmt"
	"testing"
)

func TestSliceStatement(t *testing.T) {
	fmt.Println("切片的声明")
	var names []string = []string{"a", "b", "c"}
	fmt.Printf("names类型:%T,names的值为:%q\n", names, names)

	var namesNil []string                // nil切片
	var namesBlank []string = []string{} //空切片
	fmt.Printf("namesNil:%t\nnamesBlank:%t\n", namesNil == nil, namesBlank == nil)

	// 利用make创建切片
	// make(type, len, cap)
	nameMake := make([]string, 3, 10)
	fmt.Printf("nameMake类型:%T\tnameMake的值为:%q\nnameMake的len:%d\tnameMake的cap:%d\n", nameMake, nameMake, len(nameMake), cap(nameMake))
	// 访问和修改
	fmt.Println("切片的访问")
	fmt.Printf("names第一个元素是:%v\n", names[0])
	fmt.Println("切片的修改")
	names[0] = "0"
	fmt.Printf("names:%q\n", names)
	// 长度,切片中已经存在的元素 len
	fmt.Printf("names的长度:%d\n", len(names))

	name3 := nameMake
	name3 = append(name3, "a")        // ["", "", "", "aa"]
	nameMake = append(nameMake, "aa") // ["", "", "", "aa"]
	// 变量赋值都是复制
	fmt.Printf("nameMake的值为:%q\nname3的值%q\n", nameMake, name3)
	// 0 <= start <= end <= cap
	// len = end - start
	// cap = cap - start

	// 0 <= start <= end <= cap_end <= cap
	// len = end -start
	// cap = cap_end = start
	fmt.Println(nameMake[1:3:5])

	/*
		=== RUN   TestSliceStatement
		切片的声明
		names类型:[]string,names的值为:["a" "b" "c"]
		namesNil:true
		namesBlank:false
		nameMake类型:[]string	nameMake的值为:["" "" ""]
		nameMake的len:3	nameMake的cap:10
		切片的访问
		names第一个元素是:a
		切片的修改
		names:["0" "b" "c"]
		names的长度:3
		nameMake的值为:["" "" "" "aa"]
		name3的值["" "" "" "aa"]
		[ ]
		--- PASS: TestSliceStatement (0.00s)
		PASS
		ok  	gonotes/slice	3.044s
	*/
}
