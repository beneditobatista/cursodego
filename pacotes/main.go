package main

import (
	"fmt"

	"github.com/beneditobatista/cursodego/pacotes/operadora"
	"github.com/beneditobatista/cursodego/pacotes/prefixo"
)

//NomeDoUsuario Nome do usuario
var NomeDoUsuario = "Benedito"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)

}
