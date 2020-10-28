package dir

import (
	"fmt"
	"github.com/alancesar/music-tidy/metadata"
	"os"
	"strings"
)

func CreateArtistAndAlbumDirectory(root string, metadata metadata.Metadata) (string, error) {
	artist := removeSlash(metadata.Artist)
	album := removeSlash(metadata.Album)
	path := fmt.Sprintf("%s/%s/[%d] %s", root, artist, metadata.Year, album)
	sanitized := strings.TrimSpace(path)
	return path, os.MkdirAll(sanitized, os.ModePerm)
}

func BuildFilename(extension string, metadata metadata.Metadata) string {
	title := removeSlash(metadata.Title)
	filename := fmt.Sprintf("%02d - %s%s", metadata.Track, title, extension)
	sanitized := strings.TrimSpace(filename)
	return sanitized
}

func removeSlash(string string) string {
	return strings.ReplaceAll(string, "/", "-")
}
