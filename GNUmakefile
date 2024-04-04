.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=go-spanish-currency
VERSION=1.0.0
PREFIX=/usr/local

## -- BLOCK:go --
build/parse-monetary$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/parse-monetary
all: build/parse-monetary$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/parse-monetary$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/parse-monetary$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
