package main

import "fmt"

//定义接口后必须把接口内所有方法实现
type person interface {
	eat()
	sleep()
}

type tom struct {
	name string
}

type pet struct {
	name string
}

func (tom tom) eat() {
	fmt.Printf("name:%v\n", tom.name)
	fmt.Println("eat...")
}

func (tom tom) sleep() {
	fmt.Printf("name:%v\n", tom.name)
	fmt.Println("sleep...")
}

func (tom tom) player() {
	fmt.Printf("name:%v\n", tom.name)
	fmt.Println("player...")
}

func (p pet) eat() {
	fmt.Printf("name:%v\n", p.name)
	fmt.Println("eat...")
}

func (p pet) sleep() {
	fmt.Printf("name:%v\n", p.name)
	fmt.Println("sleep....")
}

func main() {
	t := tom{"tom"}
	t.eat()
	t.sleep()
	t.player()

	p := pet{"jerry"}
	p.eat()
	p.sleep()
}
