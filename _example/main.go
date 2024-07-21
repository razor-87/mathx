package main

import (
	"fmt"

	"github.com/razor-87/mathx/vec"
)

func main() {
	// Regression model, pages 42-45 from https://web.stanford.edu/~boyd/vmls/vmls-slides.pdf
	β := []float32{148.73, -18.85} // n-vector β is the weight vector
	v := float32(54.40)            // scalar v is the offset
	regressors := [][]float32{     // x is a feature vector; its elements xi are called regressors (area, beds)
		{0.846, 1},
		{1.324, 2},
		{1.150, 3},
		{3.037, 4},
		{3.984, 5},
	}
	prices := []float32{115, 234.0, 198, 528, 572.5}     // scalar y is selling price of house in $1000
	predictions := vec.Zeros[[]float32](len(regressors)) // scalar yˆ is the prediction
	for i := range regressors {
		predictions[i] = vec.DotProd(regressors[i], β) + v // yˆ = xTβ + v
	}
	for i := range predictions {
		fmt.Printf("House #%d, Predicted price: %.2f, Actual price: %.2f\n", i+1, predictions[i], prices[i])
	}
}
