// StructDemo
package cal

import (
	"fmt"
)

type MyStruct struct {
	name   string
	age    int
	weight int
	height int
	favor  string
}

func PrintHumun(humun MyStruct) {
	fmt.Printf("MyStruct name : %s\n", humun.name)
	fmt.Printf("MyStruct age : %d\n", humun.age)
	fmt.Printf("MyStruct weight : %d\n", humun.weight)
	fmt.Printf("MyStruct height : %d\n", humun.height)
	fmt.Printf("MyStruct favor : %s\n", humun.favor)
}

// Range Demo
func RangeTest(nums []int) {
	for i, num := range nums {
		fmt.Printf("第%d个值:%d\n", i, num)
	}

	for _, num := range nums {
		fmt.Printf("列表值分为别：%d\n", num)
	}
}

// Range Demo02
func RangeTest02(mymap map[string]string) {
	for k, v := range mymap {
		fmt.Printf("%s --> %s\n", k, v)
	}

	for _, v := range mymap {
		fmt.Printf("map value分为别：%s\n", v)
	}
}

// map demo
func MyMap() {
	var maymap01 = make(map[string]string)
	maymap01["1"] = "one"
	maymap01["2"] = "two"
	maymap01["3"] = "three"

	for k, v := range maymap01 {
		fmt.Printf("%s -> %s\n", k, v)
	}

	mymap02 := map[int]string{1: "One", 2: "Two", 3: "Three"}
	delete(mymap02, 1) // delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key
	for k, v := range mymap02 {
		fmt.Printf("%d -> %s\n", k, v)
	}

}

// 递归
func Factorial(n uint64) (result uint64) {

	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
