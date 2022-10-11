package main

import "fmt"

func main() {
	a := 100
	b := 20
	x := a + b
	fmt.Println(x)
	x = a - b
	fmt.Println(x)
	x = a * b
	fmt.Println(x)
	x = a / b
	fmt.Println(x)
	x = a % b
	fmt.Println(x)

	c := 10
	c++
	fmt.Println(c)
	c--
	fmt.Println(c)

	d := 100
	d += 1
	fmt.Println(d)
	d -= 1
	fmt.Println(d)
	d *= 2
	fmt.Println(d)
	d /= 10
	fmt.Println(d)
}
