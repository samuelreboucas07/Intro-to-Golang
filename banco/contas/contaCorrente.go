package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	if c.Saldo >= saque && saque > 0 {
		c.Saldo -= saque
		return "Saque realizado com sucesso"
	}
	return "Valor inválido"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.Saldo += valor
		return "Depósito realizado com sucesso.", c.Saldo
	} else {
		return "Valor do depósito inválido.", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor <= c.Saldo && valor > 0 {
		c.Saldo -= valor
		contaDestino.Depositar(valor)
		return true
	}
	return false
}
