# golang_practice

## env

Version 1.15.4 or others on Ubuntu 20.04 LTS

```
$ sudo apt install golang
```

or

```
$ wget https://golang.org/dl/go1.15.4.linux-amd64.tar.gz
$ tar xf go1.15.4.linux-amd64.tar.gz
$ sudo mv go /usr/local/go1.15.4
$ sudo ln -s /usr/local/go1.15.4 /usr/local/go
$ /usr/local/go/bin/go version
$ cp -i .profile .profile.org
$ echo "PATH=\"/usr/local/go/bin:\$PATH\"" >> .profile
$ echo "export PATH" >> .profile
$ go version
$ mkdir ~/go
$ echo "GOPATH=\"\$HOME/go\"" >> .profile
$ echo "export GOPATH" >> .profile
$ env | grep GOPATH
```

### gocode

```
$ go get github.com/nsf/gocode
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
