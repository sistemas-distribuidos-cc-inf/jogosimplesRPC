package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type JogoDaVelha struct {
	Connection *rpc.Client
}

type EstadoDoJogo struct {
	Tabuleiro [][]string
	Marca     string
	Linha     int
	Coluna    int
}

func (t *EstadoDoJogo) AlocaTabuleiro() {
	for i, _ := range t.Tabuleiro {
		t.Tabuleiro[i] = make([]string, 3)
		for j, _ := range t.Tabuleiro[i] {
			t.Tabuleiro[i][j] = "-"
		}
	}
}

func (t *JogoDaVelha) Inserir(args *EstadoDoJogo) {

	fmt.Println("Digite a linha: ")
	fmt.Scanf("%d", &args.Linha)
	fmt.Println("Digite a coluna: ")
	fmt.Scanf("%d", &args.Coluna)

	if (args.Linha <= 2) || (args.Linha <= 0) {
		err := t.Connection.Call("Jogo.Marcar", args, &args)
		if err != nil {
			log.Fatal("erro no jogo1: ", err)

		}

	} else {
		fmt.Println("linha ou coluna fora invalidas")

	}

}

func (t *JogoDaVelha) ImprimirJogo(args *EstadoDoJogo) {

	var resposta int

	err := t.Connection.Call("Jogo.ImprimeTabuleiro", args, &resposta)
	if err != nil {
		log.Fatal("erro no jogo: ", err)

	}

}

func (t *JogoDaVelha) VerificarVencedor(args *EstadoDoJogo) bool {

	var verificador bool

	err := t.Connection.Call("Jogo.VerificaVencedor", args, &verificador)
	if err != nil {
		log.Fatal("erro no jogo3: ", err)

	}

	fmt.Println(verificador)

	return verificador

}

func main() {

}
