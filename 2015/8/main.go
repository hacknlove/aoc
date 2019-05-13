package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hacknlove/aoc/helpers"
)

var ascii = regexp.MustCompile(`\\x..`)

func primeraParte() {
	respuesta := 0

	for line := range helpers.ForEachLine("input") {
		respuesta += len(line)

		line = strings.Replace(line, `\\`, "_", -1)
		line = strings.Replace(line, `\"`, "_", -1)
		line = ascii.ReplaceAllString(line, "_")

		respuesta -= (len(line) - 2)
	}
	fmt.Println("Primera parte", respuesta)
}

func segundaParte() {
	respuesta := 0

	for line := range helpers.ForEachLine("input") {
		respuesta -= len(line)

		line = strings.Replace(line, `\`, "__", -1)
		line = strings.Replace(line, `"`, "__", -1)

		respuesta += len(line) + 2
	}
	fmt.Println("Segunda parte", respuesta)

}

func main() {
	primeraParte()
	segundaParte()
}
