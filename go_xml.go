package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
}

func main() {
	p := person{
		Name:  "lxj",
		Age:   18,
		Email: "lxj@gmail.com",
	}

	// 结构体转xml
	date, err := xml.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(date))

	strByte := []byte(`<name>lxj</name><age>18</age><email>lxj@gmail.com</email>`)
	// xml转结构体
	xml.Unmarshal(strByte, &p)
	fmt.Println(p)
}
