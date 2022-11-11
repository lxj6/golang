package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (per Person) eat() {
	fmt.Printf("%v eat..\n", per.name)
}
func (per Person) sleep() {
	fmt.Printf("%v sleep..\n", per.name)
}

func (per Person) work() {
	fmt.Printf("%v work..\n", per.name)
}

//以new的方式实例化struct,并初始化属性
func newPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name 不得为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("age 不得小于0")
	}
	return &Person{name, age}, nil
}

func main() {
	tom := Person{"tom", 19}
	tom.eat()
	tom.sleep()
	tom.work()

	per, err := newPerson("jerry", 1)
	if err != nil {
		fmt.Println(err)
	} else {
		per.work()
		per.sleep()
		per.eat()
	}

}
