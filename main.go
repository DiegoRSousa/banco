package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaDiego := ContaCorrente{titular: "Diego", numeroAgencia: 589, numeroConta: 100, saldo: 100.90}
	contaBruna := ContaCorrente{"Bruna", 587, 101, 100.90}
	var contaPedro *ContaCorrente = new(ContaCorrente)
	contaPedro.titular = "Pedro"
	fmt.Println(*contaPedro)
	fmt.Println(contaBruna)
	fmt.Println(contaDiego)
}
