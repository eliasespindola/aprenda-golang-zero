package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoLower := strings.ToLower(endereco)
	primeiraPalavraDoEnderco := strings.Split(enderecoLower, " ")[0]
	//Nao temos o contains no GO
	enderecoTemUmTipoValido := contains(tiposValidos, primeiraPalavraDoEnderco)
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEnderco)
	}
	return "Tipo invalido"
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
