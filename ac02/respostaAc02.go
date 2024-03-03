// *** Resposta da AC02
// Questão 01 
//Cria um vetor com 10 inteiros
// Imprimir os elementos do vetor um abaixo do outro

package main

import "fmt"

func main() {
    vetor := [10]int{2, 9, 7, 6, 10, 55, 62, 33, 12,74}

    for _, elemento := range vetor {
        fmt.Println(elemento)
    }
}


// Questão 02
// ***Descrição de Rune em Go***
// Um rune pode ser um único byte ou quatro bytes, um intervalo determinado pelo int32.
// Um rune em Go representa um ponto de código Unicode e pode representar um caractere.
// enquanto os caracteres ASCII podem ser representados apenas por um tipo de dados int32.

package main

import "fmt"

// A função InverterString recebe uma string e a retorna invertida
func InverterString(palavra string) string {
  runas := []rune(palavra)
  for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {
    runas[i], runas[j] = runas[j], runas[i]
    }
  return string(runas)
}

func main() {
  var texto string
  fmt.Print("Digite uma string: ")
  fmt.Scanln(&texto)

  textoInvertido := InverterString(texto)
  fmt.Println("String invertida:", textoInvertido)
}


// Questão 03
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


// Questão 04
//Package "geometria"
package geometria

func AreaRetangulo(largura, altura, float64) float64 {
	return largura*altura
}

func PerimetroRetangulo(largura, altura, float64) float64{
	return 2*(largura+altura)
}

// Código para calcular a Area e Largura do retangulo utilizando o pacote "geometria"

package main

import (
	"fmt"
	"geometria"
)

func main() {
	var largura, altura float64

	fmt.Print("Digite a largura do retângulo: ")
	fmt.Scanln(&largura)
	fmt.Print("Digite a altura do retângulo: ")
	fmt.Scanln(&altura)

	area := geometria.AreaRetangulo(largura, altura)
	perimetro := geometria.PerimetroRetangulo(largura, altura)

	fmt.Println("A área do retângulo é igual a: ", area)
	fmt.Println("O perimetro do retangulo é igual a: ", perimetro)

}

