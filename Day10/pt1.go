package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(leggiFile())
}

func leggiFile() int {
	myScanner := bufio.NewScanner(os.Stdin)
	var cicli, x, sum int
	x = 1
	controllo := 20

	for myScanner.Scan() {
		var val int
		str := strings.Fields(myScanner.Text())
		if str[0] == "noop" {
			cicli += 1
		} else {
			val, _ = strconv.Atoi(str[1])
			x += val
			cicli += 2
		}

		if cicli >= controllo && controllo <= 220 {
			sum += (x - val) * controllo
			controllo += 40
		}
	}
	return sum
}
