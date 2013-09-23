package hamming_test

import (
	"github.com/twmb/hamming"
	"testing"
)

func TestDistance(t *testing.T) {
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			if hamming.Distance(i, j) != hamming.DistanceBits(uint64(i), uint64(j)) {
				t.Errorf("wrong")
			}
			if hamming.Distance(i, j) != hamming.HammingDistance(uint64(i), uint64(j)) {
				t.Errorf("wrong")
			}
		}
	}
}

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := 0; k < 200; k++ {
			for j := 7000; j < 8000; j++ {
				hamming.Distance(k, j)
			}
		}
	}
}

func BenchmarkHamDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := uint64(0); k < 200; k++ {
			for j := uint64(7000); j < 8000; j++ {
				hamming.HammingDistance(k, j)
			}
		}
	}
}

func BenchmarkDistanceBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := uint64(0); k < 200; k++ {
			for j := uint64(7000); j < 8000; j++ {
				hamming.DistanceBits(k, j)
			}
		}
	}
}
