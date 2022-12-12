package main

import (
	"fmt"
	"strings"
)

func main() {
	numCaratteri := contaStr()
	fmt.Println(numCaratteri)
}

func contaStr() int {
	var str string
	fmt.Scan(&str)

	var ris string

	for i := 0; i < len(str); i++ {
		ris += string(str[i])
		for j := i + 1; j < len(str); j++ {
			if !strings.ContainsRune(ris, rune(str[j])) {
				ris += string(str[j])
				if len(ris) == 14 { //per parte 1 inserire sequenze di 4
					return j + 1
				}
			} else {
				ris = ""
				break
			}
		}
	}
	return -1
}
