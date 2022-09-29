package main

import "fmt"

//函数类型

func test() {

}

func main() {
	var name = "jars"
	age := 20
	bol := true
	flo := 3.14
	//fmt.printf %t 打印变量类型
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", bol)
	fmt.Printf("%T\n", flo)

	//b是一个指针的int类型
	a := 100
	b := &a
	fmt.Printf("%T\n", b)

	//数组类型 [...]代表不设长度
	arr := [2]int{1, 2}
	fmt.Printf("%T\n", arr)

	//切片类型 使用[]定义
	ar := []int{1, 2}
	fmt.Printf("%T\n", ar)

	fmt.Printf("%T\n", test)

}
