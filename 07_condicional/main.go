package main

import "fmt"

func main() {
	alunos := map[string]int{"Gustavo": 7, "Mariane": 9}
	alunos["José"] = 0

	if alunos["Gustavo"] >= 6 {
		fmt.Println("Passou")
	} else if alunos["Gustavo"] < 6 {
		fmt.Println("Não Passou")
	}

	cor := 1
	switch cor {
	case 1:
		fmt.Println("Amarelo")
	case 2:
		fmt.Println("Vermelho")
	default:
		fmt.Println("Sem cor")
	}

	for aluno, nota := range alunos {
		if nota == 0 {
			continue
		}
		fmt.Printf("Aluno %v com nota %v\n", aluno, nota)
	}
}
