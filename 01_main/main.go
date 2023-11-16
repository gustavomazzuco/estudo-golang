package main

import "fmt"

var idade int

type nome string

var razao nome

const UF string = "SC"

var cidade = "Criciuma"

var ativo = false
var inativo bool
var tempo int
var sobrenome string

func main() {
	razao = "Gustavo"
	idade = 26

	fmt.Println("Hello Word!!", razao, idade, UF, cidade, ativo)
	//Imprimir a tipagem das variaveis
	fmt.Printf("Hello Word!! %T %T %T %T %T\n ", razao, idade, UF, cidade, ativo)

	estado := "SC"
	fmt.Printf("Hello Word!! %T %v\n ", estado, estado)

	fmt.Printf("%T %v\n ", inativo, inativo)
	fmt.Printf("%T %v\n ", tempo, tempo)
	fmt.Printf("%T %v\n ", sobrenome, sobrenome)
}
