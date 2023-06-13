package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}

	var num int

	fmt.Print("Informe um número inteiro: ")
	fmt.Scanln(&num)

	encontrado := false
	for _, valor := range slice {

		if num == valor {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println("O número pertence à Slice!")
	} else {
		slice = append(slice, num)
		fmt.Println("O número inserido foi adicionado à Slice:", slice)
	}
}
