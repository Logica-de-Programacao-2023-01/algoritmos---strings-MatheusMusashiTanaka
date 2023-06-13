package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var num int

	fmt.Print("Digite um número: ")
	fmt.Scanln(&num)

	encontrado := false
	for _, valor := range array {

		if num == valor {
			encontrado = true
			break
		}
	}
	if encontrado {
		fmt.Print("O seu número pertence a Array!")
	} else {
		fmt.Print("O seu número não pertence a Array!")
	}

}
