package main

import (
	"encoding/json"
	"fmt"

	"github.com/beneditobatista/cursodego/mapas/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25

	if err := casa.SetValor(60000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição de valor a casa: ", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor, por favor verifique a sua avaliação")
		}
		return
	}

	fmt.Printf("Casa é: %+v\r\n", casa)

	//objJSON, _ := json.Marshal(casa)
	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geracao do objeto JSON: ", err.Error())
		return
	}

	fmt.Println("A casa em JSON: ", string(objJSON))

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", imovel)
	}

	apto := model.Imovel{}

	apto.Nome = "Apto Azul"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(70000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, imovel)
	}

	_, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, "Casa Amarela")
	}

	fmt.Println("Quantos itens há no cache? ", len(cache))

}
