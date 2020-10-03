package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(rb []byte) (int, error) {
	n, e := r.r.Read(rb)
	if e == nil {
		for i, v := range rb {
			switch {
			case 'A' <= v && v <= 'Z':
				rb[i] = (v-'A'+13)%26 + 'A'
			case 'a' <= v && v <= 'z':
				rb[i] = (v-'a'+13)%26 + 'a'
			}
		}
	}

	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
