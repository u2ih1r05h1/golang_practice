# golang_practice

## env

on Ubuntu 20.04 LTS

```
sudo apt install golang
```

or

```
wget https://golang.org/dl/go1.15.4.linux-amd64.tar.gz
```

## build
```
$ go build foo.go
```

## run
```
$ go run foo.go
```
or
```
$ ./foo
```

## clean

```
$ go clean
```

## refactoring
```
$ gofmt foo.go

go fmt <=> gofmt -l -w
```

## cannot find package
```
go get -u bar/foo/bar/foo
```
