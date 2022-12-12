package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(leggiFile2())
}

func leggiFile2() int {
	myScanner := bufio.NewScanner(os.Stdin)
	var cont int
	for myScanner.Scan() {
		tmp := strings.Split(myScanner.Text(), ",")
		slc1 := strings.Split(tmp[0], "-")
		slc2 := strings.Split(tmp[1], "-")
		s1, _ := strconv.Atoi(slc1[0])
		e1, _ := strconv.Atoi(slc1[1])
		s2, _ := strconv.Atoi(slc2[0])
		e2, _ := strconv.Atoi(slc2[1])

		if !(e1 < s2 || e2 < s1) {
			cont++
		}

	}
	return cont
}
