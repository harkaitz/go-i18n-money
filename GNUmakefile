.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=go-spanish-currency
VERSION=1.0.0
PREFIX=/usr/local

## -- BLOCK:go --
.PHONY: all-go install-go clean-go $(BUILDDIR)/parse-monetary$(EXE)
all: all-go
install: install-go
clean: clean-go
all-go: $(BUILDDIR)/parse-monetary$(EXE)
install-go:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp  $(BUILDDIR)/parse-monetary$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(BUILDDIR)/parse-monetary$(EXE)
##
$(BUILDDIR)/parse-monetary$(EXE): $(GO_DEPS)
	mkdir -p $(BUILDDIR)
	go build -o $@ $(GO_CONF) ./cmd/parse-monetary
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: README.md LICENSE
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp README.md LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
