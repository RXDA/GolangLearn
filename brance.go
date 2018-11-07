package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	const path = "abc.txt"
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//or
	if contents, err := ioutil.ReadFile(path); err != nil {
		fmt.Println()
	} else {
		fmt.Println(contents)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "D"
	case score < 90:
		g = "C"
	case score < 95:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprint(
			"Wrong score: %d", score))
	}
	return g
}

func main() {
	readFile()

	fmt.Println(
		grade(0),
		grade(40),
		grade(65),
		grade(85),
		grade(94),
		grade(96),
		grade(101),
	)
}
