a.out: e_log.c fbsd_log.c main.c openlibm_log.c
	gcc -O3 -o "$@" $^ -lm

.PHONY: clean
clean:
	rm -rf a.out *.o
