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

func main() {

}
