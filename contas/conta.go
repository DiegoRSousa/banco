package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor <= 0 {
		return "valor inválido!"
	}
	if valor <= c.Saldo {
		return "saldo insuficiente"
	}

	c.Saldo -= valor
	return "saque realizado com sucesso"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor <= 0 {
		return "valor inválido!", c.Saldo
	}

	c.Saldo += valor
	return "deposito realizado com sucesso", c.Saldo
}

func (c *ContaCorrente) Transferir(valor float64, destino *ContaCorrente) bool {
	if valor <= 0 {
		return false
	}

	c.Sacar(valor)
	destino.Depositar(valor)
	return true
}
