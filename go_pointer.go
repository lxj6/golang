package main

import "fmt"

func main() {
	var ip *int
	fmt.Printf("ip:%V\n", ip)
	fmt.Printf("type:%T\n", ip)

	var i = 100
	ip = &i
	fmt.Printf("ip:%v\n", ip)

	var sp *string
	var s = "hello"
	sp = &s
	fmt.Printf("sp:%T\n", sp)
	fmt.Printf("sp:%v\n", sp)
}
