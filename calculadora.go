package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		return operador1 + operador2
	case "-":
		return operador1 - operador2
	case "*":
		return operador1 * operador2
	default:
		log.Println(operador, "these operator is not valid! :D")
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada) // Atoi permite convertir valores String a Valores numericos
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

	fmt.Println(entrada)
	fmt.Println(operador)

	c := calc{}

	fmt.Println(c.operate(entrada, operador))

}
