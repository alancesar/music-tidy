package path

import (
	"bytes"
	"github.com/alancesar/tidy-music/metadata"
	"html/template"
	"path/filepath"
	"strings"
)

var pathTemplate = template.New("path")

const defaultSeparator = "/"

func BuildPath(pattern, extension string, metadata metadata.Metadata) (string, error) {
	parsed, err := pathTemplate.Parse(pattern)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = parsed.Execute(buf, metadata)
	elements := strings.Split(buf.String()+extension, defaultSeparator)
	return filepath.Join(elements...), err
}
