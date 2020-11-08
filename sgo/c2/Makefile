GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build

run:
	$(GORUN) hello.go

build:
	$(GOBUILD) -o hello hello.go

run_bin:
	make build
	./hello

symbols:
	readelf -s hello
