package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

var numero = regexp.MustCompilePOSIX(`[-0-9]+`)

func primeraParte() {
	content, _ := ioutil.ReadFile("input")
	texto := string(content)

	numeros := numero.FindAllString(texto, -1)

	suma := 0
	for _, x := range numeros {
		v, _ := strconv.Atoi(x)
		suma += v
	}

	fmt.Println(suma)

}

func numeros(data interface{}) []int {
	respuesta := []int{}

	switch data := data.(type) {
	case float64:
		respuesta = append(respuesta, int(data))
	case []interface{}:
		for _, elemento := range data {
			respuesta = append(respuesta, numeros(elemento)...)
		}
	case map[string]interface{}:
		rojo := false

		for _, elemento := range data {
			if str, ok := elemento.(string); ok && str == "red" {
				rojo = true
				break
			}
		}

		if !rojo {
			for _, elemento := range data {
				respuesta = append(respuesta, numeros(elemento)...)
			}
		}
	}

	return respuesta
}

func segundaParte() {
	content, _ := ioutil.ReadFile("input")

	var data interface{}
	json.Unmarshal(content, &data)

	suma := 0
	for _, numero := range numeros(data) {
		suma += numero
	}

	println(suma)
}

func main() {
	primeraParte()
	segundaParte()
}
