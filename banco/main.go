package main

// go mod init myapp

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	// contaSamuel := contas.ContaCorrente{Titular: "Samuel", Agencia: 10000, Conta: 200000, Saldo: 1000.33}
	// contaPedro := contas.ContaCorrente{Titular: "Pedro", Agencia: 107676, Conta: 124332434, Saldo: 200}
	// fmt.Println(contaSamuel)

	// var contaCris *ContaCorrente
	// contaCris := new(ContaCorrente)
	// contaCris.titular = "Cris"

	// fmt.Println(*contaCris)
	// contaSilvia := ContaCorrente{}
	// contaSilvia.titular = "Silvia"
	// contaSilvia.saldo = 2000
	// fmt.Println(contaSilvia)
	// contaSilvia.depositar(100)
	// fmt.Println(contaSilvia)
	// fmt.Println("Antes de transferência")
	// fmt.Println(contaSamuel)
	// fmt.Println(contaPedro)
	// status := contaSamuel.Transferir(200, &contaPedro)
	// if status {
	// 	fmt.Println("Após de transferência")
	// 	fmt.Println(contaSamuel)
	// 	fmt.Println(contaPedro)
	// } else {
	// 	fmt.Println("Erro na transferência")
	// }
	contaBruno := contas.ContaCorrente{
		Titular: clientes.Titular{Nome: "Bruno", CPF: "123456789", Profissao: "Programador"},
		Agencia: 10000,
		Conta:   200000,
		Saldo:   1000.33,
	}

	fmt.Println(contaBruno)

	clienteSamuel := clientes.Titular{Nome: "Samuel", CPF: "123456789", Profissao: "Programador"}
	contaSamuel := contas.ContaCorrente{
		Titular: clienteSamuel,
		Agencia: 10000,
		Conta:   200000,
		Saldo:   1000.33}

	fmt.Println(contaSamuel)

	pagarBoleto(&contaSamuel, 100)

	fmt.Println(contaSamuel)
}

func pagarBoleto(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}

type verificarConta interface {
	Sacar(valor float64) string
}
