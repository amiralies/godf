PRGM = godf
BINDIR ?= /usr/bin

build:
	make clean
	go build

clean:
	rm -rf $(PRGM)

install:
	install -m 755 $(PRGM) $(BINDIR)

uninstall:
	rm -rf $(BINDIR)/$(PRGM)

.PHONY: build clean install uninstall

