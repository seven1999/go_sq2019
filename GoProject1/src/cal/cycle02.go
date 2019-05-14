// cycle02
package cal

import (
	"fmt"
)

func GetNum(num int, tag string) int {

	if num == 0 {
		num = num + 10000
		if tag == "Y" {
			num *= 2
		} else {
			num = 0
		}
	} else if num > 0 {
		num += 9999
	} else {
		num -= 9999
	}

	return num
}

func SwithTest(marks int) {

	/* 定义局部变量 */
	var grade string = "B"

	switch {
	case (marks > 90):
		grade = "A"
	case marks > 80:
		grade = "B"
	case marks > 70:
		grade = "C"
	case marks >= 60:
		grade = "D"
	case marks < 60:
		grade = "E"
	}
	fmt.Printf("grade is %s\n", grade)
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B":
		fmt.Printf("良好\n")
	case grade == "C":
		fmt.Printf("一般\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "E":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
}

func ForTest(name []string) {

	var count int = 0
	for a := 0; a < len(name); a++ {
		for name[a] != "" {
			count += 1
			fmt.Printf("name is %s\n", name[a])
			if count == 2 {
				count = 0
				break
			}
		}
	}

}

func main() {
	fmt.Println(GetNum(0, "Y"))
	fmt.Println(GetNum(0, "N"))
	fmt.Println(GetNum(1, "Y"))
	fmt.Println(GetNum(1, "N"))

	var names = []string{"", "Lily", "Hanna"}
	ForTest(names)
	SwithTest(89)

}
