package main

import (
	"banco/cliente/titular"
	"fmt"
)

type ContaCorrente struct {
	titular     titular.Titular
	numeroConta int
	saldo       float64
}

func main() {
	//primeiro jeito de usar orientação a objeto
	contaDoMattheus := ContaCorrente{
		titular: titular.Titular{
			Nome:     "Mattheus",
			Cpf:      "899.899.899.89",
			Profisao: "Cienceiro"},
		numeroConta: 666,
		saldo:       1300}
	fmt.Println(contaDoMattheus)

	// segundo metodo
	clienteAlex := titular.Titular{"alex", "021.021.021.21", "analista"}
	contaDoAlex := ContaCorrente{clienteAlex, 222, 20000}
	fmt.Println(contaDoAlex)

	// terceiro modo ultilizando ponteiros
	var contaDoBuiu *ContaCorrente

	contaDoBuiu = new(ContaCorrente)
	contaDoBuiu.titular.Titular.Nome = "Buiu"
	contaDoBuiu.titular.Titular.Cpf = "555.555.555.55"
	contaDoBuiu.titular.Titular.Profisao = "professor "
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
	fmt.Println(contaDoAlex.Depositar(1000))

	contaDoAlex.Tranferir(500, &contaDoMattheus)
	fmt.Println(contaDoAlex)
	fmt.Println(contaDoMattheus)

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

// Multiplos retornos para uma função
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito invalido", c.saldo
	}
}

// Ponteiro em funçãoes apontando para multiplos objetos
func (c *ContaCorrente) Tranferir(valorDaTranferencia float64, contaDestino *ContaCorrente) bool {

	if valorDaTranferencia < c.saldo {
		c.saldo -= valorDaTranferencia
		contaDestino.Depositar(valorDaTranferencia)
		return true
	} else {
		return false
	}
}

// Ultilizando interfaces #### ESTUDAR MAIS A RESPEITO---------------------------
type verificarConta interface{
	Sacar(valor float64) string
}
func PagarBoleto(conta verificarConta, valorBoleto float64){
	conta.Sacar(valorBoleto) 
}
