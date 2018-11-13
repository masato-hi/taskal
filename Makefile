GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
BINDIR=./bin
DISTDIR=./dist
TMPDIR=./tmp
BINARY_NAME=taskal
BINARY_UNIX=$(BINARY_NAME)_unix
COVERFILE=./tmp/cover.out

all: build

build: fmt
	$(GOBUILD) -o "$(BINDIR)/$(BINARY_NAME)" -v

build-release: test
	$(GOBUILD) -o "$(DISTDIR)/$(BINARY_NAME)" -v -tags=release

test: fmt
	$(GOTEST) -tags debug -v -coverprofile=$(COVERFILE) ./...

fmt:
	$(GOFMT)

cover:
	$(GOTOOL) cover -func=$(COVERFILE)

clean:
	$(GOCLEAN)
	rm -f $(TMPDIR)/*
	rm -f $(DISTDIR)/*

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o "$(DISTDIR)/$(BINARY_UNIX)" -v


