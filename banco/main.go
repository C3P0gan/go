package main

import (
	// "banco/clientes"
	"banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	//	contaDoCristoffer := contas.ContaCorrente{
	//		Titular:    "Cristoffer",
	//		nroAgencia: 589,
	//		nroConta:   123456,
	//		Saldo:      125.5,
	//	}
	//	contaDaAline := contas.ContaCorrente{"Aline", 589, 654321, 200.8}
	//	contaDoCris := contas.ContaCorrente{"Cris", 589, 123456, 12.1}
	//	fmt.Println(contaDoCris)

	//	var contaDaBruna *contas.ContaCorrente
	//	contaDaBruna = new(contas.ContaCorrente)
	//	contaDaBruna.Titular = "Bruna"
	//	contaDaBruna.nroAgencia = 616
	//	contaDaBruna.nroConta = 616616
	//	contaDaBruna.Saldo = 666

	//	var contaDaBruna2 *contas.ContaCorrente
	//	contaDaBruna2 = new(contas.ContaCorrente)
	//	contaDaBruna2.Titular = "Bruna"
	//	contaDaBruna2.nroAgencia = 616
	//	contaDaBruna2.nroConta = 616616
	//	contaDaBruna2.Saldo = 666

	//  fmt.Println(contaDoCristoffer)
	//  fmt.Println(contaDaAline)

	//	fmt.Println(contaDaBruna, contaDaBruna2, contaDaBruna == contaDaBruna2)     // false memAddr
	//	fmt.Println(&contaDaBruna, &contaDaBruna2, &contaDaBruna == &contaDaBruna2) // false pointer
	//	fmt.Println(*contaDaBruna, *contaDaBruna2, *contaDaBruna == *contaDaBruna2) // true  content

	//	contaDaSilvia := contas.ContaCorrente{}
	//	contaDaSilvia.Titular = "Silvia"
	//	contaDaSilvia.Saldo = 500

	//	fmt.Println(contaDaSilvia.Saldo)
	//	fmt.Println(contaDaSilvia.Sacar(100))
	//	fmt.Println(contaDaSilvia.Saldo)
	//    status, valor := contaDaSilvia.Depositar(2000)
	//    fmt.Println(status, valor)

	//	fmt.Println("Silvia", contaDaSilvia.Saldo)
	//	fmt.Println("Aline", contaDaAline.Saldo)
	//	status := contaDaSilvia.Transferir(300, &contaDaAline)
	//	fmt.Println(status)
	//	fmt.Println("Silvia", contaDaSilvia.Saldo)
	//	fmt.Println("Aline", contaDaAline.Saldo)

	//	contaDoBruno := contas.ContaCorrente{
	//		Titular: clientes.Titular{
	//			Nome:      "Bruno",
	//			CPF:       "123.111.123.12",
	//			Profissao: "Desenvolvedor Go",
	//		},
	//		NroAgencia: 123,
	//		NroConta:   123456,
	//		Saldo:      100,
	//	}

	//  clienteBruno := clientes.Titular{"Bruno", "123.111.123.12", "Desenvolvedor Go"}
	//  contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	//	fmt.Println(contaDoBruno)

	//  contaExemplo := contas.ContaCorrente{}
	//  contaExemplo.Depositar(100)
	//  fmt.Println(contaExemplo.ObterSaldo())

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())

    contaDaLuisa := contas.ContaCorrente{}
    contaDaLuisa.Depositar(500)
    PagarBoleto(&contaDaLuisa, 400)
    fmt.Println(contaDaLuisa.ObterSaldo())
}
