package main

import (
	"fmt"
	"os"
)


func main(){

	if len(os.Args)!=2 {
		fmt.Println("no file inputted")
		return
	}
	input:= os.Args[1]
	data, err:= os.ReadFile(input + ".txt")
	if err !=nil{
		fmt.Println("file not found")
		return
	}
	result:=string(data)
	fmt.Println(result)
}