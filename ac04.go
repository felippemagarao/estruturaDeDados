//****** QUESTÃO 01 *****
//Implemente em Go a solução recursiva para o problema da Torre de Hanói.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var discos int

	// Solicita ao usuario a quantidade de discos
	fmt.Println("*****Solução para Torre de Hanói*****")
	fmt.Println("Insira a quantidade de discos: ")
	fmt.Scanln(&discos)
	fmt.Println("")

	movimentos := hanoi(discos, "A", "C", "B")

	//Faz o calculo para saber quantos movimentos foram utilizados
	//Strconv converte a entrada em um numero inteiro
	fmt.Printf("\nPara resolver esse a Torre de Hanói com %s discos, são necessários %s movimentos", strconv.Itoa(discos), strconv.Itoa(movimentos))
	fmt.Println("")
}

func hanoi(n int, origem, destino, auxiliar string) int {
	movimentos := 0

	if n == 1 {
		fmt.Println("Disco " + strconv.Itoa(n) + ": " + origem + " vai para " + destino)
		return 1
	}

	// O disco na torre de origem vai para a torre auxiliar
	// A torre de destino vira auxiliar
	movimentos += hanoi(n-1, origem, auxiliar, destino)

	// Os discos da torre de origem vão a torre de destino
	fmt.Println("Disco " + strconv.Itoa(n) + ": " + origem + " vai para " + destino)
	movimentos++

	//O disco na torre auxiliar vai para torre destino e usa origem como auxiliar
	movimentos += hanoi(n-1, auxiliar, destino, origem)

	return movimentos
}

//****** QUESTÃO 02 *****
//Implemente em Go um algoritmo que resolva o seguinte problema: 
// dado um array de números inteiros positivos, considerado ordenado crescentemente, 
// e um valor alvo, encontre um par de números no array cuja soma seja igual ao valor alvo. 
// Se nenhum par for encontrado, retorne um valor (-1, -1) indicando que nenhum par foi encontrado. 
// Resolva esse problema com um algoritmo cuja complexidade é O(n).

package main

import "fmt"

func encontrarPar(arr []int, alvo int) (int, int) {
	// Inicia os dois ponteiros e
	//O for faz um looping para eles se encontrarem
	esquerda := 0
	direita := len(arr) - 1

	for esquerda < direita {
		soma := arr[esquerda] + arr[direita]
		if soma == alvo {
			return arr[esquerda], arr[direita]
		} else if soma < alvo {
			esquerda++
		} else {
			direita--
		}
	}
	return -1, -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	alvo := 17

	x, y := encontrarPar(arr, alvo)
	if x != -1 && y != -1 {
		fmt.Printf("Par encontrado: %d, %d\n", x, y)
	} else {
		fmt.Println("Nenhum par encontrado.")
	}
}
