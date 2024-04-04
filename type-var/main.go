package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d"}
	for k, v := range a {
		fmt.Printf("%v, %v\n", k, v)
	}
}
