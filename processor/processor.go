package processor

import (
	"github.com/alancesar/tidy-music/command"
	"github.com/alancesar/tidy-music/metadata"
	"github.com/alancesar/tidy-music/path"
	"os"
	"path/filepath"
)

func Process(sourcePath, rootDestination, pattern string, commands ...command.Command) (string, error) {
	source, err := os.Open(sourcePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = source.Close()
	}()

	m, err := metadata.ExtractMetadata(source)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(sourcePath)
	outputPath, err := path.BuildPath(pattern, ext, m)
	if err != nil {
		return "", err
	}

	destinationPath := filepath.Join(rootDestination, outputPath)
	destinationPath = filepath.Clean(destinationPath)

	for _, cmd := range commands {
		if err := cmd(sourcePath, destinationPath); err != nil {
			return destinationPath, err
		}
	}

	return destinationPath, nil
}
