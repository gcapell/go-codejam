include $(GOROOT)/src/Make.inc

TARG=SavingTheUniverse
GOFILES=$(TARG).go

include $(GOROOT)/src/Make.cmd

test: $(TARG)
	./$(TARG) < sample.txt
