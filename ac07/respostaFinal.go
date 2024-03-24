package main

import (
	"fmt"
	"math"
)

// Função para calcular o MDC (Máximo Divisor Comum) entre dois números.
func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return int(math.Abs(float64(a)))
}

// Função para simplificar uma fração.
func simplificar(numerador, denominador int) (int, int) {
	m := mdc(numerador, denominador)
	return numerador / m, denominador / m
}

func main() {
	var N int
	fmt.Scanln(&N) //qntd testes que devem ser lidos

	for i := 0; i < N; i++ {
		var numerador1, denominador1, numerador2, denominador2 int
		var operador string
		fmt.Scanf("%d / %d %s %d / %d\n", &numerador1, &denominador1, &operador, &numerador2, &denominador2)

		//Casos a serem analisados
		var numerador, denominador int
		switch operador {
		case "+":
			numerador = numerador1*denominador2 + numerador2*denominador1
			denominador = denominador1 * denominador2
		case "-":
			numerador = numerador1*denominador2 - numerador2*denominador1
			denominador = denominador1 * denominador2
		case "*":
			numerador = numerador1 * numerador2
			denominador = denominador1 * denominador2
		case "/":
			numerador = numerador1 * denominador2
			denominador = numerador2 * denominador1
		}

		simplifiedNumerador, simplifiedDenominador := simplificar(numerador, denominador)
		fmt.Printf("%d/%d = %d/%d\n", numerador, denominador, simplifiedNumerador, simplifiedDenominador)
	}
}


//*****************************************************************************

package main

import "fmt"

func main() {

	var N int // Numero de pessoas
	var L, Q float64 // Qntd de agua na garrafa e qntd de agua na cuia
	var nomes []string //nomes dos participantes

	part := 0

	//solicita os dados
	fmt.Scan(&N, &L, &Q) 

	for x := 0; x < N; x++ {
		var nome string
		fmt.Scan(&nome)
		nomes = append(nomes, nome)
	}

	for L-Q > 0 {
		L -= Q
		part++
	}

	part = part % N
	fmt.Printf("%s %.1f\n", nomes[part], L)

}


// *************************************************************************

package main

import "fmt"

func main() {

	var codigo1, numeroPecas1, codigo2, numeroPecas2 int
	var valorPeca1, valorPeca2 float64

	fmt.Scan(&codigo1, &numeroPecas1, &valorPeca1)

	fmt.Scan(&codigo2, &numeroPecas2, &valorPeca2)

	valorTotal := 0.0

	valorTotal = (float64(numeroPecas1) * valorPeca1) + (float64(numeroPecas2) * valorPeca2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valorTotal)
}
