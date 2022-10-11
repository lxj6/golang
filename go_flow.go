package main

import "fmt"

func f1() {
	f := 10
	if f > 5 {
		fmt.Println("大于5")
	} else {
		fmt.Println("小于5")
	}
}

func f2() {
	day := 22
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("未知")
	}
}

func f3() {
	score := 10
	switch score {
	case 1:
		fmt.Println(1)
	case 10: //fallthrough 向下穿透，不加则自动视为break
		fallthrough
	default:
		fmt.Println(score)
	}
}

func main() {
	f1()
	f2()
	f3()
	a := 100
	if a > 20 {
		fmt.Println("a大于20")
	} else {
		fmt.Println("a小于20")
	}

	b := 30
	switch b {
	case 10:
		fmt.Println("b等于10")
	case 30:
		fmt.Println("b等于30")
	default:
		fmt.Println("不等于任何")
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	arr := [...]int{1, 2, 3, 4}
	for key, value := range arr {
		fmt.Printf("%v:%v\r\n", key, value)
	}

	// golang中不可以使用0或非0标示真假
	/*d := 10
	if d {
		fmt.Println(d)
	}*/

	//golang可以把定义写在if条件中，但仅仅可以在if作用域中中调用age
	if age := 20; age > 18 {
		fmt.Println("你已成年")
	} else {
		fmt.Println("你未成年")
	}
	//fmt.Println(age) 会报错
	point := 30
	if point >= 90 {
		fmt.Println("你已达到90")
	} else if point >= 60 && point <= 80 {
		fmt.Println("你已及格")
	} else {
		fmt.Println("你未及格")
	}
}
