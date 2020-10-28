package processor

import (
	"errors"
	"fmt"
	"github.com/alancesar/music-tidy/dir"
	"github.com/alancesar/music-tidy/file"
	"github.com/alancesar/music-tidy/metadata"
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

	artistAndAlbumPath := dir.BuildArtistAndAlbumPath(m)
	completePath := fmt.Sprintf("%s/%s", root, artistAndAlbumPath)
	completeCleanPath := filepath.Clean(completePath)
	if err := os.MkdirAll(completeCleanPath, os.ModePerm); err != nil {
		return "", err
	}

	ext := filepath.Ext(path)
	filename := dir.BuildFilename(m, ext)
	destination := fmt.Sprintf("%s/%s", completeCleanPath, filename)

	return destination, file.Move(path, destination)
}
