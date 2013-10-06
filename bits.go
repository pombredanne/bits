// Package bits has functions to compute hamming
// distance and set bits count.
package bits

var table [256]int = [256]int{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8}

// Hamming returns the hamming distance between two integers.
func Hamming(a, b int) int {
	x := uint(a ^ b)
	distance := 0
	for ; x > 0; x >>= 8 {
		distance += table[x&0xFF]
	}
	return distance
}

// SetTable returns the number of bits set on i.
// SetU* functions are faster if you can use unsigned integers.
func SetTable(x int) int {
	i := uint(x)
	c := 0
	for ; i > 0; i >>= 8 {
		c += table[i&0xFF]
	}
	return c
}

// SetKernighan returns the number of bits set on i.
// SetTable is usually faster if the average bits set is larger than 5.
func SetKernighan(i uint) int {
	c := 0
	for ; i != 0; c++ {
		i &= i - 1
	}
	return c
}

// SetU32 returns the number of bits set on i.
func SetU32(i uint32) uint32 {
	i = i - ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	return (((i + (i >> 4)) & 0x0F0F0F0F) * 0x01010101) >> 24
}

// SetU64 returns the number of bits set on i.
func SetU64(i uint64) uint64 {
	word := i
	word -= (word >> 1) & 0x5555555555555555
	word = (word & 0x3333333333333333) + ((word >> 2) & 0x3333333333333333)
	return (((word + (word >> 4)) & 0xf0f0f0f0f0f0f0f) * 0x101010101010101) >> 56
}
