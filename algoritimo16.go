package main

import (
	"fmt"
	"strings"
)

func main() {

	var str1, str2 string
	fmt.Println("Insira sua primeira string:")
	fmt.Scanln(&str1)
	fmt.Println("Insira sua segunda string:")
	fmt.Scanln(&str2)

	if strings.Contains(str1, str2) {
		fmt.Printf("%s é substring de %s!", str2, str1)
	} else {
		fmt.Printf("%s não é substring de %s!", str2, str1)
	}

}
