package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	nome    string
	dimFile int
	padre   *File
	figli   map[string]*File
}

func main() {
	fmt.Println(leggiFile())
}

func leggiFile() int {
	myScanner := bufio.NewScanner(os.Stdin)
	var curDir *File
	var dir []*File

	for myScanner.Scan() {
		str := strings.Split(myScanner.Text(), " ")
		if str[1] == "cd" {
			if str[2] != ".." {
				if str[2] == "/" {
					curDir = &File{"/", 0, nil, make(map[string]*File)}
					dir = append(dir, curDir)
				} else {
					curDir = curDir.figli[str[2]]
				}
			} else {
				curDir = curDir.padre
			}
		}
		if str[0] == "dir" {
			curDir.figli[str[1]] = &File{str[1], 0, curDir, make(map[string]*File)}
			dir = append(dir, curDir.figli[str[1]])
		}
		if val, err := strconv.Atoi(str[0]); err == nil {
			curDir.dimFile += val
		}
	}
	return somma(dir)
}

func somma(dir []*File) (sum int) {
	for _, val := range dir {
		tmp := sommaFigli(val)
		if val.dimFile+tmp <= 100000 {
			sum += val.dimFile + tmp
		}
	}
	return
}

func sommaFigli(n *File) (size int) {
	for _, d := range n.figli {
		size += d.dimFile + sommaFigli(d)
	}
	return
}
