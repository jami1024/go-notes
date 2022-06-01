package point

import (
	"fmt"
	"testing"
)

// int float string 数组 都是属于值类型

func TestPointStatement(t *testing.T) {
	fmt.Println("指针的定义")
	// 定义指针,初始值为nil

	var username string = "zhangsan"
	var name *string = &username // 取出username的地址赋值给name =》取引用

	// *name 取出name的内存地址中对应的存储的值 =》解引用
	fmt.Printf("name的类型%T\tname的值%v\n", name, *name)
}
