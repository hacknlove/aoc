package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/hacknlove/aoc/helpers"
)

func permutar(a []int, f func([]int)) {
	perm(a, f, 0)
}

func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

var parse = regexp.MustCompile("(.*) to (.*) = (.*)")

func main() {
	distancias := make(map[int]map[int]int)
	ciudades := make(map[string]int)
	indice := 0

	for line := range helpers.ForEachLine("input") {
		distancia := parse.FindStringSubmatch(line)

		ciudad1, ok := ciudades[distancia[1]]
		if !ok {
			ciudad1 = indice
			ciudades[distancia[1]] = indice
			indice++
		}

		ciudad2, ok := ciudades[distancia[2]]
		if !ok {
			ciudad2 = indice
			ciudades[distancia[2]] = indice
			indice++
		}

		v, _ := strconv.Atoi(distancia[3])
		if _, ok1 := distancias[ciudad1]; !ok1 {
			distancias[ciudad1] = make(map[int]int)
		}
		if _, ok2 := distancias[ciudad2]; !ok2 {
			distancias[ciudad2] = make(map[int]int)
		}
		distancias[ciudad1][ciudad2] = v
		distancias[ciudad2][ciudad1] = v
	}

	ruta := make([]int, len(ciudades))

	for i := range ruta {
		ruta[i] = i
	}

	minimo := math.MaxInt64
	maximo := 0
	permutar(ruta, func(arg1 []int) {
		fmt.Println(arg1)
		distancia := 0
		for i := 0; i < len(arg1)-1; i++ {
			fmt.Println(arg1[i], arg1[i+1], distancias[arg1[i]][arg1[i+1]])
			distancia += distancias[arg1[i]][arg1[i+1]]
		}
		if distancia < minimo {
			minimo = distancia
		}
		if distancia > maximo {
			maximo = distancia
		}
	})
	fmt.Println("Primera parte:", minimo)
	fmt.Println("Primera parte:", maximo)
}
