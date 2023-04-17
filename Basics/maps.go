package main

import "fmt"

func main() {
	//maps of string

	price_list := map[string]int{
		"Bat":    1550,
		"Ball":   90,
		"Stump":  150,
		"Gloves": 785,
	}

	//Looping Maps

	for key, value := range price_list {
		fmt.Println(key, "=", value)
	}
	//maps of int key

	name := map[int]string{
		9860621606: "Samundra",
		9860704111: "Sagar",
		9876534243: "Samir",
	}

	name[9860621606] = "Mundre"
	fmt.Println(name)
}
