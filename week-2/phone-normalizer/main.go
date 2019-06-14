package main

import (
	"fmt"
	"strings"
)

func main() {
	phones := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	fmt.Println(phoneNormalizer(phones))
}

func phoneNormalizer(phones []string) (m map[string]int) {
	m = map[string]int{}
	for _, p := range phones {
		p = strings.ReplaceAll(p, " ", "")
		p = strings.ReplaceAll(p, "-", "")
		p = strings.ReplaceAll(p, "(", "")
		p = strings.ReplaceAll(p, ")", "")
		m[p] = m[p] + 1
	}
	return m
}
