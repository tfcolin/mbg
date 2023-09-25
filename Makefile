PREFIX := /usr/local
SOFTNAME := sp

OBJLIB := lib$(SOFTNAME).a
OPTOPT := -O0

all : $(OBJLIB) 

CC := gcc
LINKER := gcc

CFLAGS := -g $(OPTOPT) -I. 
LFLAGS := -g $(OPTOPT) 

MAKE_DEP := gcc -MM

TESTS := sp_cs sp_gtk

SRCS := $(filter-out $(addsuffix .c, $(TESTS)), $(wildcard *.c))
DEPS := $(patsubst %.c, %.d, $(wildcard *.c))
OBJS := $(patsubst %.c, %.o, $(SRCS))

$(OBJLIB) : $(OBJLIB)($(OBJS))

sinclude $(DEPS)

%.d : %.c
	@$(MAKE_DEP) $^ 2>/dev/null | sed 's/\($*\)\.o[ :]*/\1.o $@ :/g' > $@

sp_gtk.o : sp_gtk.c
	$(CC) -c $(CFLAGS) `pkg-config --cflags gtk+-3.0` $< -o $@

sp_gtk : sp_gtk.o
	$(LINKER) $(LFLAGS) $< -o $@ -L. -l$(SOFTNAME) `pkg-config --libs gtk+-3.0` -lm

sp_cs.o : sp_cs.c
	$(CC) -c $(CFLAGS) $< `pkg-config --cflags ncursesw` -o $@

sp_cs : sp_cs.o
	$(LINKER) $(LFLAGS) $< -o $@ -L. -l$(SOFTNAME) `pkg-config --libs ncursesw` -lm

$(TESTS) : $(OBJLIB)

$(SRCS) *.h : Makefile
	@touch $@

clean:  
	rm -rf *.o $(OBJLIB) *.d $(TESTS) 

install:
	rm -rf $(PREFIX)/lib/$(OBJLIB)
	rm -rf $(PREFIX)/include/$(SOFTNAME)/
	cp -i $(OBJLIB) $(PREFIX)/lib/
	mkdir -p $(PREFIX)/include/$(SOFTNAME)/
	cp -i *.h $(PREFIX)/include/$(SOFTNAME)/

dist:
	cd ..; tar cvf $(SOFTNAME)-0.0.8.tar $(SOFTNAME)/

.PHONY: all clean install 

