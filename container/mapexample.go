package main

import (
	"fmt"
	"unicode/utf8"
)

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, char := range []byte(s) {

		if lastI, ok := lastOccurred[char]; ok == true && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[char] = i
	}
	return maxLength
}

func lengthOfLongestSubstringWithRune(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, char := range []byte(s) {

		if lastI, ok := lastOccurred[char]; ok == true && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[char] = i
	}
	return maxLength
}

func main() {
	s := "我！韩瑞达！打钱！"
	fmt.Println(s, len(s))

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		//fmt.Println(i,ch)
		fmt.Printf("%d,%c  ", i, ch)
	}

}
