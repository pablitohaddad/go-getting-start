package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço temum tipo vpalido
func TipoDeEndereco(endereco string) string{
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(strings.ToLower(endereco), " ") [0] // pega a primeira palavra de um string

	// RUA DOS BOBOS
	// ["RUA", "DOS", "BOBOS"]

	enderecoTipoValid := false

	for _, tipo := range tiposValidos{
		if tipo == primeiraPalavraDoEndereco{
			enderecoTipoValid = true
		}
	}

	if enderecoTipoValid {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Invalido"

}