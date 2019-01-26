package secret_sharing

import (
	"fmt"
	"log"
)



func main() {
	point1 := Point{x: 1, y: 22267}
	point2 := Point{x: 2, y: 20601}
	point3 := Point{x: 3, y: 18125}
	points := []Point{point1, point2, point3}
	output := interpolate(points)
	fmt.Println(len(output.Coefficients))
	fmt.Println(output.eval(0))
	fmt.Println(output.Coefficients)
}

func removeDupes (input []float64) (output []float64) {
	encountered := map[float64]bool{}
	for v := range input {
		if encountered[input[v]] == true {
		} else {
			encountered[input[v]] = true
			 output = append(output, input[v])
		}
	}
	return
}
func interpolate(points []Point) (output Polynomial) {
	output.Coefficients = []float64{}
	fmt.Println(len(points))
	if len(points) == 0 {
		log.Fatal("Must provide at least one point")
	}
	x_values := []float64{}
	for _, value := range points {
		x_values = append(x_values, value.x)
	}
	if len(removeDupes(x_values)) != len(x_values) {
		log.Fatal("Not actually a function!")
	}
	terms := []Polynomial{}
	for i := 0; i < len(points); i++ {
		terms = append(terms, single_term(points, i))
		fmt.Println(terms)
	}
	for _, term := range terms {
		output = output.add(term)
	}
	return
}

func single_term(points []Point, i int) (output Polynomial){
	term := NewPolynomial([]float64{1.0})
	xi, yi := points[i].x, points[i].y
	for j, point := range points {
		if j == i {
			continue
		}
		xj := point.x
		term = term.mult(NewPolynomial([]float64{ (float64(-xj) / float64(xi-xj)), float64(1)/float64(xi-xj)}))
	}
	output = term.mult(NewPolynomial([]float64{yi}))
	return
}