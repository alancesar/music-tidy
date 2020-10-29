package path

import (
	"bytes"
	"errors"
	"github.com/alancesar/tidy-music/metadata"
	"html/template"
	"path/filepath"
	"strings"
)

const defaultSeparator = "/"

var InvalidArgErr = errors.New("invalid argument")

func BuildPath(pattern, extension string, metadata metadata.Metadata) (string, error) {
	parsed, err := template.New("path").Parse(pattern)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := parsed.Execute(buf, metadata); err != nil {
		return "", InvalidArgErr
	}

	elements := strings.Split(buf.String()+extension, defaultSeparator)
	return filepath.Join(elements...), nil
}
