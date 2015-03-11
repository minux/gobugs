package arm64

/*
#include <stdint.h>
#include <assert.h>
typedef int bool;
enum { false, true };

/// isMask_64 - This function returns true if the argument is a sequence of ones
/// starting at the least significant bit with the remainder zero (64 bit
/// version).
inline bool isMask_64(uint64_t Value) {
  return Value && ((Value + 1) & Value) == 0;
}

/// isShiftedMask_64 - This function returns true if the argument contains a
/// sequence of ones with the remainder zero (64 bit version.)
inline bool isShiftedMask_64(uint64_t Value) {
  return isMask_64((Value - 1) | Value);
}

static inline int countLeadingOnes(uint64_t x) {
	return __builtin_clzl(~x);
}
static inline int countTrailingZeros(uint64_t x) {
	return __builtin_ctzl(x);
}
static inline int countTrailingOnes(uint64_t x) {
	return __builtin_ctzl(~x);
}

/// processLogicalImmediate - Determine if an immediate value can be encoded
/// as the immediate operand of a logical instruction for the given register
/// size.  If so, return true with "encoding" set to the encoded value in
/// the form N:immr:imms.
static int processLogicalImmediate(uint64_t Imm, unsigned RegSize,
                uint64_t *Encoding) {
  if (Imm == 0ULL || Imm == ~0ULL ||
      (RegSize != 64 && (Imm >> RegSize != 0 || Imm == ~0U)))
    return false;

  // First, determine the element size.
  unsigned Size = RegSize;

  do {
    Size /= 2;
    uint64_t Mask = (1ULL << Size) - 1;

    if ((Imm & Mask) != ((Imm >> Size) & Mask)) {
      Size *= 2;
      break;
    }
  } while (Size > 2);

  // Second, determine the rotation to make the element be: 0^m 1^n.
  uint32_t CTO, I;
  uint64_t Mask = ((uint64_t)-1LL) >> (64 - Size);
  Imm &= Mask;

  if (isShiftedMask_64(Imm)) {
    I = countTrailingZeros(Imm);
    assert(I < 64 && "undefined behavior");
    CTO = countTrailingOnes(Imm >> I);
  } else {
    Imm |= ~Mask;
    if (!isShiftedMask_64(~Imm))
      return false;

    unsigned CLO = countLeadingOnes(Imm);
    I = 64 - CLO;
    CTO = CLO + countTrailingOnes(Imm) - (64 - Size);
  }

  // Encode in Immr the number of RORs it would take to get *from* 0^m 1^n
  // to our target value, where I is the number of RORs to go the opposite
  // direction.
  assert(Size > I && "I should be smaller than element size");
  unsigned Immr = (Size - I) & (Size - 1);

  // If size has a 1 in the n'th bit, create a value that has zeroes in
  // bits [0, n] and ones above that.
  uint64_t NImms = ~(Size-1) << 1;

  // Or the CTO value into the low bits, which must be below the Nth bit
  // bit mentioned above.
  NImms |= (CTO-1);

  // Extract the seventh bit and toggle it to create the N field.
  unsigned N = ((NImms >> 6) & 1) ^ 1;

  *Encoding = (N << 12) | (Immr << 6) | (NImms & 0x3f);
  return true;
}

uint64_t enc;
int f(long n, unsigned long x) {
	int i;
	for (i = 0; i < n; i++)
		processLogicalImmediate(x+i, 64, &enc);
}

int g(long n) { return n; }
*/
import "C"

func F(n uint64, x uint64) {
	C.f(C.long(n), C.ulong(x))
}

func G(n uint64) {
	C.g(C.long(n))
}
