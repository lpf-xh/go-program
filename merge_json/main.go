package main

// 对多条json中的key进行去重合并，value不做特殊处理。

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

func main() {
	filename := flag.String("f", "", "存放json的文件")
	flag.Parse()

	// 读取文件
	b, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	root := map[string]interface{}{}

	// 逐行合并
	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}
		data := map[string]interface{}{}
		err := json.Unmarshal([]byte(line), &data)
		if err != nil {
			fmt.Printf("%d line: %s\n", i, line)
			fmt.Printf("json.Unmarshal error: %v\n", err)
			continue
		}
		merge(root, data)
	}

	// 输出结果
	b, err = json.Marshal(root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Merge result:")
	fmt.Println(string(b))
}

var (
	objType map[string]interface{} // 对象类型
	arrType []interface{}          // 数组类型
)

func merge(dst map[string]interface{}, src map[string]interface{}) {
	// 遍历原始对象的元素
	for k, v := range src {
		if v == nil {
			continue
		}

		if reflect.TypeOf(v).String() == reflect.TypeOf(arrType).String() { // value是数组类型
			// 把v转成对象数组
			objarr := v.([]interface{})
			// 检查原始数组是不是空的, 为空不处理
			if len(objarr) > 0 {
				// 检查是不是字符串数组
				if reflect.TypeOf(objarr[0]).String() == reflect.TypeOf(objType).String() { // value是对象数组
					// 判断dst[k]是否为空
					if dst[k] == nil || len(dst[k].([]interface{})) == 0 {
						dst[k] = []interface{}{map[string]interface{}{}}
					}
					// 把第一个元素转成对象类型
					newdst := dst[k].([]interface{})[0].(map[string]interface{})
					// 遍历数组，逐一合并
					for _, obj := range objarr {
						newsrc := obj.(map[string]interface{})
						if len(newsrc) > 0 {
							merge(newdst, newsrc)
						}
					}
				} else { // value是基础类型数组
					if dst[k] == nil {
						dst[k] = v
					}
				}
			} else {
				dst[k] = []interface{}{}
			}
		} else if reflect.TypeOf(v).String() == reflect.TypeOf(objType).String() { // value是对象类型
			// 判断dst[k]是否为空
			if dst[k] == nil {
				dst[k] = map[string]interface{}{}
			}
			// 直接转成对象
			newdst := dst[k].(map[string]interface{})
			newsrc := v.(map[string]interface{})
			if len(newsrc) > 0 {
				merge(newdst, newsrc)
			}
		} else { // value是基础类型（例如：字符串、布尔、整形）
			dst[k] = v
		}
	}
}
