package main

import (
	"fmt"
)

func torreHanói(qtdDiscos int, origem, destino, auxiliar string) {
	if qtdDiscos == 1 {
		fmt.Printf("Mova o disco 1 de %s para %s\n", origem, destino)
		return
	}
	torreHanói(qtdDiscos-1, origem, auxiliar, destino)
	fmt.Printf("Mova o disco %d de %s para %s\n", qtdDiscos, origem, destino)
	torreHanói(qtdDiscos-1, auxiliar, destino, origem)
}

func calcularTotalMovimentos(qtdDiscos int) int {
	totalMovimentos := 0
	for i := 1; i <= qtdDiscos; i++ {
		totalMovimentos = totalMovimentos*2 + 1
	}
	return totalMovimentos
}

func main() {
	var qtdDiscos int
	fmt.Print("Digite a quantidade de discos na Torre de Hanói: ")
	_, err := fmt.Scan(&qtdDiscos)
	if err != nil || qtdDiscos <= 0 {
		fmt.Println("Por favor, insira um número válido de discos maior que zero.")
		return
	}
	torreHanói(qtdDiscos, "Pino A", "Pino C", "Pino B")
	totalMovimentos := calcularTotalMovimentos(qtdDiscos)
	fmt.Printf("Total de movimentos necessários: %d\n", totalMovimentos)
}
