package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println(visibili(leggiFile()))
}

func leggiFile()(griglia [][]int){
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		row := make([]int, len(line))
		for i, c := range line {
			row[i] = int(c - '0')
		}
		griglia= append(griglia, row)
	}
	return griglia
}

func visibili(griglia [][]int)(vis int){
	for i:=1;i<len(griglia)-1;i++{
		for j:=1;j<len(griglia)-1;j++{
			if visit(griglia[i][j+1:],griglia[i][j])|| visit(griglia[i][:j],griglia[i][j])  ||visitBasso(griglia,i,j) || visitaAlto(griglia,i,j)  {
				vis++
			}
		}
	}
	vis+=2*len(griglia)+((len(griglia)-2)*2)
	return
}

func visit(tmp[]int,val int ) bool{
	for i:=0;i<len(tmp);i++{
		if tmp[i] >= val{
			return false
		}
	}
	return true
}

func visitBasso(griglia [][]int,i,j int)bool{
	for k:=i+1;k<len(griglia);k++{
		if griglia[k][j] >= griglia[i][j]{
			return false
		}
	}
	return true
}

func visitaAlto(griglia [][]int,i,j int)bool{
	for k:=i-1; k>=0 ; k--{
		if griglia[k][j] >= griglia[i][j]{
			return false
		}
	}
	return true
}