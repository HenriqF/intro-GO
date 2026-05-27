package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("texto.txt")
	cont := string(bytes)
	var resultado string = ""
	if len(bytes) <= 0 {
		cont += "1"
	}

	c := 0
	analise := cont[0]

	for i := 0; i < len(cont); i++ {
		if cont[i] != analise {
			resultado += fmt.Sprintf("%v%c", c, analise)
			c = 1
			analise = cont[i]
			continue
		}
		c++
	}
	resultado += fmt.Sprintf("%v%c", c, analise)

	os.WriteFile("texto.txt", []byte(resultado), 0666)
}
