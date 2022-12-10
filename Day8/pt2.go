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

func visibili(griglia [][]int)(int){
	var max int
	for i:=1;i<len(griglia)-1;i++{
		for j:=1;j<len(griglia)-1;j++{
			ok1,val1:=visit(griglia[i][j+1:],griglia[i][j])
			ok2,val2:=visit(rev(griglia[i][:j]),griglia[i][j])
			ok3,val3:=visitBasso(griglia,i,j)
			ok4,val4:=visitaAlto(griglia,i,j)
			if ok1 || ok2 || ok3 || ok4{
				ri:=val1*val2*val3*val4
				if ri > max{
					max=ri
				}
			}
		}
	}
	return max
}

func rev(tmp[]int)[]int{
	var ris []int
	for i:=len(tmp)-1;i>=0;i--{
		ris=append(ris,tmp[i])
	}
	return ris
}

func visit(tmp[]int,val int ) (bool,int){
	var cont int
	for i:=0;i<len(tmp);i++{
		if tmp[i] >= val{
			cont++
			return false,cont
		}
		cont++
	}
	return true,cont
}

func visitBasso(griglia [][]int,i,j int)(bool,int){
	var cont int
	for k:=i+1;k<len(griglia);k++{
		if griglia[k][j] >= griglia[i][j]{
			cont++
			return false,cont
		}
		cont++
	}
	return true,cont
}

func visitaAlto(griglia [][]int,i,j int)(bool,int){
	var cont int
	for k:=i-1; k>=0 ; k--{
		if griglia[k][j] >= griglia[i][j]{
			cont++
			return false,cont
		}
		cont++
	}
	return true,cont
}
