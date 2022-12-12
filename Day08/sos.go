package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pair struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	trees := make(map[pair]int)
	y := 0
	for scanner.Scan() {
		input := scanner.Text()
		for x := 0; x < len(input); x++ {
			height, _ := strconv.Atoi(string(input[x]))
			trees[pair{x, y}] = height
		}
		y++
	}
	p1, p2 := advent(trees)
	fmt.Println("Parte 1: ", p1)
	fmt.Println("Parte 2: ", p2)
}
func advent(trees map[pair]int) (part1, part2 int) {
	for p := range trees { //fisso un punto
		visible, score := 0, 1

		for _, tipoRicerca := range []pair{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			i := 0
			newp := pair{p.x + tipoRicerca.x, p.y + tipoRicerca.y}
			for {
				_, ok := trees[newp] //il nuovo albero esiste?
				if !ok {
					visible, score = 1, score*i //faccio visible=1 invece che +=
					//se arrivo con una delle ricerche al bordo, vuol dire che l'albero fissato è visibile
					//calcolo lo scenic score con i
					break
				}

				if trees[newp] >= trees[p] {
					score *= i + 1
					//dato l'albero fissato, ho trovato un albero piu alto, che quindi blocca la visuale
					//calcolo scenic score con i+1 perche quello che blocca conta
					break
				}
				i++
				newp = pair{newp.x + tipoRicerca.x, newp.y + tipoRicerca.y}
			}
		}
		//alla fine per l albero fissato ho eseguito le 4 ricerche
		part1 += visible
		if score > part2 {
			part2 = score
		}
	}
	return
}
