package hamming_test

import (
  "github.com/twmb/hamming"
  "testing"
  "time"
)

var uints6 []uint32
var uints7 []uint32

func init() {
  loopCount := 1
  uints6 = make([]uint32, 0, 906192 * loopCount)
  uints7 = make([]uint32, 0, 3365856 * loopCount)
  // loops at beginning to prevent caching of Set, SetK stuff
  // all uints6 have 6 bits set
  // all uints7 have 7 bits set
  for loops := 0; loops < loopCount; loops++ {
    for i := uint32(1); i < 1 << 32 - 1 && i != 0; i <<= 1 {
      for j := i << 1; j < 1 << 32 - 1 && j != 0; j <<= 1 {
        for k := j << 1; k < 1 << 32 - 1 && k != 0; k <<= 1 {
          for z := k << 1; z < 1 << 32 - 1 && z != 0; z <<= 1 {
            for l := z << 1; l < 1 << 32 - 1 && l != 0; l <<= 1 {
              for r := l << 1; r < 1 << 32 - 1 && r != 0; r <<= 1 {
                uints6 = append(uints6, i | j | k | z | l | r)
                for e := r << 1; e < 1 << 32 - 1 && e != 0; e <<= 1 {
                  uints7 = append(uints7, i | j | k | z | l | r | e)
                }
              }
            }
          }
        }
      }
    }
  }
}

func TestSet(t *testing.T) {
  for i := 0; i < 256; i++ {
    if hamming.Set(uint32(i)) != hamming.Distance(i, 0) {
      t.Errorf("Error in set: wanted %v for %v, got %v", 
      hamming.Distance(i, 0), i, hamming.Set(uint32(i)))
    }
  }
   testLargeSet(t)
}

func testLargeSet(t *testing.T) {
  count := 0
  start := time.Now()
  for _, i := range uints6 {
    count += hamming.Set(i)
  }
  t.Logf("Time to check Set on %v uints6: %v. Count of all bits set: %v",
  len(uints6), time.Since(start), count)

  kCount := 0
  start = time.Now()
  for _, i := range uints6 {
    kCount += hamming.SetK(i)
  }
  t.Logf("Time to check SetK on %v uints6: %v. Count of all bits set: %v",
  len(uints6), time.Since(start), kCount)

  sCount := uint32(0)
  start = time.Now()
  for _, i := range uints6 {
    sCount += hamming.SetP(i)
  }
  t.Logf("Time to check SetP on %v uints6: %v. Count of all bits set: %v",
  len(uints6), time.Since(start), sCount)

  if kCount != count {
    t.Errorf("expected equal counts, Set: %v, SetK: %v", count, kCount)
  }

  // dense

  count = 0
  start = time.Now()
  for _, i := range uints7 {
    count += hamming.Set(i)
  }
  t.Logf("Time to check Set on %v uints7: %v. Count of all bits set: %v",
  len(uints7), time.Since(start), count)

  kCount = 0
  start = time.Now()
  for _, i := range uints7 {
    kCount += hamming.SetK(i)
  }
  t.Logf("Time to check SetK on %v uints7: %v. Count of all bits set: %v",
  len(uints7), time.Since(start), kCount)

  sCount = uint32(0)
  start = time.Now()
  for _, i := range uints7 {
    sCount += hamming.SetP(i)
  }
  t.Logf("Time to check SetP on %v uints7: %v. Count of all bits set: %v",
  len(uints7), time.Since(start), sCount)
}

