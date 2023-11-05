PROJECT=go-spanish-currency
VERSION=1.0.0
PREFIX=/usr/local
all:
clean:
install:

## -- BLOCK:go --
all: all-go
install: install-go
clean: clean-go
deps: deps-go

build/parse-monetary$(EXE): deps
	go build -o $@ $(GO_CONF) ./cmd/parse-monetary

all-go:  build/parse-monetary$(EXE)
deps-go:
	mkdir -p build
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp  build/parse-monetary$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f  build/parse-monetary$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE  $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
