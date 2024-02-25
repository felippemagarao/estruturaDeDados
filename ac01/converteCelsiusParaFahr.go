//Converte Celsius para Farhenheit

package main

import "fmt"

func converteCelsiusParaFahrenheit(celsius float64) float64 {
    return celsius*9/5 + 32
}

func main() {
    celsius := 30.0
    fahrenheit := converteCelsiusParaFahrenheit(celsius)
    fmt.Printf("%.2f graus Celsius é igual a %.2f graus Fahrenheit\n", celsius, fahrenheit)
}
