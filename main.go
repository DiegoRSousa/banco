package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor <= 0 {
		return "valor para saque inválido!"
	}
	if valor <= c.saldo {
		c.saldo -= valor
		return "saque realizado com sucesso!"
	}
	return "saldo insuficiente!"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor <= 0 {
		return "valor para saque inválido!", c.saldo
	}
	c.saldo += valor
	return "deposito realizado com sucesso!", c.saldo
}

func main() {

	contaDiego := ContaCorrente{titular: "Diego", numeroAgencia: 589, numeroConta: 100, saldo: 100.90}
	contaBruna := ContaCorrente{"Bruna", 587, 101, 100.90}
	var contaPedro *ContaCorrente = new(ContaCorrente)
	contaPedro.titular = "Pedro"
	fmt.Println(*contaPedro)
	fmt.Println(contaBruna)
	fmt.Println(contaDiego)

	fmt.Println(contaBruna.Sacar(-187.), " novo saldo: ", contaBruna.saldo)

	status, valor := contaDiego.Depositar(100)
	fmt.Println(status, valor)

}
