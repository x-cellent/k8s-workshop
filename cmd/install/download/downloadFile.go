package download

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var client = http.Client{
	CheckRedirect: func(r *http.Request, via []*http.Request) error {
		r.URL.Opaque = r.URL.Path
		return nil
	},
}

func Bytes(fileURL string) ([]byte, error) {
	resp, err := client.Get(fileURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func File(fileURL, destinationDir string) (string, error) {
	resp, err := client.Get(fileURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	u, err := url.Parse(fileURL)
	if err != nil {
		return "", err
	}

	path := filepath.Join(destinationDir, filepath.Base(u.Path))
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return path, nil
}
