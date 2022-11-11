package main

import "fmt"

type kt struct {
	name string
}

func showkt1(tom kt) {

	tom.name = "james"
}

func showkt2(tom *kt) {
	//tom.name 自动解引用
	//(*tom).name = "tom...."
	tom.name = "kty"
}

func (kt kt) showkt3(tom kt) {
	tom.name = "james"
}
func (kt kt) showkt4(tom *kt) {
	tom.name = "kty"
}

func main() {

	p1 := kt{"tom"}
	p2 := &kt{"tom"}
	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
	fmt.Println("------------------------")
	/*	fmt.Printf("p1调用前:%v\n", p1)
		showkt1(p1)
		fmt.Printf("p1调用后:%v\n", p1)
		fmt.Println("------------------------")
		fmt.Printf("p2调用前:%v\n", *p2)
		showkt2(p2)
		fmt.Printf("p2调用后:%v\n", *p2)*/
	fmt.Printf("p1调用前:%v\n", p1)
	p1.showkt3(p1)
	fmt.Printf("p1调用后:%v\n", p1)
	fmt.Println("------------------------")
	fmt.Printf("p2调用前:%v\n", *p2)
	p1.showkt4(p2)
	fmt.Printf("p2调用后:%v\n", *p2)
}
