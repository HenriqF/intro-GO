package vec2

type Vec2 struct {
	X int
	Y int
}

func Subir(vetor *Vec2, valor int) {
	vetor.X *= valor
	vetor.Y *= valor
}
