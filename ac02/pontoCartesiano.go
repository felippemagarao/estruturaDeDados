//Cria um ponto no plano cartesiano com coordenadas X e Y que calcula a distância desse ponto até a origem (0,0).

package main

import (
	"fmt"
	"math"
)

type Ponto struct {
	X, Y float64
}

func (p Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
func main() {
	var x, y float64
	fmt.Print("Digite a coordenada no ponto X:")
	fmt.Scanln(&x)
	fmt.Print("Digite a coordenada no ponto Y:")
	fmt.Scanln(&y)
	ponto := Ponto{x, y}
	distancia := ponto.DistanciaOrigem()
	fmt.Printf("A distância do ponto (%.2f, %.2f) até a origem é %.2f\n", ponto.X, ponto.Y, distancia)
}
