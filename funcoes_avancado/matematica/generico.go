package matematica

//Calculo executa qquer tipo de calculo
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador multiplica x por y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor - Divide x por y
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto - Divide  x por y retorna o resultado e o resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
