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
