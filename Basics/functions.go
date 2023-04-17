package main

import (
	"fmt"
	"math"
)

func good(str string) { //Non Return Type function
	fmt.Println("Good morning,", str)
}

func message(x []string, f func(string)) { //Passing functions as argument
	for _, value := range x {
		f(value)
	}
}

func calc(r float64) float64 { //Retrun type function
	return math.Pi * r * r
}

func main() {
	good("Samundra")
	st := []string{"Samundra", "Kiran", "Sagar"}
	message(st, good)

	a1 := calc(7.5)
	a2 := calc(9.32)

	fmt.Println(a1, a2)

}
