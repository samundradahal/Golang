package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func NewBill(x string) bill {
	b := bill{
		name:  x,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) format() string {
	fs := "Bill Preview\n"

	var total float64 = 0
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v... : Rs %0.2f\n", key, value)
		total += value
	}
	fs += fmt.Sprintf("%-25v ... Rs %0.2f\n", "tip", b.tip)
	fs += fmt.Sprintf("%-25v ... Rs %0.2f", "Total", total+b.tip)
	return fs
}

func (b *bill) addtip(t float64) {
	b.tip = t
}

func (b *bill) additem(str string, s float64) {
	b.items[str] = s
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("/Users/samundradahal/Desktop/Golang/struct"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill Saved")

}
