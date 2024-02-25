//Construa uma função verificaParidade que receba um inteiro e retorne se o número é par ou ímpar.

package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Digite um número e descubra se ele é par ou ímpar: ")
	fmt.Scanln(&numero)

	if numero%2 == 0 {
		fmt.Println("O numero é par")
	} else if numero%2 != 0 {
		fmt.Println("O numero é impar")
	}
}
