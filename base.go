package main

import (
	"fmt"
	"math"
)

//包内可用
var number = 3

//number2:=3
var (
	n1     = 3
	n2     = 4
	n3 int = 5
)

func variableZeroValue() {
	var age int
	var name string
	fmt.Printf("%d %q\n", age, name)
}

func variableInitalValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "sss"
	b = 5
	fmt.Println(a, b, c, s)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	var x, y float64 = 3, 4
	z := math.Sqrt(x*x + y*y)

	fmt.Println(c, z)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	z := math.Sqrt(a*a + b*b)
	fmt.Println(z)
}

func enums() {
	const (
		cpp = iota
		_
		java
		python
		C
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)

	fmt.Println(cpp, java, python, C)
	fmt.Println(b, kb, mb, gb, tb)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swapwithoutpointer(a, b int) (int, int) {
	return b, a
}

func main() {
	//fmt.Println("Hello World!")
	//variableZeroValue()
	//variableInitalValue()
	//variableTypeDeduction()
	//variableShorter()
	//fmt.Println(n1, n2, n3)
	//euler()
	//triangle()
	//enums()
	//readFile()
	a := 1
	b := 5
	swap(&a, &b)
	fmt.Println(a, b)
}
