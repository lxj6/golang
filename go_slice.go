package main

import "fmt"

//切片的低层也是引用了数组，但是数组是固定的长度，切片是可变长度，可使用append对slice进行填充
//切片属于动态数组，可以对其进行curd
func test1() {
	var s1 []int
	var s2 []string
	fmt.Println(s1)
	fmt.Println(s2)
}
func test2() {
	//make方式定义 打印结果为[0 0] 因为已经分配了2个内存 还没有赋值
	var s3 = make([]int, 2)
	fmt.Printf("s3:%v", s3)
}

func test3() {
	var s1 = []int{1, 2, 3}
	fmt.Printf("len(s1):%v\n", len(s1)) //打印切片长度
	fmt.Printf("cap(s1):%v\n", cap(s1)) //打印切片容量
	fmt.Println(s1[0])
}

func test4() {
	s4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s4[0:5]) //截取切片从0到5
	fmt.Println(s4[5:])  //截取切片从5到末尾
	fmt.Println(s4[:4])  //从0截取到4
}

func test5() {
	//for range 遍历切片
	s5 := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range s5 {
		fmt.Printf("k:%v -- v:%v\n", k, v)
	}
}

func test6() {
	//for循环遍历切片
	s5 := []string{"a", "b", "c", "d", "e", "f"}
	for i := 0; i < len(s5); i++ {
		fmt.Printf("v:%v\n", s5[i])
	}
}

func test7() {
	//切片的添加
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300, 400)
	fmt.Printf("s1:%v", s1)
}

func test8() {
	//切片的删除
	s2 := []int{1, 2, 3, 4}
	s2 = append(s2[0:2], s2[3:]...) //通过拼接方式删除不需要的数据
	fmt.Println(s2)
}

func test9() {
	//切片的修改，直接通过索引进行修改
	s2 := []int{1, 2, 3, 4}
	s2[0] = 100
	fmt.Println(s2)
}

func test10() {
	//切片的查询
	s2 := []int{1, 2, 3, 4}
	key := 2
	for k, _ := range s2 {
		if k == key {
			fmt.Println(s2[k])
		}
	}
}

func test11() {
	//两个切片使用同一块内存，有一个修改，则另一个也会修改,如果不想使用一个内存地址可以使用copy函数
	s1 := []int{1, 2, 3}
	s2 := s1
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}

func test12() {
	//make一个s2切片并分配内存地址，然后使用copy函数把s1复制到s2，这样操作修改s2的值 s1不会改变
	s1 := []int{1, 2, 3}
	var s2 = make([]int, 3)
	copy(s2, s1)
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}

func main() {
	test12()
}
