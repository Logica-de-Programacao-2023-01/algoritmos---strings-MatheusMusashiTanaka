package main
import "fmt"

func main() {
  array := [4]float64{2.6,8.4,20.2,14.5}
  var multiplicacao float64
  
  for _, x := range array {
      multiplicacao = multiplicacao * x
  }
  produto := multiplicacao * float64(len(array))
  fmt.Println("O produto de 2.6, 8.4, 20.2, 14.5 é", produto)
}
