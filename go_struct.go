package main

import "fmt"

//结构体定义，多个同类型的字段可以放在一行
type person struct {
	id, age     int
	name, email string
}

//匿名结构体定义
var dog struct {
	name string
	age  int
}

func main() {
	/*var tom person
	fmt.Println(tom) //未初始化的字段默认为0(int)或空(string)或者false(bool),成员访问通过tom.id来访问
	tom.id = 101
	tom.age = 18
	tom.name = "tom"
	tom.email = "tom@email.com"
	fmt.Println(tom)
	dog.name = "小黑"
	dog.age = 3
	fmt.Println(dog)*/
	//值初始化
	tom := person{
		id:    101,
		age:   18,
		name:  "tom",
		email: "tom@email.com",
	}
	fmt.Println(tom)
	//也可以直接顺序定义 注意：顺序需要和定义的顺序一样
	tom = person{102, 19, "toms", "toms@email.com"}
	fmt.Println(tom)
	//部分初始化
	tom = person{
		id:    103,
		name:  "tomss",
		email: "tomss@email.com",
	}
	fmt.Println(tom)
}
