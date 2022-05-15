package contas

import (
	"banco/clientes"
)

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) Sacar(saque float64) string {
	if c.saldo >= saque && saque > 0 {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	}
	return "Valor inválido"
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Depósito realizado com sucesso.", c.saldo
	} else {
		return "Valor do depósito inválido.", c.saldo
	}
}
