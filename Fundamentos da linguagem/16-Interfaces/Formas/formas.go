package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea (f forma){
	// minha forma deve ter um método area() que retorna um float64
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct{
	altura float64
	largura float64
}

func (r retangulo) area() float64{
	return r.altura * r.largura
}

func (c circulo) area() float64{
	return math.Pi * math.Pow(c.raio, 2)
}

type circulo struct{
	raio float64
}


func main(){

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

}