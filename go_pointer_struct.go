package main

import "fmt"

func test1(dodo person) {
	dodo.id = 1
	dodo.name = "dodos"

}

func test2(kite *person) {
	kite.id = 2
	kite.name = "kites"
}

type person struct {
	id    int
	name  string
	age   int
	email string
}

func main() {

	//var ps *person 结构体指针
	tom := person{1, "tom", 18, "tom@email.com"}
	ps := &tom
	fmt.Println(tom)
	fmt.Printf("v:%p\n", ps)
	fmt.Println(*ps)

	//test1 属于拷贝副本，函数体内修改的只是副本，外部没有被修改
	dodo := person{
		id:   3,
		name: "dodo",
	}
	fmt.Println(dodo)
	test1(dodo)
	fmt.Println("-----")
	fmt.Println(dodo)

	//test2通过指针传递，函数内部修改了之后外部也被修改了
	kite := person{
		id:   4,
		name: "kite",
	}
	fmt.Println(kite)
	fmt.Println("----------")
	test2(&kite)
	fmt.Println(kite)
}
