package main

import (
	"log"
	"sort"
	"strings"
)

var s = [][]string{
	{
		"abc", "bca",
	},
	{
		"abc", "bad",
	},
}

func main() {
	for i, strings := range s {
		permutation := CheckPermutation(strings[0], strings[1])
		log.Println(i, permutation)
	}
}
func CheckPermutation(s1 string, s2 string) bool {
	s1s := strings.Split(s1, "")
	s2s := strings.Split(s2, "")
	sort.Strings(s1s)
	sort.Strings(s2s)
	s1j := strings.Join(s1s, "")
	s2j := strings.Join(s2s, "")
	return s1j == s2j
}
