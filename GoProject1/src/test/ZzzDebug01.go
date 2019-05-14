// debug
package main

import (
	"cal"
	"fmt"
)

func main() {
	fmt.Println("=================")

	var nums = []int{1, 2, 100, 999, 777}
	var mymap map[string]string
	mymap = make(map[string]string)
	mymap["who"] = "i"
	mymap["how"] = "love"
	mymap["what"] = "China"

	cal.RangeTest(nums)
	cal.RangeTest02(mymap)

	cal.MyMap()

	fmt.Println(cal.Factorial(5))

	//接口demo
	var Iphone1 cal.Iphone
	Iphone1.Message = "你好，我是iphone"
	name := cal.Sumsung("你好，我是sumsung")
	var phone cal.Phone
	var phone01 cal.Phone
	phone = Iphone1
	phone01 = name
	fmt.Println(phone.Receive())
	fmt.Println(phone.Getsumsung())
	phone.Call()
	fmt.Println(phone01.Receive())
	fmt.Println(phone01.Getsumsung())
	phone01.Call()

}
