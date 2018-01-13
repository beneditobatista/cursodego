package main

import (
	"fmt"
)

var cache map[string]string

func main() {

	cache = make(map[string]string, 0)

	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1
	fmt.Println("Qual a capacidade deste array? ", len(teste))
	fmt.Printf("O que há nesse Array? \n\r%v\r\n", teste)

	cantores := [2]string{"Michael Jackson", "John Lennon"}
	fmt.Println("Qual a capacidade deste array? ", len(cantores))
	fmt.Printf("O que há nesse Array? \n\r%v\r\n", cantores)

	paises := [...]string{"Portugal", "Brasil", "Angola", "Moçambique"}
	capitais := [...]string{"Lisboa", "Brasilia", "Luanda", "Maputo"}
	fmt.Println("Qual a capacidade deste array? ", len(capitais))
	for indice, cidade := range capitais {
		cache[paises[indice]] = capitais[indice]
		fmt.Printf("Capital[%d] é %s\n\r", indice, cidade)
	}

	fmt.Printf("A Capital do Brasil é  %s\n\r", cache["Brasil"])

}
