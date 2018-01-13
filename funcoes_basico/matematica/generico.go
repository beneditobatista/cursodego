package matematica

//Calculo executa qquer tipo de calculo
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador multiplica x por y
func Multiplicador(x int, y int) int {
	return x * y
}
