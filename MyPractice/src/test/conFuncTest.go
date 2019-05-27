package main

import (
	"MyPractice/confunction"
	"MyPractice/confunction/fetch_sale_orderlist"
	"fmt"
)

/*
此为示例用例文件，可在此文件内完成用例，如有需要，可以添加文件
自己拉一个分支提交测试代码，不允许提交到test分支

*/

var (

	yamlPath = "E:\\Go_project\\MyPractice\\confunction\\fetch_sale_orderlist\\fetch_sale_orderlist.yml"

)

// 查询订单列表--成功
func main() {
	// 设置需要替换的参数
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
}

