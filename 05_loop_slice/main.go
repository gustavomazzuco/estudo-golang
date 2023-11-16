package main

import "fmt"

func main() {
	array := [2]string{"Banana", "Maçã"}
	slice := []string{"Banana", "Maçã", "Melancia"}

	fmt.Printf("Valores Array %v\n", array)

	slice = append(slice, "Uva")
	slice = append(slice, "Goiaba")
	fmt.Printf("Valores Slice %v tamanho %d capacidade %d\n\n\n", slice, len(slice), cap(slice))

	sliceCarro := []string{"Fusca"}
	slice = append(slice, sliceCarro...)
	fmt.Printf("Valores Slice %v tamanho %d capacidade %d tipo %T\n", slice, len(slice), cap(slice), sliceCarro)

	for _, valor := range slice {
		fmt.Printf("Valor: %v\n\n", valor)
	}

	fmt.Printf("Pedaço do slice %v %v\n\n", slice[0:1], slice[3:4])

	sliceLinha := [][]string{
		[]string{"Banana", "Uva"},
		[]string{"Goiaba"},
	}

	fmt.Printf("Multi %v\n\n", sliceLinha)
}
