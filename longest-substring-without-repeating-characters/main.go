package main

import "fmt"

func main() {
	fmt.Println("1", lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println("2", lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println("3", lengthOfLongestSubstring("pwwkew") == 3)
}

func lengthOfLongestSubstring(s string) int {
	// use a window of values and count while there arent any repeating vals
	// create a map of byte bool
	chars := map[byte]bool{}
	var out int
	l := 0
	for r := range s {
		for chars[s[r]] {
			delete(chars, s[l])
			l++
		}
		chars[s[r]] = true
		out = max(out, r-l+1)

	}
	fmt.Println(s, out)
	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
