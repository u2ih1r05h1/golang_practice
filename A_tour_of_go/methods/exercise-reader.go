package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(rb []byte) (int, error) {
	for i := range rb {
		rb[i] = 'A'
	}

	return len(rb), nil
}

func main() {
	reader.Validate(MyReader{})
}
