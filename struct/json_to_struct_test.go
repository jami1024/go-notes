package struct_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person3 struct {
	Name string `json:"name"` //指定json字段为⼩写输出
}

func TestJsonToStruct(t *testing.T) {
	fmt.Println("json字符串转结构体")
	jsonstr := `{"name":"王二"}`
	var p Person3
	if err := json.Unmarshal([]byte(jsonstr), &p); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("结构体:%+v\n", p) // %+v 将结构体的字段名称打印出来
	// 结构体:{Name:王二}

}
