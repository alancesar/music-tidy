package dir

import (
	"fmt"
	"github.com/alancesar/music-tidy/metadata"
	"os"
	"strings"
)

func CreateArtistAndAlbumDirectory(root string, metadata metadata.Metadata) (string, error) {
	artist := sanitize(metadata.Artist)
	album := sanitize(metadata.Album)
	path := fmt.Sprintf("%s/%s/[%d] %s", root, artist, metadata.Year, album)
	return path, os.MkdirAll(path, os.ModePerm)
}

func BuildFullPath(root string, metadata metadata.Metadata) string {
	title := sanitize(metadata.Title)
	return fmt.Sprintf("%s/%02d - %s", root, metadata.Track, title)
}

func sanitize(filename string) string {
	return strings.ReplaceAll(filename, "/", "-")
}
