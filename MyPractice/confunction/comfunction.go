package confunction

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unicode"

	//"gopkg.in/yaml.v2"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

/**
===================读取json==================
**/
type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//读取的数据为json格式，需要进行解码	json转struct
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

// struct转成map
func StructToMap(jsonStr interface{}) (mapResult map[string]interface{}) {
	marshalContent, err := json.Marshal(jsonStr)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(marshalContent, &mapResult)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("==========")
	//fmt.Println(mapResult)
	return mapResult
}

// map转成string(json)
func MapToString(inputMap map[string]interface{}) string {
	mjson, _ := json.Marshal(inputMap)
	StringData := string(mjson)
	return StringData
}

/**
===================读取yaml==================
**/

type YamlStruct struct {
}

func NewYamlStruct() *YamlStruct {
	return &YamlStruct{}
}

func (jst *YamlStruct) LoadYaml(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//读取的数据为json格式，需要进行解码	json转struct
	err = yaml.Unmarshal(data, v)
	if err != nil {
		return
	}
}

// 读取yaml 转成map （为后续处理使用）
func (jst *YamlStruct) LoadYamlToMap(filename string) (mapResult map[string]interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	j2, err := yaml.YAMLToJSON(data) //yaml转json
	// json转成map
	err01 := json.Unmarshal([]byte(j2), &mapResult)
	if err01 != nil {
		fmt.Println("JsonToMapDemo err: ", err01)
	}
	return mapResult
}

// 获取随机数 （手机号，英文名称， 英文+数字， 邮箱等）
func GetRandomParam(width int, prefix string, endfix string) string {

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return prefix + sb.String() + endfix
}

//// 处理获取的随机数	$g_no("10", "t", "@mail")
//func GetVar(sval string) (int, string) {
//
//	if strings.HasPrefix(sval, "$") {
//		sval = strings.Replace(strings.Trim(strings.Trim(sval, "$g_no("), ")"), " ", "", -1)
//		width, prefix, endfix := strings.Split(sval, ",")[0], strings.Split(sval, ",")[1], strings.Split(sval, ",")[2]
//		width01, _ := strconv.Atoi(width)
//		sval = GetRandomMobile(width01, prefix, endfix)
//	}
//	intSval, _ := strconv.Atoi(sval)
//	return intSval, sval
//}

// 处理获取的随机数-int	$g_no(10, 1)
func GetVarInt(sval string) int {

	if strings.HasPrefix(sval, "$") {
		sval = strings.Replace(strings.Trim(strings.Trim(sval, "$g_no("), ")"), " ", "", -1)
		width, prefix, endfix := strings.Split(sval, ",")[0], strings.Split(sval, ",")[1], strings.Split(sval, ",")[2]
		width01, _ := strconv.Atoi(width)
		sval = GetRandomParam(width01, prefix, endfix)
	}
	intSval, _ := strconv.Atoi(sval) // string 转成int
	return intSval
}

// 处理获取的随机数-string	$g_no("0, t, @mail)
func GetVarString(sval string) string {

	if strings.HasPrefix(sval, "$") {
		sval = strings.Replace(strings.Trim(strings.Trim(sval, "$g_ranstr("), ")"), " ", "", -1)
		width, prefix, endfix := strings.Split(sval, ",")[0], strings.Split(sval, ",")[1], strings.Split(sval, ",")[2]
		width01, _ := strconv.Atoi(width)
		sval = GetRandomParam(width01, prefix, endfix)
	}
	return sval
}

// 封装GetVarInt和GetVarString
func GetVar(sval string) (interface{}, interface{}) {
	rs := []rune(sval)
	if strings.HasPrefix(sval, "$") {
		//sval = strings.Replace(strings.Trim(strings.Trim(sval, "$g_ranstr("), ")"), " ", "", -1)
		sval = strings.Replace(string(rs[strings.Index(sval, "(")+1:len(rs)-1]), " ", "", -1)
		width, prefix, endfix := strings.Split(sval, ",")[0], strings.Split(sval, ",")[1], strings.Split(sval, ",")[2]
		width01, _ := strconv.Atoi(width)
		sval = GetRandomParam(width01, prefix, endfix)
	}
	intSval, _ := strconv.Atoi(sval) // string 转成int
	return sval, intSval
}

// 参数替换
func ChangeParam(mapYaml map[string]interface{}, data interface{}, tag string) interface{} {
	mapData := make(map[string]interface{})
	if data != "" {
		// data转成map
		mapData = data.(map[string]interface{})
	}
	mapResult := make(map[string]interface{})
	// 确定替换的部分
	if tag == "param" || tag == "input" || tag == "caseName" || tag == "output" {
		TestCases := mapYaml["TestCases"].([]interface{})[0]
		fmt.Println(TestCases)
		mapResult = TestCases.(map[string]interface{})[tag].(map[string]interface{})
	} else {
		mapResult = mapYaml[tag].(map[string]interface{})
	}
	for k, v := range mapData {
		switch v.(type) {
		case string:
			v = GetVarImpl(v.(string)) // 处理随机数替换
		}

		if !(strings.Contains(k, ".") || strings.Contains(k, "..")) {
			mapResult[k] = v
		} else {
			var sliceTemp []interface{} // 定义临时空切片
			mapTemp := mapResult        // 临时map					// 疑惑：:= 赋值以后mapTemp内部的值变了，mapResult也跟着变了。。不需要重新赋值？
			//fmt.Println(&mapResult == &mapTemp)
			ss := strings.Split(k, ".")
			for _, obj := range ss {
				if !IsNumber(obj) {
					switch mapTemp[obj].(type) {
					case map[string]interface{}:
						mapTemp = mapTemp[obj].(map[string]interface{})
					case []interface{}:
						sliceTemp = mapTemp[obj].([]interface{})
					case string:
						mapTemp[obj] = v
					case int:
						mapTemp[obj] = v
					case float64:
						mapTemp[obj] = v
					case bool:
						mapTemp[obj] = v
					}
				} else {
					objInt, _ := strconv.Atoi(obj)
					mapTemp = sliceTemp[objInt].(map[string]interface{})
				}
				//fmt.Println(mapTemp)
			}
		}
	}
	// map转成string
	//fmt.Println(mapResult)
	mjson, _ := json.Marshal(mapResult)
	reqData := string(mjson)
	return reqData
}

// 替换随机参数方法 处理map
func ChangeRandowParam(tagMap map[string]interface{}) map[string]interface{} {

	for k, v := range tagMap {
		switch v.(type) {
		case string: // 随机替换的格式一定是string的
			if strings.Contains(v.(string), "g_no") {
				tagMap[k] = GetVarInt(v.(string))
			} else {
				tagMap[k] = GetVarString(v.(string))
			}
		}
	}
	return tagMap
}

func GetVarImpl(savl string) interface{} {
	var resultString interface{}
	resultString = savl
	sval01, sval02 := GetVar(savl)
	if strings.Contains(savl, "g_no") {
		resultString = sval02
	} else if strings.Contains(savl, "g_str") {
		resultString = sval01
	}
	return resultString
}

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
			//fmt.Println(mapTemp)
		}
	}
	return mapYaml[tag]
}
