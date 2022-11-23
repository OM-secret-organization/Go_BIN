package rpc

import (
	"encoding/json"
	"fmt"
)

type Str struct {
	Content string
}

type MyInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func Marshal(v interface{}) []byte {
	m, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return m
}

func Run  () {
	fmt.Println("hi rpc")
	encoded := Marshal("hi rpc")
	fmt.Println((encoded))

	// json marshal(JSON.stringify())
	info1 := MyInfo{"namename", 35}
	b, err := json.Marshal(info1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshal", string(b))

	// json unmarshal(JSON.parse())
	jsonStr := `{"name":"hanbin", "age":29}`
	var info2 MyInfo
	json.Unmarshal([]byte(jsonStr), &info2)
	fmt.Println("Unmarshal", info2)
	fmt.Printf("%+v\n", info2) // +를 붙이면 구조체의 필드까지 같이 표시됨.

	// 하... 자바스크립트에서는 JSON.parse랑 JSON.stringify만 있으면 뚝딱이었는데 진짜 개귀찮고 어렵다.... 슬프다.....
}