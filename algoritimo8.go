package main

import "fmt"

func main() {

	slice := []string{"a", "a", "b", "c", "d", "a", "ab", "c"}
	var valor string
	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)

	var resultado []string

	for _, x := range slice {
		if x != valor {
			resultado = append(resultado, x)
		}
	}

	fmt.Println(resultado)
}
