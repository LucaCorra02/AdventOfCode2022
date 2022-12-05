package main

import (
	"bufio"
	"fmt"
	"os"
)

type mossa struct {
	nome   string
	corisp string
	valore int
}

func main() {
	mosse := map[string]mossa{"Z": mossa{"B", "C", 3}, "Y": mossa{"A", "B", 2}, "X": mossa{"C", "A", 1}}
	fmt.Println(leggiFile(mosse))
}

func leggiFile(mosse map[string]mossa) int {
	myScanner := bufio.NewScanner(os.Stdin)
	var punteggio int
	for myScanner.Scan() {
		str := myScanner.Text()

		if mosse[string(str[2])].nome == string(str[0]) {
			punteggio += 6 + mosse[string(str[2])].valore
		} else {
			if mosse[string(str[2])].corisp == string(str[0]) {
				punteggio += 3 + mosse[string(str[2])].valore
			} else {
				punteggio += mosse[string(str[2])].valore
			}
		}
	}
	return punteggio
}
