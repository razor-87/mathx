package gen

import "math/rand"

const seed = 42

var r = rand.New(rand.NewSource(seed)) //nolint:gosec
