package main

import "fmt"

//golang执行顺序 初始化变量->init()->main()
var i = initaVar()

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2..")
}

func initaVar() int {
	fmt.Println("initVar...")
	return 100
}

func main() {
	fmt.Println("main")
}
