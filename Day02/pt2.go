package main

import (
	"bufio"
	"fmt"
	"os"
)

type mossa struct {
	nome   string
	valore int
}

func main() {
	mosse := map[string]mossa{"A": mossa{"C", 1}, "B": mossa{"A", 2}, "C": mossa{"B", 3}}
	fmt.Println(leggiFile(mosse))
}

func leggiFile(mosse map[string]mossa) int {
	myScanner := bufio.NewScanner(os.Stdin)
	var punteggio int
	for myScanner.Scan() {
		str := myScanner.Text()

		if str[2] == 'Y' {
			punteggio += 3 + mosse[string(str[0])].valore
		}
		if str[2] == 'Z' {
			n := mosse[string(str[0])].nome
			p := mosse[n].nome
			punteggio += 6 + mosse[p].valore
		}

		if str[2] == 'X' {
			n := mosse[string(str[0])].nome
			punteggio += mosse[n].valore
		}

	}
	return punteggio
}
