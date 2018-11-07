package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	result := cmplx.Pow(math.E, 1i*math.Pi) + 1
	result = cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(result)
	fmt.Printf("%.3f \n", result)
}
