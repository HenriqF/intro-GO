package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argv := os.Args
	if len(argv) != 3 {
		return
	}

	num, err := strconv.Atoi(argv[2])

	if err != nil {
		fmt.Printf("Erro: '%v' numero inválido", argv[2])
		return
	}
	runes := []rune(argv[1])
	for i := 0; i < len(runes); i++ {
		runes[i] += rune(num)
	}
	resultado := string(runes)

	fmt.Println(resultado)
}
