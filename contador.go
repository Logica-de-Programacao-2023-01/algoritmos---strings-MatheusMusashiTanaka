package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	frase_1 := scanner.Text()

	fmt.Print("Digite um caractere:")
	scanner.Scan()
	caractere := scanner.Text()

	contador := 0
	for i := 0; i < len(frase_1); i++ {
		if string(frase_1[i]) == caractere {
			contador++
		}
	}
	fmt.Print(contador)
}
