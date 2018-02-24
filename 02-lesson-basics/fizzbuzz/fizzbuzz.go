package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(n int) {
	for i := 0; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
	// If the number is divisible by 15, print “FizzBuzz”.
	// If the number is divisible by 3, print “Fizz” instead of the number.
	// If the number is divisible by 5, print “Buzz”.
}
func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
