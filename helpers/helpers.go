package helpers

import (
	"bufio"
	"os"
)

// ForEachCharacter devuelve el archivo caracter por caracter
func ForEachCharacter(path string) chan string {
	channel := make(chan string)

	go forEachCharacter(path, channel)

	return channel
}

func forEachCharacter(path string, channel chan string) {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		channel <- scanner.Text()
	}
	close(channel)
}

// ForEachLine devuelve el archivo linea por linea
func ForEachLine(path string) chan string {
	channel := make(chan string)

	go forEachLine(path, channel)

	return channel
}

func forEachLine(path string, channel chan string) {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		channel <- scanner.Text()
	}
	close(channel)
}
