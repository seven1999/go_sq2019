package fetch_sale_orderlist

import (
	swim "MyPractice/confunction"
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"strings"
)

/**
desciption:	每个接口与yaml参数对应的struct

*/

//定义配置文件解析后的结构
type yamlConfig struct {
	Testsuit    string      `yaml:"Testsuit"`
	Description string      `yaml:"description"`
	TestCases   []TestCases `yaml:"TestCases"`
}

type TestCases struct {
	CaseName string `json:"caseName"`
	Input    Input  `json:"input"`
	Output   Output `json:"output"`
	Param    Param  `json:"param"`
}

type Input struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

type Output struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

type Param struct {
	Pageno     int    `json:"pageno"`
	Count      int    `json:"count"`
	PlatformId int    `json:"platform_id"`
	LmNumber   string `json:"lm_number"`
	Status     int    `json:"status"`
	ShopId     string `json:"shop_id"`
	From       int    `json:"from"`
	To         int    `json:"to"`
	OrderId    string `json:"order_id"`
	Courier    string `json:"courier"`
	Mykey      Mykey  `json:"mykey"`
}

type Mykey struct {
	Key01 []Key01 `json:"key01"`
	Key02 string  `json:"key02"`
}

type Key01 struct {
	Vv int `json:"vv"`
	Kk int `json:"kk"`
}

// 直接获取yaml转成string(data) -- interface{}  可直接传入接口参数
func GetYamlData(filePath string, no int) interface{} {
	v := yamlConfig{}
	YamlParse := swim.NewYamlStruct()
	YamlParse.LoadYaml(filePath, &v)

	// 转成json格式
	data, _ := json.Marshal(v.TestCases[no]) // 获取testcase
	return string(data)

}

// ========================获取具体某一个入参param==============================
//yaml转map[string]interface{} ,在处理成interface{}, 同时处理随机参数替换
func GetYamlDataParam(filePath string, no int) interface{} {
	YamlParse := swim.NewYamlStruct()
	var mapResult map[string]interface{}
	mapResult = YamlParse.LoadYamlToMap(filePath)
	// map对应到struct
	yamlStruct := yamlConfig{}
	err := mapstructure.Decode(mapResult, &yamlStruct) // map转到对应的struct  保证参数与接口一致
	if err != nil {
		fmt.Println(err)
	}
	// 对应参数后struct再转回map
	mapResult = swim.StructToMap(yamlStruct)
	TestCases := mapResult["TestCases"].([]interface{})[no]
	//fmt.Println(TestCases)
	mapTestCases := TestCases.(map[string]interface{})["param"].(map[string]interface{})
	for k, v1 := range mapTestCases {
		switch v1.(type) {
		case string: // type是string才需要走进替换
			if strings.Contains(v1.(string), "g_no") {
				mapTestCases[k] = swim.GetVarInt(v1.(string))
			} else {
				mapTestCases[k] = swim.GetVarString(v1.(string))
			}
		}
	}
	// mapTestCases 转string
	mjson, _ := json.Marshal(mapTestCases)
	reqData := string(mjson)
	//fmt.Println(reqData)
	return reqData
}

// ========================获取具整个yaml--结果是Map==============================
func GetYamlDataAll(filePath string) map[string]interface{} {
	YamlParse := swim.NewYamlStruct()
	var mapResult map[string]interface{}
	mapResult = YamlParse.LoadYamlToMap(filePath)
	// map对应到struct
	yamlStruct := yamlConfig{}
	err := mapstructure.Decode(mapResult, &yamlStruct) // map转到对应的struct  保证参数与接口一致
	if err != nil {
		fmt.Println(err)
	}
	// 对应参数后struct再转回map
	mapResult = swim.StructToMap(yamlStruct)
	return mapResult
}

// ========================获取[]TestCases-==============================
func GetYamlDataTestCases(filePath string) []interface{} {
	YamlParse := swim.NewYamlStruct()
	var mapResult map[string]interface{}
	mapResult = YamlParse.LoadYamlToMap(filePath)
	// map对应到struct
	yamlStruct := yamlConfig{}
	err := mapstructure.Decode(mapResult, &yamlStruct) // map转到对应的struct  保证参数与接口一致
	if err != nil {
		fmt.Println(err)
	}
	// 对应参数后struct再转回map
	mapResult = swim.StructToMap(yamlStruct)
	TestCases := mapResult["TestCases"].([]interface{})
	//mapTestCases := TestCases.(map[string]interface{})["Param"].(map[string]interface{})

	return TestCases
}
