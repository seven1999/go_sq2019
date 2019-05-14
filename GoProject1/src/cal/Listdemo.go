// ListDemo
package cal

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	mylist := list.New() //通过 container/list 包的 New 方法初始化 list

	var mylist01 list.List //通过声明初始化list

	element := mylist.PushBack("U")
	mylist.PushFront("I")
	mylist.InsertBefore("Love", element)
	//	mylist.Remove(element)

	mylist01.PushBack("KKKK")

	//	fmt.Println(mylist)

	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, "\t")
	}
	fmt.Println()

}