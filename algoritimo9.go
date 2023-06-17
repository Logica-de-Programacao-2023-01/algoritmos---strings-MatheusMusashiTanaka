package main

import (
	"fmt"
	"strings"
)

func main() {

	var str, cha, cha2 string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	fmt.Print("Digite o caractere que deseja alterar: ")
	fmt.Scanln(&cha)
	fmt.Print("Digite o caractere que deseja colocar no lugar desse: ")
	fmt.Scanln(&cha2)

	substituir := strings.Replace(str, cha, cha2, -1)

	fmt.Printf("%s fica %s", str, substituir)

}
