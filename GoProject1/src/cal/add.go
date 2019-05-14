package cal

import "strconv"

//函数名必须大写首字母，不然外部包找不到
func Add(a int,b int)(int){
	return a+b
}

func Sub(a int,b int)(int){
	return a-b
}

func GetArray(n int) [10]int{
	var a [10]int 		//int array with length 3
	for i := 1; i <= n; i++{
		a[i-1] = i
	}
	return a
}


func MyMulty(num [10]int)(map[string]string) {
	var countryCapitalMap map[string]string 		/*创建集合 */
	countryCapitalMap = make(map[string]string)
	for i := 0; i < len(num); i++{
		for j := 0; j < len(num); j++{
			if i*j == 40{
				countryCapitalMap[strconv.Itoa(i)] = strconv.Itoa(j)
			}
		}
	}
	return countryCapitalMap
}
