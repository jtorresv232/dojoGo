package main
import "fmt"
import "math"

func  Sqrt(x, y float64) float64  {
  // codigo que implementa la raiz
  // hint 1: usar ciclo que itere 10 veces
  z := float64(x/2)
  var w float64 =0
  var cont int = 0
  for {
    cont++
    w=z
    z=z-(((z*z)-x)/(2*z))

    if w-z < y {
      fmt.Println(cont)
      return z
    }

  }
}

func main()  {
  fmt.Println(Sqrt(10, 1e-6))
  fmt.Println(math.Sqrt(10))
}
