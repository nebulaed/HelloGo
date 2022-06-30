package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rd *rot13Reader) Read(b []byte) (int, error) {
	n, err := rd.r.Read(b)
	if err == io.EOF {
		return 0, err
	}
	for id := 0; id < n; id++ {
		if curr := &b[id]; (65 <= *curr && *curr <= 77) || (97 <= *curr && *curr <= 109) {
			*curr += 13
		} else if (78 <= *curr && *curr <= 90) || (110 <= *curr && *curr <= 122) {
			*curr -= 13
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
