package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaDiego := contas.ContaCorrente{Titular: "Diego", NumeroAgencia: 589, NumeroConta: 100, Saldo: 100.90}
	contaBruna := contas.ContaCorrente{Titular: "Bruna", NumeroAgencia: 587, NumeroConta: 101, Saldo: 100.90}
	var contaPedro *contas.ContaCorrente = new(contas.ContaCorrente)
	contaPedro.Titular = "Pedro"
	fmt.Println(*contaPedro)
	fmt.Println(contaBruna)
	fmt.Println(contaDiego)

	fmt.Println(contaBruna.Sacar(-187.), " novo saldo: ", contaBruna.Saldo)

	status, valor := contaDiego.Depositar(100)
	fmt.Println(status, valor)

	statusTransferencia := contaDiego.Transferir(100, &contaBruna)

	fmt.Println(statusTransferencia)
	fmt.Println(contaDiego.Titular, " saldo: ", contaDiego.Saldo)
	fmt.Println(contaBruna.Titular, " saldo: ", contaBruna.Saldo)

}
