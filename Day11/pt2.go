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
	items []int64
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
				cur = Monkey{[]int64{}, "", 0, 0, 0}
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

func converti(val []string)(ris []int64){
	for _,v:= range val{
		ri,_:=strconv.Atoi(strings.ReplaceAll(v,",",""))
		ris=append(ris,int64(ri))
	}
	return
}


func rounds(in []Monkey)(item int64){
	var ri int
	max :=make([]int64,8)

	for ri < 10000{
		for i:=0;i<len(in);i++{
			max[i]+=round(in,i)
		}
		ri++
	}
	fmt.Println(max)
	return
}


func round(in []Monkey,indice int )(it int64 ){
	it=int64(len(in[indice].items))
	lc:=lcm(in)

	for _,val:= range in[indice].items{
		sos:=op(val,in[indice].op)%lc
		if sos % int64(in[indice].test)==0{
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

func lcm(in []Monkey)(lc int64){
	lc=1
	for _,m:=range in {
		t:=int64(m.test)
		lc*=t
	}
	return
}

func op( old int64,operazione string)(ris int64){
	new:=strings.Fields(strings.ReplaceAll(operazione,"old",strconv.FormatInt(old,10)))
	p1,_:=strconv.ParseInt(new[0],10,64)
	p2,_:=strconv.ParseInt(new[2],10,64)
	if new[1]=="+"{
		ris=(p1+p2)
	}else{
		ris=(p1*p2)
	}
	return
}
