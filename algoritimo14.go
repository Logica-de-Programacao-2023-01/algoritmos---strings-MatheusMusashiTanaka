package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	for i := 0; i < len(str)-1; i++ {
		if strings.Compare(string(str[i]), string(str[i+1])) <= 0 {
			fmt.Println("Não está em ordem decrescente")
			return
		}
	}
	fmt.Println("Está em ordem decrescente")
}
