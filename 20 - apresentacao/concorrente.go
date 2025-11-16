package main

import (
	"fmt"
	"sync"
	"time"
)

// Cria um WaitGroup. Ele é usado para esperar que um conjunto de goroutines termine.
// É uma forma de sincronização essencial.
var waitGroup sync.WaitGroup

// Função que simula a mesma tarefa demorada.
func trabalhoConcorrente(numero int) {
	// 1. Ao final da execução desta função, notifica o WaitGroup que uma Goroutine terminou.
	defer waitGroup.Done()

	// 2. O loop demorado que consome tempo.
	for i := 0; i < 500000000; i++ {
		// Simulação do trabalho da CPU.
	}
	
	// 3. Imprime uma mensagem de conclusão.
	fmt.Println("Tarefa", numero, "concluída concorrentemente.")
}

func main() {
	// 4. Marca o tempo de início.
	inicio := time.Now()

	fmt.Println("Iniciando execução concorrente...")

	// 5. Adiciona 3 contadores ao WaitGroup. Isso diz ao programa: "Espere por 3 tarefas".
	waitGroup.Add(3)

	// 6. Chamada de múltiplas tarefas usando Goroutines:
	// A palavra-chave 'go' inicia a função em uma nova Goroutine (thread leve).
	// O programa principal NÃO espera que 'trabalhoConcorrente(x)' termine para iniciar o próximo.
	go trabalhoConcorrente(1)
	go trabalhoConcorrente(2)
	go trabalhoConcorrente(3)

	// 7. A função Wait() bloqueia a execução da função main até que o contador do WaitGroup
	// chegue a zero (ou seja, até que todas as 3 goroutines chamem 'Done()').
	waitGroup.Wait()

	// 8. Calcula o tempo total decorrido.
	fim := time.Since(inicio)

	// 9. Imprime o resultado. Este tempo será significativamente menor que o Exemplo 1.
	fmt.Println("Execução Concorrente concluída em:", fim)
}