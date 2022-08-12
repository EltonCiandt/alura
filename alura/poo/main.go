package main

import (
	"fmt"
)


func main() {

	contaSilvia := ContaCorrente{titular: "Silvia", saldo: 3000.00}
	fmt.Println(contaSilvia)

	contaGustavo := ContaCorrente{titular: "Gustavo", saldo: 9000}
	fmt.Println(contaGustavo)

	status := contaSilvia.tranferir(200, &contaGustavo)
	fmt.Println(status)

	fmt.Println("o saldo final da conta do gustavo é:", contaGustavo)
	fmt.Println("o saldo final da conta do Silvia é:", contaSilvia)

}
