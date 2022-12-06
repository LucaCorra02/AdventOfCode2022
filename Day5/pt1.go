package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	//mappa:= map[string][]string{"1":{"Z","N"},"2":{"M","C","D"},"3":{"P"}}
	mappa:= map[string][]string{"1":{"Q","M","G","C","L"},"2":{"R","D","L","C","T","F","H","G"},"3":{"V","J","F","N","M","T","W","R"},"4":{"J","F","D","V","Q","P"},"5":{"N","F","M","S","L","B","T"},"6":{"R","N","V","H","C","D","P"},"7":{"H","C","T"},"8": {"G","S","J","V","Z","N","H","P"},"9":{"Z","F","H","G"}}
	leggiFile(mappa)
	//fmt.Println((mappa))
	fmt.Println(estraiParola(mappa))

}

func leggiFile(mappa map[string][]string){
	myScanner:=bufio.NewScanner(os.Stdin)
	for myScanner.Scan(){
		tmp:=strings.Fields(myScanner.Text())
		num,_:=strconv.Atoi(tmp[1])
		ris:=len(mappa[tmp[3]])-num
		mappa[tmp[5]]=append(mappa[tmp[5]], reverse(mappa[tmp[3]][ris:len(mappa[tmp[3]])])...)
		mappa[tmp[3]]=mappa[tmp[3]][:ris]
	}
}

func estraiParola(mappa map[string][]string) string{
	var str string
	tmp:=[]string{}
	for key,_:=range mappa{
		tmp=append(tmp,key)
	}
	sort.Strings(tmp)
	for _,value:=range tmp{
		str+=mappa[value][len(mappa[value])-1]
	}
	return str
}

func reverse(str []string)[]string{
	for i:=0; i<len(str)/2;i++{
		str[i],str[len(str)-i-1]=str[len(str)-i-1],str[i]
	}
	return str
}