package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hacknlove/aoc/helpers"
)

// func Split(s, sep string) []string

func parsearLinea(line string) []int {
	dimensionesString := strings.Split(line, "x")

	dimensiones := make([]int, 3)

	for i, s := range dimensionesString {
		d, _ := strconv.Atoi(s)
		dimensiones[i] = d
	}

	sort.Ints(dimensiones)
	return dimensiones
}

func papel(dimensiones []int) int {
	areas := make([]int, 3)
	areas[0] = int(dimensiones[0]) * int(dimensiones[1])
	areas[1] = int(dimensiones[0]) * int(dimensiones[2])
	areas[2] = int(dimensiones[2]) * int(dimensiones[1])

	return 3*areas[0] + 2*areas[1] + 2*areas[2]
}

func primeraParte() {

	papelTotal := 0

	for line := range helpers.ForEachLine("input") {
		papelTotal += papel(parsearLinea(line))
	}

	fmt.Println("papel total", papelTotal)
}

func cinta(dimensiones []int) int {
	return 2*(dimensiones[0]+dimensiones[1]) + dimensiones[0]*dimensiones[1]*dimensiones[2]
}

func segundaParte() {
	cintaTotal := 0

	for line := range helpers.ForEachLine("input") {
		cintaTotal += cinta(parsearLinea(line))
	}

	fmt.Println("cinta total", cintaTotal)
}

func main() {

	primeraParte()
	segundaParte()

}
