target = bin2dec
obj = main.o convert.o

bin2dec: $(obj)
	cc -o $(target) $(obj)

main.o:
	cc -c main.c 

convert.o:
	cc -c convert.c convert.h

clean:
	-rm *.o *.h.gch
