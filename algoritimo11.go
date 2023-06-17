package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	vogais := []string{"A", "a", "e", "E", "i", "I", "o", "O", "U", "u"}

	for _, vogal := range vogais {
		str = strings.ReplaceAll(str, vogal, "")
	}
	fmt.Printf("A string sem vogais é: %s", str)
}
