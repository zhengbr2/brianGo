package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	//return len(s[i]) < len(s[j])
	sri := []rune(s[i])
	srj := []rune(s[j])

	return sri[0] < srj[0]
}

func main() {

	strs := []string{"ac", "aa", "bb"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)

	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"

}
