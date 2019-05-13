package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hacknlove/aoc/helpers"
)

func dobleLetra(line string) bool {
	anterior := ' '
	for _, letra := range line {
		if anterior == letra {
			return true
		}
		anterior = letra
	}
	return false
}

func primeraParte() {
	nice := 0
	//
	vocales := regexp.MustCompile("[a,e,i,o,u]")

	prohibido := regexp.MustCompile("ab|cd|pq|xy")

	for line := range helpers.ForEachLine("input") {
		if len(vocales.FindAllStringIndex(line, 3)) < 3 {
			continue
		}

		if prohibido.FindString(line) != "" {
			continue
		}

		if !dobleLetra(line) {
			continue
		}

		nice++
	}
	fmt.Println("respuesta", nice)
}

func abxxxab(line string) bool {
	for i := 0; i < len(line)-3; i++ {
		j := strings.LastIndex(line, line[i:i+2])
		if i+1 < j {
			return true
		}
	}
	return false
}

func axa(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func segundaParte() {
	nice := 0
	for line := range helpers.ForEachLine("input") {
		if !axa(line) {
			continue
		}

		if !abxxxab(line) {
			continue
		}

		nice++
	}
	fmt.Println("respuesta", nice)

}

func main() {

	primeraParte()
	segundaParte()

}
