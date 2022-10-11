package main

import "fmt"

func main() {
	//布尔类型定义
	var b1 bool = true
	var b2 bool = false
	var b3 = true
	var b4 = false
	b5 := true
	b6 := false

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
	fmt.Println(b5)
	fmt.Println(b6)

	age := 18

	if age >= 18 {
		fmt.Println("你已成年")
	} else {
		fmt.Println("你还未成年")
	}

}
