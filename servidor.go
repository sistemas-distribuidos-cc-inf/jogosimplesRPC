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
