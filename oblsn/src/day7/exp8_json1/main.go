package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"` //json tag can be called like alias
	NickName string `json:"nickname"`
	Age      int
	Email    string
	Birthday string
}

func testStruct() {
	user1 := &User{
		UserName: "Lee",
		NickName: "ming",
		Age:      16,
		Email:    "111@126.com",
		Birthday: "1994/7/15",
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func testInt() {
	var age = 20
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	fmt.Printf("%s\n", data)
}

func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "111"
	m["age"] = 20
	m["sex"] = "male"
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}
func testSlice() {
	var m map[string]interface{}
	var s []map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "111"
	m["age"] = 20
	m["sex"] = "male"
	s = append(s, m)
	s = append(s, m) //twice
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func main() {
	testStruct()
	testInt()
	testMap()
	testSlice()
}
