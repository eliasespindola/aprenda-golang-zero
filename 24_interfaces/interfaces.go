package main

import "math"

// Vamos ter formas geometricas
type forma interface {
	area() float64
}

// Implementacao da interface, a implementacao é implicita
func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) float64 {
	return f.area()
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)
	c := circulo{10}
	escreverArea(c)
}
