package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("1: expect true - ", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("2: expect false - ", isPalindrome("race a car"))
	fmt.Println("3: expect true - ", isPalindrome(" "))
	fmt.Println("3: expect false - ", isPalindrome("0P"))
}

func isPalindrome(str string) bool {
	// initiate l and r pointers at ends of str
	// move inwards ignoring non alphanumerica values
	// return false if evaluated string values are not equal

	l, r := 0, len(str)-1
	runeStr := []rune(str)
	for l < r {
		if !isAlnum(runeStr[l]) {
			l++
			continue
		}
		if !isAlnum(runeStr[r]) {
			r--
			continue
		}
		if !strings.EqualFold(string(str[l]), string(str[r])) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlnum(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}
