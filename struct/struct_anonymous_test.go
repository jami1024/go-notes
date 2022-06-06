package struct_test

import (
	"fmt"
	"testing"
)

func TestStructAnonymouns(t *testing.T) {
	fmt.Println("匿名结构体")
	p := struct {
		name string
	}{
		name: "zhangsan",
	}
	fmt.Println(p.name)
}
