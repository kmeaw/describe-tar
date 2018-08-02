package main

import (
	"archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kr/pretty"
)

func main() {
	r := tar.NewReader(os.Stdin)
	for {
		h, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("% #v\n", pretty.Formatter(h))
		_, err = io.CopyN(ioutil.Discard, r, h.Size)
		if err != nil {
			panic(err)
		}
	}
}
