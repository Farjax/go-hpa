package main

import(
	"math"
	"fmt"
)
func main(){
	calcular_raiz()
  fmt.Printf("Code.education Rocks!!")
}

func calcular_raiz() float64{
  var x float64 =0.0001
  for i:=1;i<100000000;i++ {
    x+= math.Sqrt(x)
  }
  return x
}