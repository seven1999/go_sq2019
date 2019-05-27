
package main

import (
	"fmt"
	"github.com/buger/jsonparser"
)

func main() {
	data := []byte(`{
					  "person": {
						"name":{
						  "first": "Leonid",
						  "last": "Bugaev",
						  "fullName": "Leonid Bugaev"
						},
						"github": {
						  "handle": "buger",
						  "followers": 109
						},
						"avatars": [
						  { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
						  { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=480", "type": "thumbnail" }
						]
					  },
					  "company": {
						"name": "Acme"
					  }
					}`)

	result, err := jsonparser.GetString(data, "person", "name", "fullName")
	result01, _ := jsonparser.GetString(data, "person", "avatars", "[1]", "url")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(result01)

	content, valueType, offset, err := jsonparser.Get(data, "person", "name", "fullName")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content, valueType, offset)
	//jsonparser提供了解析bool、string、float64以及int64类型的方法，至于其他类型，我们可以通过valueType类型来自己进行转化
	result1, err := jsonparser.ParseString(content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result1)

	err = jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		fmt.Printf("key:%s\n value:%s\n Type:%s\n", string(key), string(value), dataType)
		return nil
	}, "person", "name")

}