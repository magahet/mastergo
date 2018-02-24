package main

import (
	"fmt"
	"unicode"
)

func main() {
	myString := "abcde"
	for _, char := range myString {
		upper_char := unicode.ToUpper(char)
		fmt.Print(string(upper_char))
	}
	fmt.Println("\n" + myString)
}
