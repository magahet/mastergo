package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func acronym(s string) (acr string) {
	space := false
	for idx, r := range s {
		if (space || idx == 0) && unicode.IsLetter(r) {
			acr += string(unicode.ToUpper(r))
			space = false
		} else if unicode.IsSpace(r) {
			space = true
		}
	}
	return acr
}

func main() {
	s := "Pan Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronym(s))
}
