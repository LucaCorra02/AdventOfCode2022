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
	var tmp []string

	for myScanner.Scan() {
		tmp = append(tmp, myScanner.Text())
		if len(tmp) == 3 {
			r := tmp[0][contenuta(tmp)]
			if unicode.IsUpper(rune(r)) {
				score += int(r-'A') + 27
			} else {
				score += int(r-'a') + 1
			}
			tmp = nil
		}
	}
	return score
}

func contenuta(slc []string) int {
	for pos, car := range slc[0] {
		if strings.ContainsRune(slc[1], car) && strings.ContainsRune(slc[2], car) {
			return pos
		}
	}
	return -1
}
