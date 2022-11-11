package main

import "fmt"

type animal struct {
}

func (an animal) parent() {
	fmt.Println("this is parent class")
}

type dog struct {
	animal
}
type cat struct {
	animal
}

func (c cat) parent() {
	fmt.Println("this is son class")
}

func main() {
	dog := dog{}
	//如果dog没有实现父类的parent方法，可以直接使用dog.parent()来调用
	dog.parent()
	cat := cat{}
	//如果cat实现了父类的parent方法，那么cat.parent()调用自己本身关联的方法，cat.animal.parent()则会调用父类的方法
	cat.parent()
}
