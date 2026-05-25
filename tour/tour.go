package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
	"tour/vec2"
	"unsafe"
)

func fibas(n int) int {
	if n <= 1 {
		return n
	}
	return fibas(n-1) + fibas(n-2)
}

func seloko(x, y string) (string, string) {
	return y, x
}

func retornosigma(a int) (a1, a2 int) {
	a1 = a * 30
	a2 = a % a1
	return
}

func testesigma(a int) {
	fmt.Println("não é mentira: abri o arquivo", a)
	defer fmt.Println("fechei o arquivo", a)

	fmt.Println("conteudo do arquivo: ", a+20)
	return
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Sigma da bahia %v sigma %v", rand.Intn(100), math.Pi)

	for i := range 20 {
		fmt.Printf("n: %v\n", fibas(i))
	}

	var x, y string = seloko("da bahia", "sigma")
	//x, y := selok..
	fmt.Printf("%s %s\n", x, y)

	a, b := retornosigma(30)
	fmt.Printf("%d %d\n", a, b)
	fmt.Println(retornosigma(30))

	var c, python, java = 20, true, "merda"
	var i int = 1
	fmt.Println(i, c, python, java)

	sigma, da, bahia := "porra", "loca", 20
	fmt.Println(sigma, da, bahia)

	var (
		_albedo1 = 'a'
		_albedo2 = 'b'
		_albedo3 = 'c'
	)

	fmt.Printf("porraloca: %T, %d %c\n", _albedo1, _albedo2, _albedo3)

	var (
		num1 int64   = 20
		num2 float64 = 30.3
	)

	fmt.Printf("porraloca: %v\n", float64(num1)/num2)

	var (
		selokocompensa complex128 = 2 + 9i //????/
	)

	fmt.Println(selokocompensa)

	const memudaduvido int = 30

	var soma = 0
	for {
		if soma == 20 {
			break
		}
		soma++
	}
	fmt.Println("k")

	var valor = 30
	if valor := 20; valor <= 30 {
		fmt.Println("dadsdsad")
	}
	fmt.Println(valor) //bizarry

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("liinuz")
	default:
		fmt.Println(os)
	}

	dia_da_semana := time.Now().Weekday()

	fmt.Print(dia_da_semana)
	fmt.Printf(" %v/%v/%v\n", time.Now().Day(), int(time.Now().Month()), time.Now().Year())
	fmt.Printf("%v/%v/%v\n", time.Now().Day(), time.Now().Month(), time.Now().Year()) //bizarry

	testesigma(30)

	valornumerico := 30
	var ponteiro *int = &valornumerico
	*ponteiro = 1023941
	fmt.Println(valornumerico)

	var ponteirolegal unsafe.Pointer = unsafe.Pointer(&valornumerico)

	fmt.Println(&ponteirolegal, *(*int)(ponteirolegal))
	ponteirolegal = unsafe.Add(ponteirolegal, 1)
	fmt.Println(&ponteirolegal, *(*int)(ponteirolegal))

	var coordenadas vec2.Vec2 = vec2.Vec2{X: 1, Y: 0}

	fmt.Printf("%v %T\n", coordenadas, coordenadas)
	vec2.Subir(&coordenadas, 5)
	fmt.Printf("%v %T\n", coordenadas, coordenadas)

}
