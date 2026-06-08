package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("incomplete or invalid command")
		return
	}
	input := os.Args[1]
	output:= os.Args[2]
	data, err := os.ReadFile(input + ".txt")
	if err != nil {
		fmt.Println("file not found")
		return
	}
	read:= string(data)
	result:=strings.ToUpper(read)

	err = os.WriteFile(output +".txt", []byte(result), 0664)

	if err !=nil{
		fmt.Println("file empty or currpted")
		return
	}
	fmt.Println("sucessful; cat output.txt to see formatted text")
}
