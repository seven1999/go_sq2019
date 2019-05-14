// FuncTest.go
package cal

import (
	"fmt"
)

func GetMulpl(numArray []int, tagNum int) map[int]int {
	//	var x int
	//	var y int
	var g1 map[int]int
	g1 = make(map[int]int)
	for i := 0; i < len(numArray); i++ {
		for j := 0; j < len(numArray); j++ {
			if numArray[i]*numArray[j] == tagNum {
				g1[numArray[i]] = numArray[j]
			}
		}
	}
	fmt.Println("===========")
	return g1
}

//引用传递
func Swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

//闭包
func GetSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//func main() {
//	var numArray = []int{1, 2, 3, 4, 5, 6, 22, 8, 66, 33}
//	var tagNum = 66
//	fmt.Println(GetMulpl(numArray, tagNum))

//	/* 定义局部变量 */
//	var a int = 100
//	var b int = 200

//	fmt.Printf("交换前，a 的值 : %d\n", a)
//	fmt.Printf("交换前，b 的值 : %d\n", b)

//	/* 调用 swap() 函数
//	 * &a 指向 a 指针，a 变量的地址
//	 * &b 指向 b 指针，b 变量的地址
//	 */
//	Swap(&a, &b)

//	fmt.Printf("交换后，a 的值 : %d\n", a)
//	fmt.Printf("交换后，b 的值 : %d\n", b)

//	//闭包
//	/* nextNumber 为一个函数，函数 i 为 0 */
//	nextNumber := GetSeq()
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())
//}
