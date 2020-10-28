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

func Process(path, root string, move bool) (string, error) {
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

	artistAndAlbumPath, err := dir.CreateArtistAndAlbumDirectory(root, m)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(path)
	filename := dir.BuildFilename(ext, m)
	destination := fmt.Sprintf("%s/%s", artistAndAlbumPath, filename)

	if move {
		err = file.Move(path, destination)
	} else {
		err = file.Copy(source, destination)
	}

	return destination, err
}
