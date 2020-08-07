package main

import (
	"fmt"
	"math"
)

func raizQuadrada(valor float64) float64 {
	valor = math.Sqrt(valor)
	return valor
}

func main() {
	var x float64 = 0.0001

	for i := 0; i < 10000000; i++ {
		x += raizQuadrada(x)
		fmt.Println("Code.education Rocks! ", x)
	}

}
