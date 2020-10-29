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

func BuildPath(pattern, extension string, metadata metadata.Metadata) (string, error) {
	parsed, err := template.New("path").Parse(pattern)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = parsed.Execute(buf, metadata)
	elements := strings.Split(buf.String()+extension, defaultSeparator)
	return filepath.Join(elements...), err
}
