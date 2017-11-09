package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type JogoDaVelha struct {
	Tabuleiro [][]string
	Marca     string
	Linha     int
	Coluna    int
}

type Jogo int

func (t *JogoDaVelha) ImprimeTabuleiro(args *JogoDaVelha, resposta *int) {
	fmt.Println("Tabuleiro do jogo: \n")

	fmt.Println("   0 1 2")

	for i, _ := range args.Tabuleiro {
		fmt.Println(i, args.Tabuleiro[i])
	}

	fmt.Println("\n\n")

	return nil
}

func (t *Jogo) Marcar(args *JogoDaVelha, resposta *JogoDaVelha) error {

	if args.Tabuleiro[args.Linha][args.Coluna] != "-" {
		return nil

	}

	resposta.Tabuleiro = args.Tabuleiro

	args.Tabuleiro[args.Linha][args.Coluna] = args.Marca

	if args.Marca == "X" {
		(*resposta).Marca = "0"

	}
	if args.Marca == "0" {
		(*resposta).Marca = "X"

	}

	return nil

}

func (t *Jogo) VerificaVencedor(args *JogoDaVelha, resposta *bool) error {

	var troca string
	if args.Marca == "X" {
		troca = "0"

	}
	if args.Marca == "0" {
		troca = "X"

	}

	if (args.Tabuleiro[1][0] == troca && args.Tabuleiro[1][1] == troca && args.Tabuleiro[1][2] == troca) ||
		(args.Tabuleiro[0][0] == troca && args.Tabuleiro[0][1] == troca && args.Tabuleiro[0][2] == troca) ||
		(args.Tabuleiro[2][0] == troca && args.Tabuleiro[2][1] == troca && args.Tabuleiro[2][2] == troca) ||
		(args.Tabuleiro[0][0] == troca && args.Tabuleiro[1][0] == troca && args.Tabuleiro[2][0] == troca) ||
		(args.Tabuleiro[0][1] == troca && args.Tabuleiro[1][1] == troca && args.Tabuleiro[2][1] == troca) ||
		(args.Tabuleiro[0][2] == troca && args.Tabuleiro[1][2] == troca && args.Tabuleiro[2][2] == troca) ||
		(args.Tabuleiro[0][0] == troca && args.Tabuleiro[1][1] == troca && args.Tabuleiro[2][2] == troca) ||
		(args.Tabuleiro[0][2] == troca && args.Tabuleiro[1][1] == troca && args.Tabuleiro[2][0] == troca) {
		*resposta = true

	} else {
		*resposta = false

	}

	return nil

}

func main() {
	jogo := new(Jogo)
	rpc.Register(jogo)

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
