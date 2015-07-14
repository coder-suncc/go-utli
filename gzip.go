package util

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
)

func GzipFile(src, dst string) error {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	n, err := w.Write(data)
	w.Close()
	f, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, &b)
	if err != nil {
		return err
	}
	return nil
}

func GzipByte(src) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err := w.Write(src)
	if err != nil {
		return nil, err
	}

	dst := b.Bytes()
	return dst, nil
}
