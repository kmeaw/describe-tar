package main

import (
	"archive/tar"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	xattrs := false
	flag.BoolVar(&xattrs, "xattrs", xattrs, "Show only files with non-empty xattrs")
	flag.Parse()

	j := json.NewEncoder(os.Stdout)
	j.SetIndent("", "    ")
	r := tar.NewReader(os.Stdin)
	for {
		h, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if !xattrs || h.Xattrs != nil {
			j.Encode(h)
		}
		_, err = io.CopyN(ioutil.Discard, r, h.Size)
		if err != nil {
			panic(err)
		}
	}
}
