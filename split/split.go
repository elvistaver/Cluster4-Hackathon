package main

import (
	"fmt"
	"strings"
	"unicode"
)

func split(text string)[]string{
	result:= strings.Fields(text)
	return result
}

func Split(text string)[]string{
	result:=[]string{}

	tempra:=[]rune{}

	for _, chr:= range text{

		if unicode.IsPunct(chr){
			result=append(result, string(tempra))
			tempra=[]rune{}
		}

		if chr!=' '{
			tempra=append(tempra, chr)
			continue
		}
		if len(tempra) != 0{
			result = append(result, string(tempra))
			tempra=[]rune{}
			continue
		}	
	}
	if len(tempra)!=0{
		result = append(result, string(tempra))
	}
	return result
}
func main(){
	fmt.Println(split("Hello, world! How are you?"))
	fmt.Println(Split("Hello, world! How are you?"))
}