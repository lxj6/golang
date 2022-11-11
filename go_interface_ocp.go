package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}
type Dog struct {
}
type Cat struct {
}

//todo 我们可以继续添加宠物pig，只要实现了eat和sleep就可以通过person.care调用，实现了ocp思想

//Dog 实现 Pet接口
func (dog Dog) eat() {
	fmt.Println("dog eat...")
}
func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}

//Cat 实现 Pet接口
func (cat Cat) eat() {
	fmt.Println("cat eat...")
}
func (cat Cat) sleep() {
	fmt.Println("cat sleep...")
}

type Person struct {
}

//因为dog和cat都实现了pet接口，所以这里传dog和cat都可以
func (per Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{}
	cat := Cat{}
	person := Person{}
	person.care(dog)
	person.care(cat)
}
