package main

import (
	"fmt"

	"github.com/hacknlove/aoc/helpers"
)

func primeraParte() {
	casas := make(map[string]int)
	cordenadas := [2]int{0, 0}

	cordenadasS := "0,0"
	casas["0,0"] = 1

	for char := range helpers.ForEachCharacter("input") {
		switch char {
		case ">":
			cordenadas[1]++
		case "<":
			cordenadas[1]--
		case "^":
			cordenadas[0]--
		case "v":
			cordenadas[0]++
		}
		cordenadasS = fmt.Sprintf("%d,%d", cordenadas[0], cordenadas[1])

		casas[cordenadasS] = 1
	}
	fmt.Println("casas:", len(casas))
}

func segundaParte() {
	casas := make(map[string]int)
	casas["0,0"] = 1

	cordenadas := [2][2]int{{0, 0}, {0, 0}}
	cordenadasS := "0,0"

	i := 0
	for char := range helpers.ForEachCharacter("input") {
		switch char {
		case ">":
			cordenadas[i][1]++
		case "<":
			cordenadas[i][1]--
		case "^":
			cordenadas[i][0]--
		case "v":
			cordenadas[i][0]++
		}
		cordenadasS = fmt.Sprintf("%d,%d", cordenadas[i][0], cordenadas[i][1])

		i = (i + 1) % 2

		casas[cordenadasS] = 1
	}
	fmt.Println("casas:", len(casas))

}

func main() {

	primeraParte()
	segundaParte()

}
