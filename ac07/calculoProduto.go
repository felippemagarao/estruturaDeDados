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
