package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"mime"
	"path/filepath"
)

func init() {
	mime.AddExtensionType(".ico", "image/x-icon")
	mime.AddExtensionType(".eot", "font/eot")
	mime.AddExtensionType(".tff", "font/tff")
	mime.AddExtensionType(".woff", "application/font-woff")
	mime.AddExtensionType(".woff2", "application/font-woff")
}

func Asset(base, path string) ([]byte, string, string, error) {
	var b bytes.Buffer
	file := base + path
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, "", "", err
	}
	if data != nil {
		w := gzip.NewWriter(&b)
		_, err = w.Write(data)
		err = w.Close()
		data = b.Bytes()
	}
	sum := md5.Sum(data)
	return data, hex.EncodeToString(sum[1:]), mime.TypeByExtension(filepath.Ext(file)), nil
}