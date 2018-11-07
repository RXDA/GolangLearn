package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "hanruia",
		"course": "golang",
		"age":    "18",
	}
	map2 := make(map[string]string) //map2 ==empty map
	var map3 map[string]string      //map3==nil

	fmt.Println(m, map2, map3)

	//遍历map,无序
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	courseName2 := m["cause"]
	fmt.Println(courseName)
	fmt.Println(courseName2)

	courseName3, contain3 := m["course"]
	courseName4, contain4 := m["cause"]
	fmt.Println(courseName3, contain3)
	fmt.Println(courseName4, contain4)

	if courseName5, contain5 := m["cause"]; contain5 == true {
		fmt.Println(courseName5)
	} else {
		fmt.Println("not conatin")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]

	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
