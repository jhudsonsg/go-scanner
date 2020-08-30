// Pacote responsável por gerar tokens de um determinado arquivo.
package main

import (
	"log"

	"github.com/jhudsonsg/go-scanner/core"
	"github.com/jhudsonsg/go-scanner/reconhecer"
)

// proximoToken - lê varios caracteres para gerar um token reconhecido pela linguagem.
func proximoToken() core.Token {
	for caractere := lerProximoCaractere(); caractere != ""; caractere = lerProximoCaractere() {

		if reconhecerComentarios(caractere) {
			continue
		}

		if reconhecerEspacos(caractere) {
			continue
		}

		if token := reconhecerOperadorLogicos(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerOperadorAritimetica(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerLimitadores(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerNumeros(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerCadeias(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerCaracteres(caractere); token.Type != "" {
			return token
		}

		log.Fatal("Erro do tipo léxico, Token: ", caractere, ", não conhecido. Linha: ", linhaAtual)
	}

	return core.Token{}
}

// reconhecerComentarios - pula linha com comentários.
func reconhecerComentarios(caractere string) bool {
	if caractere == "/" {
		caractere = lerProximoCaractere()

		if caractere == "/" {
			for caractere := lerProximoCaractere(); caractere != "\n"; caractere = lerProximoCaractere() {
			}

			avancarLinha()
			return true
		}
	}

	return false
}

// reconhecerEspacos - pula linhas com espaços em brancos ou quebra de linha
func reconhecerEspacos(caractere string) bool {
	if caractere == " " || caractere == "\r" {
		return true
	}

	if caractere == "\n" {
		avancarLinha()
		return true
	}

	return false
}

// reconhecerOperador - reconhece tokens de operadores lógicos.
func reconhecerOperadorLogicos(caractere string) core.Token {
	if caractere == "=" {
		lexema := caractere
		caractere = lerProximoCaractere()

		if caractere == "=" {
			lexema += caractere
			return core.Token{Type: core.OPERADOR_LOGICO_IGUALDADE, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		RetrocederNoBuffer()
		return core.Token{Type: core.OPERADOR_ATRIBUICAO, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	if caractere == ">" {
		lexema := caractere
		caractere = lerProximoCaractere()

		if caractere == "=" {
			lexema += caractere
			return core.Token{Type: core.OPERADOR_LOGICO_MAIOR_QUE, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		RetrocederNoBuffer()
		return core.Token{Type: core.OPERADOR_MAIOR, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	if caractere == "<" {
		lexema := caractere
		caractere = lerProximoCaractere()

		if caractere == "=" {
			lexema += caractere
			return core.Token{Type: core.OPERADOR_LOGICO_MENOR_QUE, Lexema: lexema, LinhaDoToken: linhaAtual}
		} else if caractere == ">" {
			lexema += caractere
			return core.Token{Type: core.OPERADOR_LOGICO_DIFERENTE, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		RetrocederNoBuffer()
		return core.Token{Type: core.OPERADOR_LOGICO_MENOR, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	return core.Token{}
}

// reconhecerOperadorAritimetica - reconhece tokens de operadores aritimeticos.
func reconhecerOperadorAritimetica(caractere string) core.Token {
	if caractere == "+" {
		return core.Token{Type: core.OPERADOR_ARITMETICO_SOMA, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "-" {
		return core.Token{Type: core.OPERADOR_ARITMETICO_SUBTRACAO, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "*" {
		return core.Token{Type: core.OPERADOR_ARITMETICO_MULTIPLICACAO, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "/" {
		return core.Token{Type: core.OPERADOR_ARITMETICO_DIVISAO, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	return core.Token{}
}

// reconhecerLimitadores - reconhece tokens de limitadores.
func reconhecerLimitadores(caractere string) core.Token {
	if caractere == "{" {
		return core.Token{Type: core.ABRE_CHAVE, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "}" {
		return core.Token{Type: core.FECHA_CHAVE, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "(" {
		return core.Token{Type: core.ABRE_PARENTESE, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == ")" {
		return core.Token{Type: core.FECHA_PARENTESE, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	return core.Token{}
}

// reconhecerCaracteres - reconhe tokens de caracteres, distinguindo-os palavras chaves de variáveis.
func reconhecerCaracteres(caractere string) core.Token {
	if reconhecer.HeCaractere(caractere) {
		lexema := caractere

		for caractere := lerProximoCaractere(); reconhecer.HeCaractere(caractere) || reconhecer.HeNumero(caractere); caractere = lerProximoCaractere() {
			lexema += caractere
		}

		RetrocederNoBuffer()

		if hePalavraReservada, tipoToken := reconhecer.HePalavraReservada(lexema); hePalavraReservada {
			return core.Token{Type: tipoToken, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		return core.Token{Type: core.VARIAVEL, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	return core.Token{}
}

// reconhecerNumeros - reconhe tokens de sequ^qncia de números, distinguindo-os de reais e inteiros.
func reconhecerNumeros(caractere string) core.Token {
	if reconhecer.HeNumero(caractere) {
		lexema := caractere

		for caractere := lerProximoCaractere(); reconhecer.HeNumero(caractere) || caractere == "."; caractere = lerProximoCaractere() {
			lexema += caractere
		}

		RetrocederNoBuffer()

		if reconhecer.HeReal(lexema) {
			return core.Token{Type: core.TIPO_NUMERO_REAL, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		return core.Token{Type: core.TIPO_NUMERO_INTEIRO, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	return core.Token{}
}

// reconhecerCadeias - reconhece tokens de cadeira com inicio de uma "'".
func reconhecerCadeias(caractere string) core.Token {
	if caractere == "'" {
		lexema := caractere

		for caractere := lerProximoCaractere(); caractere != "'"; caractere = lerProximoCaractere() {
			lexema += caractere
		}
		lexema += "'"

		return core.Token{Type: core.TIPO_CADEIA, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	return core.Token{}
}
