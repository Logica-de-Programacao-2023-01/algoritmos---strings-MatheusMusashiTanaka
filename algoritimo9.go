package main

import "fmt"

func main() {

	Array := [6]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}

	var num float64
	fmt.Print("Informe um número para repor todas as posições do Array: ")
	fmt.Scanln(&num)

	for i := 0; i < len(Array); i++ {
		Array[i] = num
	}
	fmt.Println("Slice atual:", Array)
}
