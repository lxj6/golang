package main

import "fmt"

func main() {
	type person struct {
		id    int
		name  string
		age   int
		email string
	}
	tom := person{1, "tom", 18, "tom@email.com"}
	fmt.Println(tom)

}
