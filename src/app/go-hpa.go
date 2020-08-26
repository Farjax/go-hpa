package main

import(
	"math"
	"fmt"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  calcular_raiz()
  fmt.Fprintf(w, "Code.education Rocks!")
}

func main() {  
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8083", nil))
}

func calcular_raiz() float64{
  var x float64 =0.0001
  for i:=1;i<1000000;i++ {
    x+= math.Sqrt(x)
  }
  return x
}