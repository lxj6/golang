package main

import "fmt"

func main() {
	//变量使用const  const constName [type] = value  常量定义以后不可修改
	const P float64 = 3.14
	const PI = 3.14
	fmt.Println(P)
	fmt.Println(PI)
	//定义多个常量
	const (
		w = 100
		h = 200
	)
	fmt.Println(w)
	fmt.Println(h)
	//多重赋值
	const i, j = 1, 2
	const a, b, c = 3, 1.1, "c"
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//iota 可以被认为是一个可被编译器修改的常量，默认值为0,每次调用加1,遇到const关键字时被重置为0
	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	// _可以跳过自增
	const (
		b1 = iota
		_
		b3 = iota
	)
	fmt.Println(b1)
	fmt.Println(b3)

	//也可以直接赋值跳过
	const (
		c1 = iota
		c2 = 100
		c3 = iota
	)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

}
