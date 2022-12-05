package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(leggiFile())
}

func leggiFile() int {
	myScanner := bufio.NewScanner(os.Stdin)
	var score int

	for myScanner.Scan() {
		str := myScanner.Text()
		r := str[strings.IndexAny(str[:len(str)/2], str[len(str)/2:])]
		if unicode.IsUpper(rune(r)) {
			score += int(r-'A') + 27
		} else {
			score += int(r-'a') + 1
		}
	}
	return score
}
