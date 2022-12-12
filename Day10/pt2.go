package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	leggiFile2()
	fmt.Println()
}

func leggiFile2() {
	myScanner := bufio.NewScanner(os.Stdin)
	var cicle, x int
	controllo := 41
	x = 1
	cicle = 0

	for myScanner.Scan() {
		str := strings.Fields(myScanner.Text())
		cicle += 1
		t(controllo, &cicle)
		controlla(cicle, x)
		if str[0] == "addx" {
			t(controllo, &cicle)
			cicle += 1
			t(controllo, &cicle)
			controlla(cicle, x)
			val, _ := strconv.Atoi(str[1])
			x += val
		}
	}
}

func controlla(cicle, x int) {
	//fmt.Println("x: ",x," Cicle: ",cicle)
	if x+1 == cicle-1 || x-1 == cicle-1 || x == cicle-1 {
		fmt.Print("█")
	} else {
		fmt.Print(" ")
	}
}

func t(controllo int, cicle *int) {
	if *cicle >= controllo {
		fmt.Println()
		*cicle -= 40
	}
}
