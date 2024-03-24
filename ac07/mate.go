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
