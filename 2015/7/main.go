package main

import (
	"regexp"
	"strconv"

	"github.com/hacknlove/aoc/helpers"
)

var asignacion = regexp.MustCompile("^(.*) -> ([a-z]+)$")

var negacion = regexp.MustCompile("^NOT ([a-z0-9]+)$")
var operacion = regexp.MustCompile("^([a-z0-9]+) ([A-Z]+) ([0-9a-z]+)$")

var variables = make(map[string]string)
var valores = make(map[string]uint16)

func evaluar(expresion string) uint16 {

	if valor, ok := valores[expresion]; ok {
		return valor
	}

	if variable, ok := variables[expresion]; ok {
		valores[expresion] = evaluar(variable)
		return valores[expresion]
	}

	if v2, e := strconv.ParseUint(expresion, 10, 16); e == nil {
		valores[expresion] = uint16(v2)
		return valores[expresion]
	}

	comando := negacion.FindStringSubmatch(expresion)

	if len(comando) > 0 {
		valores[comando[1]] = evaluar(comando[1])
		return 0xFFFF ^ valores[comando[1]]
	}

	comando = operacion.FindStringSubmatch(expresion)
	if len(comando) > 0 {
		var valor1, valor2 uint16

		if v1, ok1 := valores[comando[1]]; ok1 {
			valor1 = v1
		} else {
			valor1 = evaluar(comando[1])
			valores[comando[1]] = valor1
		}
		if v2, ok2 := valores[comando[3]]; ok2 {
			valor2 = v2
		} else {
			valor2 = evaluar(comando[3])
			valores[comando[3]] = valor2
		}

		switch comando[2] {
		case "AND":
			return valor1 & valor2
		case "OR":
			return valor1 | valor2
		case "LSHIFT":
			return valor1 << valor2
		case "RSHIFT":
			return valor1 >> valor2
		}
	}
	panic("Aquí no")
}

func primeraParte() {

	for line := range helpers.ForEachLine("input") {
		comando := asignacion.FindStringSubmatch(line)
		variables[comando[2]] = comando[1]
	}

	println("Primera parte:", variables["a"], evaluar(variables["a"]))
}

func segundaParte() {
	a := evaluar(variables["a"])

	variables = make(map[string]string)
	valores = make(map[string]uint16)

	for line := range helpers.ForEachLine("input") {
		comando := asignacion.FindStringSubmatch(line)
		variables[comando[2]] = comando[1]
	}
	println(valores["b"])
	valores["b"] = a
	println(valores["b"])

	println("Segunda parte:", variables["a"], evaluar(variables["a"]))
}

func main() {
	primeraParte()
	segundaParte()
}
