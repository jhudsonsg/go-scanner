package main

import "fmt"

func main() {
	lerArquivo("file")

	for token := proximoToken(); token.Type != ""; token = proximoToken() {
		fmt.Println(token)
	}
}
