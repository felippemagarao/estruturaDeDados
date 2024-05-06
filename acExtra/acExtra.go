package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Fila struct {
	tamanho int
	inicio  *No
	fim     *No
}

func (f *Fila) percorre() {
	if f.inicio == nil {
		fmt.Println("Fila vazia")
	} else {
		no := f.inicio
		for no != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println()
	}
}

func (f *Fila) insere(novoValor int) {
	novoNo := &No{valor: novoValor}

	if f.inicio == nil {
		f.inicio = novoNo
		f.fim = novoNo
	} else {
		f.fim.prox = novoNo
		f.fim = novoNo
	}

	f.tamanho++
}

func (f *Fila) remove() error {
	if f.inicio == nil {
		return fmt.Errorf("Fila vazia")
	}

	if f.tamanho == 1 {
		f.inicio = nil
		f.fim = nil
	} else {
		f.inicio = f.inicio.prox
	}

	f.tamanho--
	return nil
}

func main() {
	fila := Fila{}
	fila.insere(6)
	fila.insere(9)
	fila.insere(31)

	fila.percorre()
	fmt.Println(fila.tamanho)

	fila.remove()

	fila.percorre()
	fmt.Println(fila.tamanho)

	fila.remove()
	fila.remove()

	fila.percorre()
	fmt.Println(fila.tamanho)

	err := fila.remove()
	fmt.Println(err)
}
