include $(GOROOT)/src/Make.inc

TARG=AlienLanguage
GOFILES=$(TARG).go

include $(GOROOT)/src/Make.cmd

test: $(TARG)
	./$(TARG) < sample.txt
