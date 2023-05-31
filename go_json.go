package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	p := person{
		Name:  "李晓景",
		Age:   18,
		Email: "lxj@gmail.com",
	}
	// 结构体转json
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	jsonByte := []byte("{\"name\":\"lxj\",\"age\":18,\"email\":\"lxj@gmail.com\"}")
	json.Unmarshal(jsonByte, &p)
	fmt.Println(p)
}
