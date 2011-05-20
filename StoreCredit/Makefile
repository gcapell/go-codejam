include $(GOROOT)/src/Make.inc

TARG=StoreCredit
GOFILES=StoreCredit.go

include $(GOROOT)/src/Make.cmd

test: $(TARG)
	./$(TARG) < sample.txt
