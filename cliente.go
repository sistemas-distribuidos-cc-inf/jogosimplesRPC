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

func (t *JogoDaVelha) Connectar() {

	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("erro na conexao: ", err)

	}

	t.Connection = client

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

	return verificador

}

func main() {
	var args *EstadoDoJogo = &EstadoDoJogo{}
	args.Tabuleiro = make([][]string, 3)
	args.Marca = "X"
	args.Linha = 0
	args.Coluna = 0

	args.AlocaTabuleiro()

	var jogo *JogoDaVelha = &JogoDaVelha{}
	jogo.Connectar()

	jogo.ImprimirJogo(args)

	for i := 0; i < 9; i++ {
		if jogo.VerificarVencedor(args) {
			if args.Marca == "X" {
				args.Marca = "0"

			} else {
				args.Marca = "X"

			}
			fmt.Printf("Jogador %s venceu\n", args.Marca)
			break

		}
		jogo.Inserir(args)
		jogo.ImprimirJogo(args)

	}
}
