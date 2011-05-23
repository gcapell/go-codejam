include $(GOROOT)/src/Make.inc

TARG=TrainTimetable
GOFILES=$(TARG).go

include $(GOROOT)/src/Make.cmd

test: $(TARG)
	./$(TARG) < sample.txt
