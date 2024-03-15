// Questão 1

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World!")

}

// Questão 02
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


// Questão 03

package main

import (
    "fmt"
)

func main() {
    var T int
    fmt.Scanln(&T)

    for i := 0; i < T; i++ {
        var PA, PB int
        var G1, G2 float64
        fmt.Scanln(&PA, &PB, &G1, &G2)

        anos := 0
        for PA <= PB {
            PA += int(float64(PA) * G1 / 100)
            PB += int(float64(PB) * G2 / 100)
            anos++
            if anos > 100 {
                fmt.Println("Mais de 1 seculo.")
                break
            }
        }

        if anos <= 100 {
            fmt.Printf("%d anos.\n", anos)
        }
    }
}
