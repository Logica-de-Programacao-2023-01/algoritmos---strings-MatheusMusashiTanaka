package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Digite uma string:")
	fmt.Scanln(&str)

	newstr := strings.ReplaceAll(str, "a", "*")
	newstr = strings.ReplaceAll(newstr, "e", "*")
	newstr = strings.ReplaceAll(newstr, "i", "*")
	newstr = strings.ReplaceAll(newstr, "o", "*")
	newstr = strings.ReplaceAll(newstr, "u", "*")
	newstr = strings.ReplaceAll(newstr, "A", "*")
	newstr = strings.ReplaceAll(newstr, "E", "*")
	newstr = strings.ReplaceAll(newstr, "I", "*")
	newstr = strings.ReplaceAll(newstr, "O", "*")
	newstr = strings.ReplaceAll(newstr, "U", "*")
	fmt.Println("Nova string:", newstr)
}
