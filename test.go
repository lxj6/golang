package main

import "fmt"

func main() {
	//var定义一个变量 后面跟着数据类型
	var test string
	test = "hello"
	fmt.Println(test)

	//var 批量定义类型
	var (
		a int
		b string
		c float64
	)
	a = 3
	b = "test"
	c = 5.5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//var 批量定义并赋值
	var (
		d int    = 4
		e string = "haha"
	)

	fmt.Println(d)
	fmt.Println(e)

	//类型推断 :=

	f := "f"
	g := 3.33
	h := 10

	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

}
