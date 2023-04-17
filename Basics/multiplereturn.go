package main

import (
	"fmt"
	"strings"
)

func initial(str string) (string, string) {
	s := strings.ToUpper(str)
	name := strings.Split(s, " ")
	var initials []string
	for _, value := range name {
		initials = append(initials, value[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	fn, ln := initial("Samundra Dahal")
	fmt.Println(fn, ln)

	fn, ln = initial("Kiran")
	fmt.Println(fn, ln)
}
