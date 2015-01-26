#include <stdio.h>
#include <math.h>

extern double __ieee754_log(double);
extern double __fb_ieee754_log(double);
extern double openlibm_log(double);
int main() {
#ifdef BENCH
	volatile
#endif
	double x = 0.6023203277884887;
#ifndef BENCH
#define PRINT(f) printf("%20s = %.20g\n", #f, (f))
	printf("given x = %.20g:\n", x);
	PRINT(__ieee754_log(x));
	PRINT(__fb_ieee754_log(x));
	PRINT(openlibm_log(x));
	PRINT(log(x));
#else
	int i;
	volatile double y;
	for (i = 0; i < TIMES; i++) {
		y = LOG(x);
	}
#endif
}
