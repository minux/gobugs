a.out: e_log.c fbsd_log.c main.c openlibm_log.c
	gcc -O3 -o "$@" $^ -lm

.PHONY: clean bench
clean:
	rm -rf a.out *.o

N=500000000
bench: e_log.c fbsd_log.c main.c openlibm_log.c
	@gcc $^ -DBENCH -DTIMES=$(N) -DLOG=log -lm -O3
	@echo log in libm:
	@/usr/bin/time ./a.out
	@gcc $^ -DBENCH -DTIMES=$(N) -DLOG=__ieee754_log -lm -O3
	@echo fdlibm __ieee754_log:
	@/usr/bin/time ./a.out
	@gcc $^ -DBENCH -DTIMES=$(N) -DLOG=__fb_ieee754_log -lm -O3
	@echo FreeBSD __fb_ieee754_log:
	@/usr/bin/time ./a.out
	@gcc $^ -DBENCH -DTIMES=$(N) -DLOG=openlibm_log -lm -O3
	@echo openlibm_log:
	@/usr/bin/time ./a.out
	@rm -f a.out

