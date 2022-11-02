package main

import "fmt"

func test4() {
	//数组定义
	var a1 [2]int
	var a2 [3]string
	fmt.Printf("a1:%T\r\n", a1)
	fmt.Printf("a2:%T\r\n", a2)
	fmt.Printf("a2:%v\r\n", a1) //如果没有定义值 int默认为0 string默认为空格
	fmt.Printf("a2:%v\r\n", a2)
	a1[0] = 1
	a1[1] = 2
	fmt.Printf("a2:%v\r\n", a1)
}

func test5() {
	var a3 = [3]int{1, 2} //初始化赋值不够数组长度默认后面为0
	fmt.Printf("a3:%v\r\n", a3)
	var a4 = [...]bool{true, false} //[...]省略数组的长度
	fmt.Printf("a4:%v\r\n", a4)
}

func test6() {
	var a5 = [...]int{1: 1, 3: 3, 5: 5} //指定数组的索引以后未指定的0,2,4则默认为0  如果是bool类型则未指定的是false
	for i, i2 := range a5 {
		fmt.Printf("i:%v,i2:%v\r\n", i, i2)
	}
}

func test7() {
	var a1 = [3]int{2, 3, 4}
	fmt.Println(len(a1))
}

func main() {
	//test4()
	//test5()
	test6()
	test7()
}
