package pow

import "testing"

var vfpow10SC = []float64{
	2,
	3,
}

var pow10SC = []float64{
	1024,
	59049,
}

func alike(expected, result float64) bool {
	return expected == result
}

func TestPow10(t *testing.T) {
	for i := 0; i < len(vfpow10SC); i++ {
		if f := Pow10(vfpow10SC[i]); !alike(pow10SC[i], f) {
			t.Errorf("Pow10(%f) = %g, want %g", vfpow10SC[i], f, pow10SC[i])
		}
	}
}

func BenchmarkPow10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Pow10(10)
	}
}
