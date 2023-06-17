package main

import (
	"fmt"
	"strings"
)

func main() {
	var nome string
	fmt.Println("Escreva uma frase: ")
	fmt.Scan(&nome)
	saida := strings.ReplaceAll(nome, " ", "")
	fmt.Println("String sem espaço em branco: ", saida)
}
