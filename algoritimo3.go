package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, s1, s2 string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)
	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&s1)
	fmt.Print("Digite o novo caractere: ")
	fmt.Scanln(&s2)

	saida := strings.ReplaceAll(s, s1, s2)
	fmt.Printf("Nova string: %s\n", saida)
}
