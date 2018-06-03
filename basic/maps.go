package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)//m2 == empty mpa

	var m3 map[string]int//m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range  m{//key的顺序使无序的，map是hashmap
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	if causeName, ok := m["cause"]; ok{
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}
	fmt.Println(courseName, ok)
	coursename := m["course"]//这里如果输入错了，会重新创建一个key对应value类型的初始值
	fmt.Println(coursename)

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	//删除操作
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
