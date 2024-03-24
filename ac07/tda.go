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
