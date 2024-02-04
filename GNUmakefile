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
deps:
build/parse-monetary$(EXE): deps
	@mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/parse-monetary
all-go: build/parse-monetary$(EXE)
install-go:
	install -D -t $(DESTDIR)$(PREFIX)/bin build/parse-monetary$(EXE)
clean-go:
	rm -f build/parse-monetary$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	install -D -t $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT) LICENSE
## -- BLOCK:license --
