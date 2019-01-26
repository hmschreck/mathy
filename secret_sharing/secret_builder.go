package secret_sharing

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const max = 1000
const min = -1000

func main() {
	rand.Seed(time.Now().Unix())
	secret, _ := strconv.Atoi(os.Args[3])
	n, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])
	p := NewPolynomial([]float64{})
	poly := make([]float64, 0)
	poly = append(poly, float64(secret))
	p.Coefficients = poly
	for i := 1; i < k; i++ {
		p.Coefficients = append(p.Coefficients, float64(rand.Intn(max-min)+min))
	}
	for i := 1; i <= n; i++ {
		fmt.Println(i, p.eval(float64(i)))
	}
	fmt.Println(p.Coefficients)
}
