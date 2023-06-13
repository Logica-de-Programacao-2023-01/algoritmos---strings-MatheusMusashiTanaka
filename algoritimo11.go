
package main

import "fmt"

func main() {

	array := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var linha, coluna int

	fmt.Print("Informe a linha: ")
	fmt.Scanln(&linha)
	fmt.Print("Informe a coluna: ")
	fmt.Scanln(&coluna)

	fmt.Printf("O valor armazenado na posição (%d,%d) é %d", linha, coluna, array[linha][coluna])
}
