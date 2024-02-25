//Crie uma função calculaMedia que receba um número variável de parâmetros de tipo ponto flutuante 
//e retorne a média aritmética desses valores.

package main

import "fmt"

func calculaMedia(nums ...float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func main() {
	media := calculaMedia(1.8, 2.9, 9.3)
	fmt.Println("Média:", media)
}
