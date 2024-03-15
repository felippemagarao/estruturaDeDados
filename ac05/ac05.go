package main

import (
    "fmt"
)

func descrescente(a, b, c float64) (float64, float64, float64) {
    if a < b {
        a, b = b, a
    }
    if a < c {
        a, c = c, a
    }
    if b < c {
        b, c = c, b
    }
    return a, b, c
}

func main() {
    var A, B, C float64

    fmt.Scan(&A, &B, &C)

    A, B, C = descrescente(A, B, C)

    if A >= (B + C) {
        fmt.Println("NAO FORMA TRIANGULO")
    } else {
        if (A*A) == (B*B)+(C*C) {
            fmt.Println("TRIANGULO RETANGULO")
        }

        if (A*A) > (B*B)+(C*C) {
            fmt.Println("TRIANGULO OBTUSANGULO")
        }

        if (A*A) < (B*B)+(C*C) {
            fmt.Println("TRIANGULO ACUTANGULO")
        }

        if A == B && B == C {
            fmt.Println("TRIANGULO EQUILATERO")
        }

        if (A == B && B != C) || (A == C && A != B) || (B == C && A != B) {
            fmt.Println("TRIANGULO ISOSCELES")
        }
    }
}
