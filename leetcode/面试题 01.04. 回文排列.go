package main

import "log"

var s = []string{
	"tactcoa",
	"abba",
	"aabbccbbcc",
	"abb",
	"aaa",
	"aaaaaa",
	"tactca",
	"asdasdasd",
	"atfe",
	"aabbccc",
}

func main() {
	for _, s2 := range s {
		palindrome := canPermutePalindrome(s2)
		log.Println(s2, palindrome)
	}
}

func canPermutePalindrome(s string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
		if m[v]%2 == 0 {
			delete(m, v)
		}
	}
	return len(m) <= 1
}
