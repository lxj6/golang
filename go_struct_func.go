package main

import (
	"fmt"
)

type Ren struct {
	name string
	age  int
}

//结构体和属性分开，可以在不同的文件，但是必须在一个包下
func (per Ren) eat() {
	fmt.Printf("%v,eat...\n", per.name)
}

func (per Ren) sleep() {
	fmt.Printf("%v,sleep....\n", per.name)
}

type Customer struct {
	name string
}

func (cus Customer) login(name string, pwd string) bool {
	fmt.Printf("action name:%v\n", cus.name)
	if name == "tom" && pwd == "123" {
		return true
	}
	return false
}

func main() {
	//类似面向对象的类
	james := Ren{"james", 19}
	james.eat()
	james.sleep()

	tom := Customer{"tom"}
	bl := tom.login("tom", "123")
	if bl {
		fmt.Println("登陆成功!")
	} else {
		fmt.Println("登录失败!")
	}
}
