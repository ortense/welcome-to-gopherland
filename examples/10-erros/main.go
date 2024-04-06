package main

import (
	"exemplos/10-erros/ping"
	"fmt"
)

func main() {

	fmt.Println("Go não tem try/catch")
	// mas isso pode ser bom...

	url := "https://google.com.br"
	ping.Ping(url)
}
