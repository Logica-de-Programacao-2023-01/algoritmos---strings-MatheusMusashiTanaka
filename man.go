package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	fmt.Print("digite um numero: ")
	scanner.Scan()
	numero, err := strconv.Atoi(scanner.Text())

	resultado := ""

	if err != nil {
		fmt.Println("O numero não é valido")
	} else {
		for i := 0; i < len(frase); i++ {
			if i < numero {
				resultado = resultado + (strings.ToUpper(string(frase[i])))
			} else {
				resultado = resultado + string(frase[i])
			}

		}
	}
	fmt.Println(resultado)
	fmt.Println("")
}
-------------------------------------------------------------------------

ou



package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	fmt.Print("digite um numero: ")
	scanner.Scan()
	numero, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("O numero não é valido")
	} else {
		resultado := strings.ToUpper(frase[:numero]) + frase[numero:]
		fmt.Println(resultado)
	}
}
