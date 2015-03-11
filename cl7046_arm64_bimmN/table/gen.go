package main

/*
#include <stdint.h>
#include <assert.h>
#include <errno.h>

static inline int countLeadingZeros(unsigned int x) {
	return __builtin_clz(x);
}

static inline uint64_t ror(uint64_t elt, unsigned size) {
  return ((elt & 1) << (size-1)) | (elt >> 1);
}

/// decodeLogicalImmediate - Decode a logical immediate value in the form
/// "N:immr:imms" (where the immr and imms fields are each 6 bits) into the
/// integer value it represents with regSize bits.
//static inline uint64_t decodeLogicalImmediate(uint64_t val, unsigned regSize) {
//  // Extract the N, imms, and immr fields.
//  unsigned N = (val >> 12) & 1;
//  unsigned immr = (val >> 6) & 0x3f;
//  unsigned imms = val & 0x3f;
uint64_t decodeLogicalImmediate(unsigned N, unsigned immr, unsigned imms, unsigned regSize) {
  errno = 0;
  N &= 1;
  immr &= 0x3f;
  imms &= 0x3f;

  assert((regSize == 64 || N == 0) && "undefined logical immediate encoding");
  int len = 31 - countLeadingZeros((N << 6) | (~imms & 0x3f));
  assert(len >= 0 && "undefined logical immediate encoding");
  unsigned size = (1 << len);
  unsigned R = immr & (size - 1);
  unsigned S = imms & (size - 1);
  //assert(S != size - 1 && "undefined logical immediate encoding");
  if (S == size - 1) {
  	errno = EINVAL;
	return 0;
  }
  uint64_t pattern = (1ULL << (S + 1)) - 1;
  unsigned i;
  for (i = 0; i < R; ++i)
    pattern = ror(pattern, size);

  // Replicate the pattern to fill the regSize.
  while (size != regSize) {
    pattern |= (pattern << size);
    size *= 2;
  }
  return pattern;
}
*/
import "C"
import (
	"fmt"
)

func main() {
	for _, rs := range []C.uint{32, 64} {
		for n := C.uint(0); n < rs/32; n++ {
			for immr := C.uint(0); immr < 1<<6-1; immr++ {
				for imms := C.uint(0); imms < 1<<6-1; imms++ {
					x, ok := C.decodeLogicalImmediate(n, immr, imms, rs)
					if ok == nil {
						fmt.Printf("%d: %d %02x %02x - %08x\n", rs, n, immr, imms, x)
					}
				}
			}
		}
	}
}
