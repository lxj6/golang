package main

import "fmt"

func test1() {
	i := 2
	if i > 1 {
		goto LABRY
		fmt.Println("i大于1")
	} else {
		goto LABRY
	}
LABRY:
	fmt.Println("end....")
}

func test2() {
MYLABRY:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				break MYLABRY
			}
			fmt.Printf("%v:%v\r\n", i, j)
		}
	}

}

func main() {
	//test1()
	test2()
}
