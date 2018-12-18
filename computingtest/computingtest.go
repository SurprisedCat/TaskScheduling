package main

import (
	"fmt"
	"time"

	"../matrix"
	"gonum.org/v1/gonum/mat"
)

func main() {
	iter := 100
	dim := 100
	// Initialize two matrices, a and b.

	a := matrix.RandomMatrix(dim)
	b := matrix.RandomMatrix(dim)

	// Take the matrix product of a and b and place the result in c.
	var c mat.Dense
	start := time.Now()
	for i := 0; i < iter; i++ {
		c.Mul(a, b)
	}
	cost := time.Since(start).Nanoseconds() / 1000
	fmt.Println(cost)
	// Print the result using the formatter.
	//fc := mat.Formatted(&c, mat.Prefix("    "), mat.Squeeze())
	//fmt.Printf("c = %v\n", fc)
	enc, _ := c.MarshalBinary()
	fmt.Println(enc)
	var dec mat.Dense
	dec.UnmarshalBinary(enc)
	// Print the result using the formatter.
	fc := mat.Formatted(&dec, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("dec = %v\n", fc)
}
