package main

import "fmt"

var input = "hepxcrrq"

// verdadero si abc o efg ...
func secuenciaAscendente(password []byte) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i+1] == password[i]+1 && password[i+2] == password[i]+2 {
			return true
		}
	}
	return false
}

// falso si i o l
func sinProhibidas(password []byte) bool {
	for i := 0; i < len(password); i++ {
		if password[i] == 'i' || password[i] == 'o' || password[i] == 'l' {
			return false
		}
	}
	return true
}

// verdadero si aa..rr o ggww ...
func dosDobles(password []byte) bool {
	i := 0
	primero := false
	for ; i < len(password)-3; i++ {
		if password[i] == password[i+1] {
			primero = true
			i += 2
			break
		}
	}

	if !primero {
		return false
	}

	for ; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			return true
		}
	}
	return false
}

func incrementar(password []byte, i int) []byte {
	password[i] = password[i] + 1
	if password[i] > 'z' {
		password[i] = 'a'
		return incrementar(password, i-1)
	}
	return password
}
func esValida1(password []byte) bool {
	return secuenciaAscendente(password) && sinProhibidas(password) && dosDobles(password)
}

func siguiente(password string) string {
	array := []byte(password)
	for {
		array = incrementar(array, 7)
		if esValida1(array) {
			fmt.Println(string(array))
			return string(array)
		}
	}
}

func main() {
	siguiente(siguiente(input))
}
