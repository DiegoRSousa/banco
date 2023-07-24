package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor <= 0 {
		return "valor inválido!"
	}
	if valor <= c.saldo {
		return "saldo insuficiente"
	}

	c.saldo -= valor
	return "saque realizado com sucesso"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor <= 0 {
		return "valor inválido!", c.saldo
	}

	c.saldo += valor
	return "deposito realizado com sucesso", c.saldo
}

func (c *ContaCorrente) Transferir(valor float64, destino *ContaCorrente) bool {
	if valor <= 0 {
		return false
	}

	c.Sacar(valor)
	destino.Depositar(valor)
	return true
}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}
