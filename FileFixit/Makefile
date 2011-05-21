include $(GOROOT)/src/Make.inc

TARG=FileFixit
GOFILES=$(TARG).go

include $(GOROOT)/src/Make.cmd

test: $(TARG)
	./$(TARG) < sample.txt
