package main

import "fmt"

func f5() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func f6() {
	//for初始条件也可以放在外部，自增也可以放在内部手动来break
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func f7() {
	i := 0
	for i < 10 {
		if i > 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func main() {

	//golang中只有for循环，没有while和do while
	//f5()
	//f6()
	//f7()

	/*for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}*/

	/*for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if j == 1 {
				continue
			}
			fmt.Println(j)
		}
		fmt.Println(i)
	}*/
	/**
	golang 中可以使用 for range遍历数组、切片、字符串、map及通道(channel),通过遍历的返回值有以下规律
	1.数组、切片、字符串返回 索引和值
	2.map返回键和值
	3.通道只返回通道内的值
	*/
	//数组
	arr := [5]int{1, 2, 3, 4, 5}
	for k, v := range arr {
		fmt.Printf("key:%v,val:%v\r\n", k, v)
	}
	//不需要k的话可以用_代替，go中定义的变量必须使用
	for _, v := range arr {
		fmt.Println(v)
	}

	//切片
	arr1 := []int{1, 2, 3, 4}
	for i, i2 := range arr1 {
		fmt.Printf("key:%v,val:%v\r\n", i, i2)
	}

	//map  相当于其他语言的键值对
	m := make(map[string]string, 0)
	m["name"] = "kenny"
	m["age"] = "20"
	m["email"] = "kenny@email.com"

	for s, s2 := range m {
		fmt.Printf("%v:%v\r\n", s, s2)
	}

	//string
	name := "kenny"
	for i, i2 := range name {
		fmt.Printf("%v:%c\r\n", i, i2)
	}
	// 定义MYLABLE标签，指定break跳出到标签位置
MYLABLE:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break MYLABLE
		}
	}
	fmt.Println("end...")
}
