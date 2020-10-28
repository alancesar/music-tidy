package processor

import (
	"github.com/alancesar/music-tidy/dir"
	"github.com/alancesar/music-tidy/file"
	"github.com/alancesar/music-tidy/metadata"
	"os"
)

func Process(path string, root string) (string, int64, error) {
	source, err := os.Open(path)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		_ = source.Close()
	}()

	m, err := metadata.ExtractMetadata(source)
	if err != nil {
		return "", 0, err
	}

	artistAndAlbumPath, err := dir.CreateArtistAndAlbumDirectory(root, m)
	destination := dir.BuildFullPath(artistAndAlbumPath, m)
	bytes, err := file.Copy(destination, source)
	return destination, bytes, err
}
