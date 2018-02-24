package main

import "fmt"

func longest(words ...string) (max int) {
	for _, word := range words {
		if len(word) > max {
			max = len(word)
		}
	}
	return max
}

func main() {
	fmt.Println(longest("Six", "sleek", "swans", "swam", "swiftly", "southwards"))
	fmt.Println(longest("Your", "word", "list", "here"))
}
