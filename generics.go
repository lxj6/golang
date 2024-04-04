package main

import "fmt"

func main() {
	var a, b = 1, 2
	var c, d = 1.5, 1.6

	fmt.Println("普通函数获取数字:", getMaxNumber(a, b))
	fmt.Println("普通函数获取浮点:", getMaxFloat(c, d))
	// 隐式约束，根据传入类型自动推断

	fmt.Println("泛型函数获取数字:", getMax(a, b))
	fmt.Println("泛型函数获取数字:", getMax(c, d))

	// 显示约束，根据传入的约束而约束
	fmt.Println("泛型函数获取数字:", getMax[int](a, b))

	var e, f uint8 = 1, 3
	fmt.Println("泛型函数获取数字:", getMaxs(e, f))
	var g, h MyInt64 = 5, 8
	// 衍生类型
	fmt.Println("泛型函数获取数字:", getMaxs(g, h))

	var i, j string = "abc", "abc"
	fmt.Println("comparable:", getBuiltInComparable(i, j))
}

// 下面俩方法的逻辑一样，但是类型不一样
func getMaxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 定义一个泛型函数 int|float代表只支持这两种类型，支持所有类型可使用any
func getMax[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

/*
由于有一些指针类型会被编译器认为T*int 所以可以写在interface中
type st [T*int | *float64]struct {
}

type st2[T interface{ *int | *float64 }] struct {
}
*/
// 也可以定义一个泛型接口集
type cusNumberT interface {
	// 支持unit8 int32 float64 与 int64及其衍生类型
	// ~表示支持类型的衍生类型
	// | 表示取并集
	uint8 | int32 | float64 | ~int64
	// 多行之间取交集
	//uint8 | int32 | float64 | ~int64
	//int32 | float64 | ~int64 | uint16
}

// MyInt64 为int64的衍生类型，是具有基础类型int64的新类型，与int64是不同类型
type MyInt64 int64

// MyInt32 为int32的别名，是统一类型  这就是 =和 不用=的区别
type MyInt32 = int32

func getMaxs[K cusNumberT](a, b K) K {
	if a > b {
		return a
	}
	return b
}

// comparable 类型，只支持 == != 两个操作比较的类型 如 int、float、string，以及指针、通道等，只要它们的值可以使用 == 或 != 运算符进行比较，它们就实现了 comparable 接口
func getBuiltInComparable[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

// 使用any相当于任意类型，和不约束是一样的
func printBuiltInAny[T any](a, b T) bool {
	return false //a == b
}
