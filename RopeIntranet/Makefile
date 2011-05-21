include $(GOROOT)/src/Make.inc

TARG=RopeIntranet
GOFILES=$(TARG).go

include $(GOROOT)/src/Make.cmd

test: $(TARG)
	./$(TARG) < sample.txt
