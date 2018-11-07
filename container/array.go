package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

func printArray2(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{6, 5, 4, 2, 1}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)
	fmt.Println("--------------")
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	fmt.Println("--------------")
	for i := range arr3 {
		fmt.Println(arr3[i], i)
	}
	fmt.Println("--------------")
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	fmt.Println("--------------")
	for _, v := range arr3 {
		fmt.Println(v)
	}
	fmt.Println("--------------")
	printArray(arr1)
	printArray(arr3)
	//printArray(arr2)
	fmt.Println("--------------")
	printArray2(&arr1)
	printArray2(&arr3)

}
