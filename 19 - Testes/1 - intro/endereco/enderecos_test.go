package enderecos

import (
	"testing"
)

// Criando cenários de teste

type cenarioDeTeste struct{
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T){
	// TestNomeFuncExato
	// parametro é um ponteiro t

	// t.Parallel() testes podem rodar em paralelo

	cenarioDeTeste := []cenarioDeTeste{
		{"Avenida Paulista", "Avenida"},
		{"Rodovia das Rodas", "Rodovia"},
		{"Rua ABC", "Rua"},
		{"Praça das Rosas", "Tipo Invalido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA DOS BOBOS", "Avenida"},
	}

	for _, cenario := range cenarioDeTeste{
		tipoEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoEnderecoRecebido != cenario.retornoEsperado{
			t.Errorf("O tipo recebido %s é diferente do esperado %s", 
			tipoEnderecoRecebido, cenario.retornoEsperado)
		}
	}

	/**
	go test -> executa testes
	go teste ./... entra em todos os packages e procura testes e faz os testes - testes podem ficar em caching
	go test -v -> modo verboso, mais descritivo
	go test --cover -> cobertura de teste
	go test --coverprofile cobertura.txt -> arquivo de relatorio gerado
	go tool cover --func=cobertura.txt -> leitura do nosso arquivo criado
	go tool cover --func=cobertura.html -> arquivo visual das linhas nao cobertas em golang
	
	*/ 
	

}