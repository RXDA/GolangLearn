package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//s:=arr[2:6] //2,3,4,5
	//fmt.Println("arr[2:6]=",arr[2:6])//2,3,4,5
	//fmt.Println("arr[2:6]=",arr[ :6])//0,1,2,3,4,5
	//fmt.Println("arr[2:6]=",arr[2: ])//2,3,4,5,6,7,8
	//fmt.Println("arr[2:6]=",arr[ : ])//all

	s1 := arr[2:]
	//s2:=arr[:]

	fmt.Println(s1)
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

}
