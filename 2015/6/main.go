package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hacknlove/aoc/helpers"
)

type cordenada struct {
	x int
	y int
}

func parseCordenadas(texto string) cordenada {
	s := strings.Split(texto, ",")
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])

	return cordenada{x, y}
}

func toggle(grid *[][]bool, from cordenada, to cordenada) int {
	luces := 0

	g := *grid

	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			g[x][y] = !g[x][y]
			if g[x][y] {
				luces++
			} else {
				luces--
			}
		}
	}
	return luces
}

func turn(grid *[][]bool, modo string, from cordenada, to cordenada) int {
	luces := 0

	g := *grid

	status := false
	inc := -1

	if modo == "on" {
		status = true
		inc = 1
	}

	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			if g[x][y] != status {
				luces += inc
			}
			g[x][y] = status
		}
	}
	return luces
}

func primeraParte() {
	fmt.Println("Primera parte")
	luces := 0
	grid := make([][]bool, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]bool, 1000)
	}
	for line := range helpers.ForEachLine("input") {
		fields := strings.Fields(line)

		switch fields[0] {
		case "toggle":
			luces += toggle(&grid, parseCordenadas(fields[1]), parseCordenadas(fields[3]))
		case "turn":
			luces += turn(&grid, fields[1], parseCordenadas(fields[2]), parseCordenadas(fields[4]))
		}
	}
	fmt.Println("respuesta", luces)
}

func toggle2(grid *[][]int, from cordenada, to cordenada) int {
	luces := 0

	g := *grid

	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			g[x][y] += 2
			luces += 2
		}
	}
	return luces
}

func turn2(grid *[][]int, modo string, from cordenada, to cordenada) int {
	luces := 0

	g := *grid

	inc := -1

	if modo == "on" {
		inc = 1
	}

	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			if g[x][y]+inc < 0 {
				continue
			}
			g[x][y] += inc
			luces += inc
		}
	}
	return luces
}

func segundaParte() {
	fmt.Println("Primera parte")
	luces := 0
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}
	for line := range helpers.ForEachLine("input") {
		fields := strings.Fields(line)

		switch fields[0] {
		case "toggle":
			luces += toggle2(&grid, parseCordenadas(fields[1]), parseCordenadas(fields[3]))
		case "turn":
			luces += turn2(&grid, fields[1], parseCordenadas(fields[2]), parseCordenadas(fields[4]))
		}
	}
	fmt.Println("respuesta", luces)
}

func main() {
	fmt.Println("día 6")
	primeraParte()
	segundaParte()
}
