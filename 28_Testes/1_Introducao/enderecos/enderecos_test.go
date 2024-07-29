package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTest := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"RUA ABC", "Rua"},
		{"rua abc", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"AVENIDA ABC", "Avenida"},
		{"avenida abc", "Avenida"},
		{"Estrada ABC", "Estrada"},
		{"ESTRADA ABC", "Estrada"},
		{"estrada ABC", "Estrada"},
		{"Rodovia ABC", "Rodovia"},
		{"RODOVIA ABC", "Rodovia"},
		{"rodovia ABC", "Rodovia"},
		{"Bla ABC", "Tipo invalido"},
		{"bla ABC", "Tipo invalido"},
		{"BLA ABC", "Tipo invalido"},
		{"", "Tipo invalido"},
	}

	for _, cenario := range cenariosDeTest {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado)
		}
	}

}
