package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

var input = "bgvyzdsv"

func minar(prefijo string) {
	respuesta := 1
	for {
		data := []byte(fmt.Sprintf("%s%d", input, respuesta))
		hash := fmt.Sprintf("%x", md5.Sum(data))
		for i := 0; i < 5; i++ {
			if strings.HasPrefix(hash, prefijo) {
				fmt.Println("respuesta", respuesta)
				return
			}
		}
		respuesta++
	}

}

func primeraParte() {
	minar("00000")
}

func segundaParte() {
	minar("000000")
}

func main() {
	primeraParte()
	segundaParte()
}
