package main

import (
	"fmt"
	"strings"
)

func main() {

	var palavra string

	fmt.Print("Digite uma palavra: ")
	fmt.Scanln(&palavra)

	reverso := ""
	for i := len(palavra) - 1; i >= 0; i-- {
		reverso += string(palavra[i])
	}

	if strings.EqualFold(palavra, reverso) {
		fmt.Printf("%s é um palíndromo:", palavra)
	} else {
		fmt.Printf("%s não é um palíndromo:", palavra)
	}
}
