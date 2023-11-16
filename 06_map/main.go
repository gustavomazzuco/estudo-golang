package main

import "fmt"

func main() {
	alunos := map[string]int{"Gustavo": 7, "Mariane": 9}

	fmt.Println(alunos)

	alunos["José"] = 7
	fmt.Println(alunos)
	fmt.Println(alunos["Mariane"])

	valor, ok := alunos["João"]
	fmt.Printf("\nValor é %v se existe %v\n\n", valor, ok)

	valor, ok = alunos["Gustavo"]
	fmt.Printf("\nValor é %v se existe %v\n\n", valor, ok)

	delete(alunos, "José")
	fmt.Println(alunos)

	for key, aluno := range alunos {
		fmt.Printf("Nota do aluno %v é igual a %v\n", key, aluno)
	}

	clear(alunos)
	fmt.Println(alunos)
}
