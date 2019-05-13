package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "1321131112"

func contarRepeticiones(secuencia []int, i int) int {
	contar := 0
	numero := secuencia[i]
	for secuencia[i] == numero {
		i++
		contar++
		if i == len(secuencia) {
			break
		}
	}
	return contar
}

func expandir(secuencia []int) []int {
	var respuesta []int
	for i := 0; i < len(secuencia); {
		c := contarRepeticiones(secuencia, i)
		respuesta = append(respuesta, c)
		respuesta = append(respuesta, secuencia[i])
		i += c
	}
	return respuesta
}

func primeraParte() {
	var secuencia []int
	for _, letra := range strings.Split(input, "") {
		v, _ := strconv.Atoi(letra)
		secuencia = append(secuencia, v)
	}

	for i := 0; i < 40; i++ {
		secuencia = expandir(secuencia)
		fmt.Println(secuencia)
		fmt.Println(len(secuencia))
	}
}

func segundaParte() {
	var secuencia []int
	for _, letra := range strings.Split(input, "") {
		v, _ := strconv.Atoi(letra)
		secuencia = append(secuencia, v)
	}

	for i := 0; i < 50; i++ {
		secuencia = expandir(secuencia)
		fmt.Println(secuencia)
		fmt.Println(len(secuencia))
	}
}

func main() {
	primeraParte()
	segundaParte()
}
