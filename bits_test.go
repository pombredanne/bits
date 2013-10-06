package bits_test

import (
	"github.com/twmb/bits"
	"testing"
)

func TestSet(t *testing.T) {
	for i := 0; i < 256; i++ {
		if int(bits.SetU32(uint32(i))) != bits.Hamming(i, 0) {
			t.Errorf("Error in set: wanted %v for %v, got %v",
				bits.Hamming(i, 0), i, bits.SetU32(uint32(i)))
		}
	}
}

func BenchmarkSetTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bits.SetTable(i)
	}
}
func BenchmarkSetKernighan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bits.SetKernighan(uint(i))
	}
}
func BenchmarkSetU32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bits.SetU32(uint32(i))
	}
}
func BenchmarkSetU64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bits.SetU64(uint64(i))
	}
}
