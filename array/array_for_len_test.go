package array

import (
	"fmt"
	"testing"
)

func TestArrayForLen(t *testing.T) {
	fmt.Println("利用for len方式进行对数组的遍历")
	var names [3]string = [3]string{"zhangsan1", "zhangsan2", "zhangsan3"}
	for i := 0; i < len(names); i++ {
		fmt.Printf("索引：%d, 内容：%s\n", i, names[i])
		fmt.Printf("索引：%d, 内容：%q\n", i, names[i])
		// %s 和 %q的区别：如果是字符串，%q会把值用双引号扩起来
	}
	/*
		=== RUN   TestArrayForLen
		利用for len方式进行对数组的遍历
		索引：0, 内容：zhangsan1
		索引：0, 内容："zhangsan1"
		索引：1, 内容：zhangsan2
		索引：1, 内容："zhangsan2"
		索引：2, 内容：zhangsan3
		索引：2, 内容："zhangsan3"
		--- PASS: TestArrayForLen (0.00s)
		PASS
		ok  	gonotes/array	0.475s
	*/
}
