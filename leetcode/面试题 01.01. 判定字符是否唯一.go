package main

import (
	"log"
	"sort"
	"strings"
)

// https://leetcode-cn.com/problems/is-unique-lcci/
var s = []string{
	"leetcode",
	"abc",
	"strings",
	"stringsstrings",
	"tt6j54l54rv4s6gd4crqqwm40000gn",
	"aaaaaaaaaaa",
	"123aa",
	"123456",
	"abcdef",
}

func main() {
	for i, s2 := range s {
		unique := isUnique(s2)
		log.Println(i, unique)
	}
}
func isUnique(astr string) bool {
	split := strings.Split(astr, "")
	sort.Strings(split)
	for i := 0; i < len(split)-1; i++ {
		if split[i] == split[i+1] {
			return false
		}
	}
	return true
}
