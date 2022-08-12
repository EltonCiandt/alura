package contas

import (
	"fmt"
	"strconv"
)

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}


func (c *ContaCorrente) tranferir(valorDaTranferencia float64, destinatario *ContaCorrente) bool {
	if valorDaTranferencia < c.Saldo && valorDaTranferencia > 0 {
		c.Saldo -= valorDaTranferencia
		destinatario.depositar(valorDaTranferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) depositar(valorDoDeposito float64) float64 {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		fmt.Println("Deposito realizado com sucesso", c.Saldo)
	} else {
		fmt.Println("deposito invalido")
	}
	return c.Saldo
}

func (c *ContaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		fmt.Println("Saque realizado, Saldo atual é:", c.Saldo)
	} else {
		fmt.Println("Saldo insuficiente")
	}
	return strconv.FormatBool(podeSacar)
}
