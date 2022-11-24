package httptest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get (url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("에러발생", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("에러발생", err)
		panic(err)
	}
	fmt.Println("성공")

	var myInfo Info
	err = json.Unmarshal(data, &myInfo)
	if err != nil {
		fmt.Println("json parsing error", err)
	}
	fmt.Println(myInfo)

	// fmt.Printf("%s\n", string(data))
}

type Info struct {
	Name string `json:"name"`
	Age int `json:"age"`
}