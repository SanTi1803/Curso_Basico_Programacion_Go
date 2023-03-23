package mypackage

import "fmt"

type Figuras2D interface {
	area() float64
}

type Cuadrado struct {
	Base float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (c Cuadrado) area() float64 {
	return c.Base * c.Base
}

func (r Rectangulo) area() float64 {
	return r.Base * r.Altura
}

func Calcular(f Figuras2D) {
	fmt.Println("Area:", f.area())
}
