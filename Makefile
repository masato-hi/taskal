GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
BINDIR=./bin
DISTDIR=./dist
BINARY_NAME=taskal
BINARY_UNIX=$(BINARY_NAME)_unix
COVERFILE=./tmp/cover.out

all: build

build:
	$(GOFMT)
	$(GOBUILD) -o "$(BINDIR)/$(BINARY_NAME)" -v

build-release:
	$(GOFMT)
	$(GOBUILD) -o "$(DISTDIR)/$(BINARY_NAME)" -v -tags=release

test:
	$(GOFMT)
	$(GOTEST) -tags debug -v -coverprofile=$(COVERFILE) ./...

cover:
	$(GOTOOL) cover -func=$(COVERFILE)

clean:
	$(GOCLEAN)
	rm -f "$(DISTDIR)/*"

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o "$(DISTDIR)/$(BINARY_UNIX)" -v


