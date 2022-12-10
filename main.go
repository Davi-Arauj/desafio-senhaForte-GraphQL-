package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	palavra := "tEsr21#T!@3#85"

	splitPalavra := strings.Split(palavra, "")
	var matched bool
	for _, letra := range splitPalavra {
		match := regexp.MustCompile(letra + `{2}`)
		matched = match.MatchString(palavra)
		if matched {
			fmt.Println(letra)
			break
		}
	}

	fmt.Println(matched)

}
