package main

import (
	"fmt"
	"log"
	"strconv"
)

func binToDec(s string) (int64, error) {

	bin, err := strconv.ParseInt(s, 2, 64)

	if err != nil {
		log.Println("invalid converstion")
	}
	return bin, nil
}

func main() {
	fmt.Println(binToDec("10"))
	fmt.Println(binToDec("1010"))
	fmt.Println(binToDec("11111111"))
}
