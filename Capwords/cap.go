package main

import (
	"fmt"
	"strings"
)

func capitalize(s string)string{

	input:= strings.Fields(s)
	result:=" "

	for _, text:= range input{
		text= strings.ToUpper(text[:1]) + strings.ToLower(text[1:])
		result+= " "+text
	}
	 return result
}

func main(){
	fmt.Println(capitalize("whAT iS YOU upto"))
}
