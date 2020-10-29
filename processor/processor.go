package processor

import (
	"errors"
	"github.com/alancesar/tidy-music/metadata"
	"github.com/alancesar/tidy-music/path"
	"os"
	"path/filepath"
)

var MetadataErr = errors.New("no metadata found")

func Process(sourcePath, rootDestination, pattern string, sandbox bool) (string, error) {
	var err error

	source, err := os.Open(sourcePath)
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

	ext := filepath.Ext(sourcePath)
	outputPath, err := path.BuildPath(pattern, ext, m)
	if err != nil {
		return "", err
	}

	completePath := filepath.Join(rootDestination, outputPath)
	completePath = filepath.Clean(completePath)

	if !sandbox {
		dir, _ := filepath.Split(completePath)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}

		err = os.Rename(sourcePath, completePath)
	}

	return completePath, err
}
