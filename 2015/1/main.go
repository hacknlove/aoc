package main

import (
	"fmt"

	"github.com/hacknlove/aoc/helpers"
)

func primeraParte() {
	piso := 0

	for char := range helpers.ForEachCharacter("input") {
		switch char {
		case "(":
			piso++
		case ")":
			piso--
		}
	}
	fmt.Println("piso:", piso)
}

func segundaParte() {
	piso := 0
	i := 0

	for char := range helpers.ForEachCharacter("input") {
		switch char {
		case "(":
			piso++
		case ")":
			piso--
		}
		i++
		if piso == -1 {
			fmt.Println("paso:", i)
			return
		}
	}
}

func main() {

	primeraParte()
	segundaParte()

}
