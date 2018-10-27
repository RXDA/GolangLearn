package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupport:" + op)
	}
}

func div(a, b int) (result, yu int) {
	result = a / b
	yu = a % b
	return result, yu
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args"+"(%d.%d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(values ... int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func main() {
	fmt.Println(div(19, 3))
	fmt.Println(eval(3, 4.5, "*"))

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
	fmt.Println(sum(1,2,3,4,5,6,6,7))
}
