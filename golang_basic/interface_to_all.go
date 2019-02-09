/**
 interface 类型转化为各种数据结构的sample， 主要验证下面几个问题、

1. 如果数据本身是 string形式， interface 能直接转吗？ 怎么转？
2. 定义一个 含有多种数据结构的json， 然后整个json是  map[string]interface{}， 每个key的value应该怎么转数据格式？
 */
package main

import "fmt"

func testMap()  {
	// 先声明map
	var m1 map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[string]string)
	// 最后给已声明的map赋值
	m1["a"] = "aa"
	m1["b"] = "bb"

	// 直接创建
	m2 := make(map[string]string)
	// 然后赋值
	m2["a"] = "aa"
	m2["b"] = "bb"

	// 初始化 + 赋值一体化
	m3 := map[string]string{
		"a": "aa",
		"b": "bb",
	}

	// ==========================================
	// 查找键值是否存在
	if v, ok := m1["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	// 遍历map
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}

func main() {
	var it interface{}
	it = "i am it"
	fmt.Println(it)
	fmt.Printf("%s\n", it)
	fmt.Printf("%d\n", it)
	fmt.Printf("%v\n", it)
	fmt.Printf("%+v\n", it)


	itMap := make(map[string]interface{})
	itMap["name"] = "r00x"
	itMap["age"] = 18
	fmt.Printf("%v\n", itMap)
	fmt.Println(itMap["age"])
	fmt.Printf("%s\n", itMap["age"]) 		// 虽然是以interface定义， 但是还是能获取实际值的类型是int, 输出为 %!s(int=18)
}


