// 声明包
package main

// 引入包
import (
	"fmt"
)

// 函数 (main 函数为整个程序的入口)
func main() {
	a := 922337203685477580

	fmt.Printf("c: %T \r\n", a)
	// 语句表达式
	fmt.Println("hello")
	// golang 共有成员与私有成员通过成员的标识的首字母区分
	// 首字母大写表示共有成员，首字母小写表示私有成员，只有两种情况
	//p1.F1()
}

// 初始化函数， golang每个包的引用会优先调用这个函数
func init() {
	//fmt.Println("world")
}
