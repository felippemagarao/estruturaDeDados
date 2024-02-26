package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func adicionarContato(contato Contato, contatos [5]Contato) [5]Contato {
	for i := range contatos {
		if contatos[i].Nome == "" {
			contatos[i] = contato
			fmt.Println("Contato adicionado com sucesso!")
			return contatos
		}
	}
	fmt.Println("O array de contatos está cheio. Não foi possível adicionar o contato.")
	return contatos
}

func excluirContato(contatos [5]Contato) [5]Contato {
	for i := len(contatos) - 1; i >= 0; i-- {
		if contatos[i].Nome != "" {
			contatos[i] = Contato{}
			fmt.Println("Contato excluído com sucesso!")
			return contatos
		}
	}
	fmt.Println("Nenhum contato válido para excluir.")
	return contatos
}

func main() {
	var contatos [5]Contato
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1 - Adicionar contato")
		fmt.Println("2 - Excluir contato")
		fmt.Println("3 - Sair")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			var nome, email string
			fmt.Print("Digite o nome do contato: ")
			scanner.Scan()
			nome = scanner.Text()
			fmt.Print("Digite o email do contato: ")
			scanner.Scan()
			email = scanner.Text()

			novoContato := Contato{Nome: nome, Email: email}
			contatos = adicionarContato(novoContato, contatos)
		case 2:
			contatos = excluirContato(contatos)
		case 3:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}

		fmt.Println("\nLista de contatos:")
		for _, c := range contatos {
			if c.Nome != "" {
				fmt.Printf("Nome: %s, Email: %s\n", c.Nome, c.Email)
			}
		}
		fmt.Println()
	}
}
