package main

import "fmt"

type ContaCorrente struct {
	titular     string
	numeroConta int
	saldo       float64
}

func main() {
	//primeiro jeito de usar orientação a objeto
	contaDoMattheus := ContaCorrente{
		titular:     "Mattheus",
		numeroConta: 666,
		saldo:       1300}
	fmt.Println(contaDoMattheus)

	// segundo metodo
	contaDoAlex := ContaCorrente{"Alex", 222, 20000}
	fmt.Println(contaDoAlex)

	// terceiro modo ultilizando ponteiros
	var contaDoBuiu *ContaCorrente

	contaDoBuiu = new(ContaCorrente)
	contaDoBuiu.titular = "Buiu"
	contaDoBuiu.numeroConta = 500
	contaDoBuiu.saldo = 100

	fmt.Println(*contaDoBuiu)

	/*comparação de objetos em GO com metodos
	buleanos vai comparar sempre o conteudo deles
	ou seja se um objeto tiver o mesmo conteudo
	do outro objeto sera retornado um TRUE.
	*/

	contaDoMattheus.Sacar(300)
	fmt.Println(contaDoMattheus.saldo)
}

// funçãoes e metodos para um objeto
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado"
	} else {
		return "Saldo insuficiente"
	}
}
