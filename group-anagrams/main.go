package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println("1", reflect.DeepEqual(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}))
	fmt.Println("2", reflect.DeepEqual(groupAnagrams([]string{""}), [][]string{{""}}))
	fmt.Println("3", reflect.DeepEqual(groupAnagrams([]string{"a"}), [][]string{{"a"}}))
}

func groupAnagrams(strs []string) [][]string {
	strMap := map[string][]string{}

	for _, str := range strs {
		sortedStr := sortStr(str)
		strMap[sortedStr] = append(strMap[sortedStr], str)
	}

	var out [][]string
	for _, val := range strMap {
		out = append(out, val)
	}

	return out
}

func sortStr(str string) string {
	runeStr := []rune(str)

	sort.Slice(runeStr, func(a, b int) bool {
		return runeStr[a] < runeStr[b]
	})
	return string(runeStr)
}
