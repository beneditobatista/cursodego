package main

import "fmt"
import "github.com/beneditobatista/cursodego/funcoes_avancado/matematica"

func main() {

	resultado := matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicacao foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Soma, 5, 4)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("O resultado da Divisao foi: %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("O resultado da Divisao foi: %d e o resto: %d\r\n", resultado, resto)

}
