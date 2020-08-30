/*
	Pacote responsável por fazer a leitura de um determinado arquivo e disponibilizar para
	as demais partes do compilador.
*/
package main

import (
	"io/ioutil"
	"log"
)

// buffer - carrega todo os arquivo nessa variável para o leitor poder trabalhar.
var buffer string

// limiteBuffer - defini o tamanho do buffer.
var limiteBuffer int

// posicaoAtual - mostra em qual coluna o leitor está.
var posicaoAtual = -1

// linhaAtual - defini a linha em que o leitor está trabalhando.
var linhaAtual = 1

// colunaAtual - ponteiro para posicaoAtual que representa da coluna que o leitor está trabalhando.
var colunaAtual = &posicaoAtual

// lerArquivo - lê determinado arquivo e carrega no buffer.
func lerArquivo(file string) {
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	buffer = string(bytes)
	limiteBuffer = len(buffer)
}

// lerProximoCaractere - lê cada caracterer do buffer.
func lerProximoCaractere() string {
	if chegouNoLimiteDoBuffer() {
		return ""
	}

	avancarNoBuffer()
	return lerByteAtual()
}

// chegouNoLimiteDoBuffer - verifica se o leitor chegou no limite do buffer.
func chegouNoLimiteDoBuffer() bool {
	if posicaoAtual == limiteBuffer-1 {
		return true
	}

	return false
}

// avancarNoBuffer - avança a posição atual do leitor para frente.
func avancarNoBuffer() {
	posicaoAtual++
}

// avancarLinha - avança a posição atual da linha do leitor trabalha.
func avancarLinha() {
	linhaAtual++
}

// RetrocederNoBuffer - decrementa 1 na posição atual do buffer
func RetrocederNoBuffer() {
	if posicaoAtual < limiteBuffer-1 {
		posicaoAtual--
	}
}

// lerByteAtual - retorna a string que representa aquele byte.
func lerByteAtual() string {
	return string(buffer[posicaoAtual])
}
