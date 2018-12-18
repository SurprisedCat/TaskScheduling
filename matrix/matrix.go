package matrix

import (
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

//RandomMatrix Generate a N*N matrix of random float64 value(-99999~99999)
func RandomMatrix(dim int) *mat.Dense {
	rand.Seed(time.Now().Unix())
	// Generate a dim×dim matrix of random values.
	data := make([]float64, dim*dim)
	for i := range data {
		data[i] = -99999 + rand.Float64()*(199998)
	}
	return mat.NewDense(dim, dim, data)
}
