//Cria um vetor com 10 inteiros
// Imprimi os elementos do vetor um abaixo do outro

package main

import "fmt"

func main() {
    vetor := [10]int{2, 9, 7, 6, 10, 55, 62, 33, 12,74}

    for _, elemento := range vetor {
        fmt.Println(elemento)
    }
}

