package main

import (
	"fmt"
)

func main() {
	var quantidadeDeJogos int
	fmt.Scan(&quantidadeDeJogos)

	for i := 0; i < quantidadeDeJogos; i++ {

		var numeroJogado int
		fmt.Scan(&numeroJogado)

		if numeroJogado%2 == 0 {
			fmt.Println("ALICE")
		} else {
			fmt.Println("BOB")
		}
	}
}
