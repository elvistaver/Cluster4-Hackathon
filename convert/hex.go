package main

import (
	"fmt"
	"log"
	"strconv"
)


func hexToDec(s string) (int64, error){

	hex, err:= strconv.ParseInt(s, 16, 64)

	if err !=nil{
		log.Println("invalid converstion")
	}
	return hex, nil
}

func main(){
	fmt.Println(hexToDec("1E"))
	fmt.Println(hexToDec("FF"))
}
