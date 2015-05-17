package main

/*
#cgo CFLAGS: -I${SRCDIR}/TestU01/include
#cgo LDFLAGS: ${SRCDIR}/TestU01/testu01/.libs/libtestu01.a ${SRCDIR}/TestU01/probdist/.libs/libprobdist.a ${SRCDIR}/TestU01/mylib/.libs/libmylib.a -lm

#include "unif01.h"
#include "bbattery.h"

extern unsigned int goRandInt32(void);
extern double goRandDbl(void);

static int test(void) {
   unif01_Gen *gen;

   gen = unif01_CreateExternGen01 ("rand.Float64", goRandDbl);
   bbattery_SmallCrush (gen);
   unif01_DeleteExternGen01 (gen);

   gen = unif01_CreateExternGenBits ("rand.Uint32", goRandInt32);
   bbattery_SmallCrush (gen);
   unif01_DeleteExternGenBits (gen);

   return 0;
}
*/
import "C"
import "math/rand"

//export goRandInt32
func goRandInt32() uint32 {
	return rand.Uint32()
}

//export goRandDbl
func goRandDbl() float64 {
	return rand.Float64()
}

func main() {
	C.test()
}
