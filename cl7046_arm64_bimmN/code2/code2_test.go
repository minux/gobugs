package arm64

import "testing"

func BenchmarkNonLUT(b *testing.B) {
	F(uint64(b.N), 0x0ffffffffffffffe)
}
func BenchmarkCgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		G(uint64(b.N))
	}
}
