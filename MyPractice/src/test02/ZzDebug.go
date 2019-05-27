package main

import (
	"MyPractice/confunction"
	"MyPractice/confunction/fetch_sale_orderlist"
	"fmt"
)

var (
	yamlPath = "/Users/qisong/local/go_project/entrytask01/src/comfunc/fetch_sale_orderlist/fetch_sale_orderlist.yml"
)

func main() {
	// 设置需要替换的参数							// 嵌套的后续处理
	input := map[string]interface{}{
		"pageno":           20,
		"mykey.key01.1.vv": 1000,
		"mykey.key02":      "GGGG0008",
	}
	// 获取yaml数据
	yamlDataAll := fetch_sale_orderlist.GetYamlDataAll(yamlPath)
	// 获取入参数据
	inputJson := confunction.ChangeParam(yamlDataAll, input, "param")
	fmt.Println(inputJson)

	map03 := make(map[string]interface{})
	map03["aaa"] = "AAA"
	in := make([]interface{}, 3)
	in[0] = 1
	in[1] = 2
	in[2] = map03
	map02 := make(map[string]interface{})
	map02["pageno01"] = in

	tag := "pageno.pageno01.0"
	tag01 := "pageno.pageno01.2.aaa"
	map01 := make(map[string]interface{})
	map01["pageno"] = map02

	//fmt.Println(map01["pageno"].(map[string]interface{})["pageno01"].([]interface{})[0])
	fmt.Println(map01["pageno"].(map[string]interface{})["pageno01"].([]interface{})[2])
	fmt.Println(confunction.GetNestMapData(map01, tag))
	fmt.Println(confunction.GetNestMapData(map01, tag01))

}
