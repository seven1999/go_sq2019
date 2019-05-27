package main

import (
	"MyPractice/confunction"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func IsNumber(str string) bool {
	r := []rune(str)
	for _, obj := range r {
		if !unicode.IsDigit(obj) {
			return false
		}
	}
	return true
}

// 参数替换   a.0.b
func GetNestMapData(mapYaml map[string]interface{}, tag string) interface{} {
	if strings.Contains(tag, ".") || strings.Contains(tag, "..") {
		var sliceTemp []interface{} // 定义临时空切片
		ss := strings.Split(tag, ".")
		for num, obj := range ss {
			if !IsNumber(obj) {
				switch mapYaml[obj].(type) {
				case map[string]interface{}:
					if num == len(ss) - 1{
						return mapYaml[obj]
					}else {
						mapYaml = mapYaml[obj].(map[string]interface{})
					}
				case []interface{}:
					if num == len(ss) - 1{
						return mapYaml[obj]
					}else {
						sliceTemp = mapYaml[obj].([]interface{})
					}
				case string:
					return mapYaml[obj]
				case int:
					return mapYaml[obj]
				case float64:
					return mapYaml[obj]
				case bool: //有需要再自己添加
					return mapYaml[obj]
				}
			} else {
				objInt, _ := strconv.Atoi(obj)
				switch sliceTemp[objInt].(type) {
				case map[string]interface{}:
					if num == len(ss) - 1{
						return sliceTemp[objInt]
					}else {
						mapYaml = sliceTemp[objInt].(map[string]interface{})
					}
				case string:
					return sliceTemp[objInt]
				case int:
					return sliceTemp[objInt]
				case float64:
					return sliceTemp[objInt]
				case bool: //有需要再自己添加
					return sliceTemp[objInt]
				}
			}
		}
	}
	return mapYaml[tag]
}


func main() {

	map03 := make(map[string]interface{})
	map03["aaa"] = "AAA"
	in := make([]interface{}, 3)
	in[0] = 1
	in[1] = 2
	in[2] = map03
	map02 := make(map[string]interface{})
	map02["pageno01"] = in

	tag := "pageno.pageno01"
	tag01 := "pageno.pageno01.0"
	tag02 := "pageno.pageno01.2.aaa"
	map01 := make(map[string]interface{})
	map01["pageno"] = map02

	//fmt.Println(map01["pageno"].(map[string]interface{})["pageno01"].([]interface{})[0])
	//fmt.Println(map01["pageno"].(map[string]interface{})["pageno01"].([]interface{})[2])
	fmt.Println(confunction.GetNestMapData(map01, tag))
	fmt.Println(confunction.GetNestMapData(map01, tag01))
	fmt.Println(confunction.GetNestMapData(map01, tag02))

}
