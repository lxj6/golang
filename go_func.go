package main

import "fmt"

//有返回值函数
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func sum1(a int, b int) int {
	return a + b
}

//无参数和返回值函数
func say() {
	fmt.Println("hello world")
}

//无参数返回两个值 sstring,int
func getNamgeAndAge() (string, int) {
	n := "tom"
	a := 18
	return n, a
}

//多个参数args,传过来的是一个slice数组
func test2(args ...int) {
	for k, v := range args {
		fmt.Printf("k:%v,v:%v\n", k, v)
	}
}

//
func test3(s [3]int) {
	s[0] = 100
}

//切片属于指针类型，即使修改的是拷贝副本，底层指向的也会被修改
func test4(s []int) {
	s[1] = 200
}

//定义两个相同类型函数，可以赋值给fun
type fun func(int, int) int

func sums(a int, b int) int {
	return a + b
}

func subs(a int, b int) int {
	return a - b
}

func main() {
	//s1 := sum(1, 3)
	//fmt.Println(s1)
	/*name, age := getNamgeAndAge()
	fmt.Printf("name:%v\n", name)
	fmt.Printf("age:%v\n", age)
	age := 17
	age1 := 18
	test2(age, age1)
	s := [3]int{1, 2, 3}
	test3(s)
	test4(s)
	fmt.Println(s)
	var f fun
	f = sums
	s := f(1, 2)
	fmt.Println(s)
	f = subs
	s = f(5, 1)
	fmt.Println(s)*/
}
