// Package mathx provides useful mathematical functions and types.
package mathx

type Floaty interface{ ~float32 | ~float64 }

type Vector[T Floaty] interface{ ~[]T }
