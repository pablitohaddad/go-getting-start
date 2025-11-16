package main

import (
	"fmt"
	"time"
)

// Função que simula uma tarefa demorada (Ex: cálculo complexo, I/O pesado).
func trabalhoDemorado(numero int) {
	// 1. Inicia um loop que consome tempo para simular o trabalho.
	for i := 0; i < 500000000; i++ {
		// Esta linha apenas consome ciclos de CPU, mantendo a função ocupada.
	}
	// 2. Imprime uma mensagem indicando que a tarefa (com o número específico) foi concluída.
	fmt.Println("Tarefa", numero, "concluída sequencialmente.")
}

func main() {
	// 3. Marca o tempo de início para medir a performance.
	inicio := time.Now()

	fmt.Println("Iniciando execução sequencial...")

	// 4. Chamada de múltiplas tarefas DEMORADAS em sequência.
	// O programa espera o 'trabalhoDemorado(1)' terminar antes de começar o 'trabalhoDemorado(2)'.
	trabalhoDemorado(1)
	trabalhoDemorado(2)
	trabalhoDemorado(3)

	// 5. Calcula o tempo total decorrido.
	fim := time.Since(inicio)

	// 6. Imprime o resultado para comparação.
	fmt.Println("Execução Sequencial concluída em:", fim)
}