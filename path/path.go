package path

import (
	"bytes"
	"errors"
	"github.com/alancesar/tidy-music/mime"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const defaultSeparator = "/"

var InvalidArgErr = errors.New("invalid argument")

func LookFor(rootPath string, t mime.Type) []string {
	paths := make([]string, 0)
	_ = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && info.Mode().IsRegular() && mime.Is(path, t) {
			paths = append(paths, path)
		}

		return nil
	})

	return paths
}

func BuildFromPattern(pattern string, source interface{}) (string, error) {
	parsed, err := template.New("path").Parse(pattern)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := parsed.Execute(buf, source); err != nil {
		return "", InvalidArgErr
	}

	elements := strings.Split(buf.String(), defaultSeparator)
	return filepath.Join(elements...), nil
}
