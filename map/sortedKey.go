package main

import (
	"fmt"
	"sort"
)

func main() {
	var m = map[string]int{
		"unix":         0,
		"python":       1,
		"go":           2,
		"javascript":   3,
		"testing":      4,
		"philosophy":   5,
		"startups":     6,
		"productivity": 7,
		"hn":           8,
		"reddit":       9,
		"C++":          10,
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for indx, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k], " index:", indx)
	}
}
