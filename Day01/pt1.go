package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(leggiFile())
}

func leggiFile() int {
	scanner := bufio.NewScanner(os.Stdin)
	var max, max2, max3, sum int

	for scanner.Scan() {
		str := scanner.Text()

		if str == "" || scanner.Err() != nil {
			if sum > max {
				max = sum
			} else if sum > max2 {
				max2 = sum
			} else if sum > max3 {
				max3 = sum
			}
			sum = 0
		} else {
			val, _ := strconv.Atoi(str)
			sum += val
		}
	}
	return max + max2 + max3
}
