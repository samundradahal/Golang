package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//Standard Library of Strings

	intro := "Hello World"

	fmt.Println(strings.Contains(intro, "Hi"))
	fmt.Println(strings.Contains(intro, "Hello"))

	fmt.Println(strings.ReplaceAll(intro, "Hello", "Hi"))
	fmt.Println(strings.ToLower(intro))
	fmt.Println(strings.Split(intro, " "))

	// Slices Standard Library

	age := []int{20, 10, 45, 65, 78}
	fmt.Println(sort.IntSlice(age))
	fmt.Println(sort.SearchInts(age, 65))

}
