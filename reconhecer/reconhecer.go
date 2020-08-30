// Pacote com funções auxiliares para reconher sequências de strings.
package reconhecer

import (
	"regexp"

	"github.com/jhudsonsg/go-scanner/core"
)

// HeNumero - verifica se a sequencia informada é um número.
func HeNumero(str string) bool {
	match, _ := regexp.MatchString("[0-9]", str)
	return match
}

// HeReal - verifica se a sequencia informada é um número real.
func HeReal(str string) bool {
	match, _ := regexp.MatchString("[0-9]*\\.[0-9]*", str)
	return match
}

// HeCaractere - verifica se a sequencia informada faz parte do alfabeto [a-z].
func HeCaractere(str string) bool {
	match, _ := regexp.MatchString("[a-zA-Z]", str)
	return match
}

// HeLogico - verifica se a sequencia informada é um termo verdadeiro true ou falso.
func HeLogico(str string) bool {
	match, _ := regexp.MatchString("[a-zA-Z]", str)
	return match
}

// HePalavraReservada - verifica se é uma palavra chave.
func HePalavraReservada(str string) (bool, string) {
	for valor, tipoToken := range core.PalavrasChaves {
		if valor == str {
			return true, tipoToken
		}
	}

	return false, ""
}

// HeDelimitadorDeInicio - verifica se é o operador de início do programa.
func HeDelimitadorDeInicio(str string) bool {
	if str == "INICIO" {
		return true
	}
	return false
}

// HeDelimitadorDeFim - verifica se é o operador de fim do programa.
func HeDelimitadorDeFim(str string) bool {
	if str == "FIM" {
		return true
	}
	return false
}
