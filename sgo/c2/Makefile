GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

run:
	$(GORUN) hello.go

build:
	$(GOBUILD) -o hello hello.go

run_bin:
	make build
	./hello
	
clean:
	$(GOCLEAN)

symbols:
	readelf -s hello
