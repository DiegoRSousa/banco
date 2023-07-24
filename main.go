package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	clienteBruno := clientes.Titular{Nome: "Bruno", Cpf: "038", Profissao: "Desenvolvedor"}
	contaBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 99, NumeroConta: 1234}

	contaBruno.Depositar(100.)

	fmt.Println(contaBruno.Titular.Nome, contaBruno.Saldo())
}
