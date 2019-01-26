package secret_sharing

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Polynomial struct {
	Coefficients []float64
}

func (p *Polynomial) String() (output string){

	return
}

func NewPolynomial(coefficients []float64) (p Polynomial){
	if len(coefficients) == 0 {
		return
	}
	p.Coefficients = coefficients
	return
}

func strip(list []float64, remove int) (output []float64) {
	if len(list) == 0 {
		return
	}
	i := 0
	for i = len(list)-1; i >= 0; i-- {
		if (list[i] != float64(remove)) {
			break
		}
	}
	return list[0:i+1]
}

func (p *Polynomial) mult(mult Polynomial) (output Polynomial){
	newCoefficients := make([]float64, len(p.Coefficients)+len(mult.Coefficients))
	for i, a := range p.Coefficients {
		for j, b := range mult.Coefficients {
			newCoefficients[i+j] += a*b
		}
	}
	output.Coefficients = strip(newCoefficients, 0)
	return
}

func (p *Polynomial) add(addend Polynomial) (output Polynomial) {
	long := []float64{}
	fmt.Println("P: ", len(p.Coefficients))
	short := []float64{}
	if len(p.Coefficients) > len(addend.Coefficients) {
		long = p.Coefficients
		short = addend.Coefficients
	} else {
		long = addend.Coefficients
		short = p.Coefficients
	}
	for i, val := range short {
		long[i] += val
	}
	output.Coefficients = long
	return
}

func (p *Polynomial) eval(x float64) (output float64) {
	for deg := range p.Coefficients {
		output += p.Coefficients[deg] * math.Pow(float64(x), float64(deg))
	}
	return
}

