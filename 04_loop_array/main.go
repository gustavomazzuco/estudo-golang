package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	for indice, valor := range meuArray {
		fmt.Printf("Indice é igual %d e valor igual %d\n", indice, valor)
	}

	for _, valor := range meuArray {
		fmt.Printf("Valor igual %d\n", valor)
	}
}
