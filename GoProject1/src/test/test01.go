package main

import (
	"cal"
	"fmt"
)



func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c  = 1, false, "str"	// 多重复值

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d\n", area)
	//println()
	println(a, b, c)

	fmt.Println(cal.Add(100,300))
	fmt.Println("Hello World!")

	fmt.Println(cal.GetArray(10))


	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(cal.MyMulty(numbers))
	fmt.Println(cal.MyMulty(cal.GetArray(10)))


}


