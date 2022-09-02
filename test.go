package main

import "fmt"

func getNameAndAge() (string, int) {
	return "namge", 30
}

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

	//var 批量定义并赋值 类型推断
	var (
		d = 4
		e = "haha"
	)

	fmt.Println(d)
	fmt.Println(e)

	var x, y, z = 1, "y", 10.1
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//短变量类型 := 只可在函数内部

	f := "f"
	g := 3.33
	h := 10

	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	//匿名变量 变量定义了以后必须使用，不使用可以使用_(匿名变量)
	_, age := getNameAndAge()

	fmt.Println(age)
}
