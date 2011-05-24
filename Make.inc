include $(GOROOT)/src/Make.inc

TARG=solver
test: $(TARG)
	./$(TARG) < sample.txt
