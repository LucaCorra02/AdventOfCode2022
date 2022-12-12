package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items []int
	op string
	test int
	true int
	false int
}

func main(){
	ri:=leggiFile()
	fmt.Println("ris",rounds(ri))
}

func leggiFile()(ris []Monkey){
	myScanner:=bufio.NewScanner(os.Stdin)
	var cur Monkey

	for myScanner.Scan(){
		tmp:=strings.Fields(myScanner.Text())
		if len(tmp)>0 {
			if tmp[0] == "Monkey" {
				cur = Monkey{[]int{}, "", 0, 0, 0}
			} else if tmp[0] == "Starting" {
				cur.items = converti(tmp[2:])
			} else if tmp[0] == "Operation:" {
				cur.op=strings.Join(tmp[3:]," ")
			}else if tmp[0]=="Test:"{
				val,_:=strconv.Atoi(tmp[3])
				cur.test=val
			}else if tmp[0]+tmp[1]=="Iftrue:"{
				val,_:=strconv.Atoi(tmp[5])
				cur.true=val
			}else{
				val,_:=strconv.Atoi(tmp[5])
				cur.false=val
				ris=append(ris,cur)
			}
		}
	}
	return
}

func converti(val []string)(ris []int){
	for _,v:= range val{
		ri,_:=strconv.Atoi(strings.ReplaceAll(v,",",""))
		ris=append(ris,ri)
	}
	return
}


func rounds(in []Monkey)(item int){
	var ri int
	max :=make([]int,8)

	for ri < 20{
		for i:=0;i<len(in);i++{
			max[i]+=round(in,i)
		}
		ri++
	}
	sort.Ints(max)
	return max[len(max)-1]*max[len(max)-2]
}

func round(in []Monkey,indice int )(it int ){
	it=len(in[indice].items)
	for _,val:= range in[indice].items{
		//rimuovere elemento che se ne va
		sos:=op(val,in[indice].op)/3
		if sos%in[indice].test==0{
			i:=in[indice].true
			in[i].items=append(in[i].items,sos)
		}else{
			i:=in[indice].false
			in[i].items=append(in[i].items,sos)
		}
	}
	in[indice].items=nil
	return
}


func op( old int,operazione string)(ris int){
	new:=strings.Fields(strings.ReplaceAll(operazione,"old",strconv.Itoa(old)))
	p1,_:=strconv.Atoi(new[0])
	p2,_:=strconv.Atoi(new[2])
	if new[1]=="+"{
		ris=p1+p2
	}else{
		ris=p1*p2
	}
	return
}