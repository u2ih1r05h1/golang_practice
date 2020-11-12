GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean

.PHONY: all build run clean

all:
	make build
	make run
	make clean

build:
	$(GOBUILD) test.go

run:
	$(GORUN) test.go

clean:
	$(GOCLEAN)
