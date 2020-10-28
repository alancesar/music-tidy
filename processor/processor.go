package processor

import (
	"errors"
	"fmt"
	"github.com/alancesar/tidy-music/dir"
	"github.com/alancesar/tidy-music/file"
	"github.com/alancesar/tidy-music/metadata"
	"os"
	"path/filepath"
)

var MetadataErr = errors.New("no metadata found")

func Process(path, root string) (string, error) {
	var err error

	source, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = source.Close()
	}()

	m, err := metadata.ExtractMetadata(source)
	if err != nil {
		return "", MetadataErr
	}

	completePath := dir.BuildPath(m)
	completePath = fmt.Sprintf("%s/%s", root, completePath)
	completePath = filepath.Clean(completePath)
	if err := os.MkdirAll(completePath, os.ModePerm); err != nil {
		return "", err
	}

	ext := filepath.Ext(path)
	filename := file.BuildFilename(m, ext)
	destination := fmt.Sprintf("%s/%s", completePath, filename)
	return destination, os.Rename(path, destination)
}
