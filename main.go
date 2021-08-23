package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(estrada string, operador string) int {

	entradaLimpia := strings.Split(estrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {

	case "+":
		return operador1 + operador2

	case "-":
		return operador1 - operador2

	case "/":
		return operador1 / operador2

	case "*":
		return operador1 * operador2

	default:
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	fmt.Println("entrada", entrada, "operador", operador)
	c := calc{}
	fmt.Println(c.operate(entrada, operador))
}
