package cmd

import (
	"io"
	"io/ioutil"
)

func getBodyData(r io.Reader) (d []byte) {
	d, _ = ioutil.ReadAll(r)
	return d
}
