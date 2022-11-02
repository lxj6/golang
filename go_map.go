package main

import "fmt"

func test1() {
	//map声明的两种方式
	m1 := make(map[int]string)
	m1[1] = "name"
	m1[2] = "age"
	fmt.Println(m1)

	m2 := map[int]string{}
	m2[1] = "age"
	m2[2] = "name"
	fmt.Println(m2)
	//声明并赋值
	var m3 = map[int]string{
		1: "a",
		2: "b",
	}
	fmt.Println(m3[1])
}

func test2() {
	var m1 = map[string]string{}
	m1["name"] = "tom"
	m1["age"] = "18"
	k1 := "name"
	k2 := "s"
	//检测 m1中是否含有k1,如果有返回value=tom  ok为true
	value, ok := m1[k1]
	fmt.Println(ok)
	fmt.Println(value)
	//检测 m1中是否含有k1,如果有返回value=  ok为false
	value, ok = m1[k2]
	fmt.Println(ok)
	fmt.Println(value)
}

func test3() {
	//map中的遍历
	m1 := map[string]string{
		"name": "tom",
		"age":  "18",
	}
	for k, v := range m1 {
		fmt.Printf("k:%v,v:%v\n", k, v)
	}
}

func main() {
	test3()
}
