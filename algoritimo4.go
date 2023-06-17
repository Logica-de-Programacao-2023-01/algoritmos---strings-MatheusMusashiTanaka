package main

import (
	"fmt"
)

func main() {

	var sum string
	var sdois string
	fmt.Println("Digite duas strings")
	fmt.Scan(&sum)
	fmt.Scan(&sdois)

	if sum == sdois {
		fmt.Println("As strings são iguais")
	} else {
		fmt.Println("As strings são diferentes")
	}

}
