package main

import "fmt"

var decimal = []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
var symbol = []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%v = %v\n", i, romanizer(i))
	}
}

func romanizer(n int) string {
	roman := ""
	for i, v := range decimal {
		for n >= decimal[i] {
			roman += symbol[i]
			n -= v
		}
	}
	return roman
}
