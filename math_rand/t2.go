package main

/*
#cgo CFLAGS: -I${SRCDIR}/TestU01/include
#cgo LDFLAGS: ${SRCDIR}/TestU01/testu01/.libs/libtestu01.a ${SRCDIR}/TestU01/probdist/.libs/libprobdist.a ${SRCDIR}/TestU01/mylib/.libs/libmylib.a -lm

//#define SMALL

#include "unif01.h"
#include "bbattery.h"

#define BUFSZ (1<<20)

extern void goRandInt32Batch(void *);
extern void goRandDblBatch(void *);

static unsigned int goRandInt32(void) {
	static unsigned int buffer[BUFSZ];
	static int n = 0;
	if (n == 0) {
		goRandInt32Batch(buffer);
		n = BUFSZ;
	}
	return buffer[--n];
}
static double goRandDbl(void) {
	static double buffer[BUFSZ];
	static int n = 0;
	if (n == 0) {
		goRandDblBatch(buffer);
		n = BUFSZ;
	}
	return buffer[--n];
}

static int test(void) {
   unif01_Gen *gen;

   gen = unif01_CreateExternGen01 ("rand.Float64", goRandDbl);
#ifdef SMALL
   bbattery_SmallCrush (gen);
#else
   bbattery_BigCrush (gen);
#endif
   unif01_DeleteExternGen01 (gen);

   gen = unif01_CreateExternGenBits ("rand.Uint32", goRandInt32);
#ifdef SMALL
   bbattery_SmallCrush (gen);
#else
   bbattery_BigCrush (gen);
#endif
   unif01_DeleteExternGenBits (gen);

   return 0;
}
*/
import "C"
import "math/rand"
import "unsafe"

const BUFSZ = C.BUFSZ

//export goRandInt32Batch
func goRandInt32Batch(p unsafe.Pointer) {
	s := (*[BUFSZ]uint32)(p)[:]
	for i := range s {
		s[i] = rand.Uint32()
	}
}

//export goRandDblBatch
func goRandDblBatch(p unsafe.Pointer) {
	s := (*[BUFSZ]float64)(p)[:]
	for i := range s {
		s[i] = rand.Float64()
	}
}

func main() {
	C.test()
}
