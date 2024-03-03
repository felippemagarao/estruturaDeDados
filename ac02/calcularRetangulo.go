package main

import (
	"fmt"
	"AC03/geometria"
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
