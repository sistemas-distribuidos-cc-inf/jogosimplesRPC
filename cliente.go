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

func main() {

}
