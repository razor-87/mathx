package vec

import "github.com/razor-87/mathx"

func Sum[T mathx.Vector[E], E mathx.Floaty](xs T) (ret E) {
	const k = 8
	if r := len(xs) % k; r != 0 {
		for i := range xs[:r] {
			ret += xs[i]
		}
		xs = xs[r:]
	}
	var s T
	for i := 0; i < len(xs); i += k {
		s = xs[i : i+k : i+k]
		ret += s[0] + s[1] + s[2] + s[3] + s[4] + s[5] + s[6] + s[7]
	}

	return ret
}
