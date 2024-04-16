package vec

import "github.com/razor-87/mathx"

func Sum[T mathx.Vector[E], E mathx.Floaty](v T) (ret E) {
	const k = 8
	for ; len(v) >= k; v = v[k:] {
		ret += v[0] + v[1] + v[2] + v[3] + v[4] + v[5] + v[6] + v[7]
	}
	for i := range v {
		ret += v[i]
	}

	return ret
}
