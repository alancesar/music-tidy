package dir

import (
	"fmt"
	"github.com/alancesar/music-tidy/metadata"
	"strings"
)

func BuildArtistAndAlbumPath(metadata metadata.Metadata) string {
	artist := sanitize(metadata.Artist)
	album := sanitize(metadata.Album)
	path := fmt.Sprintf("%s/[%d] %s", artist, metadata.Year, album)
	sanitized := strings.TrimSpace(path)
	return sanitized
}

func BuildFilename(metadata metadata.Metadata, extension string) string {
	title := sanitize(metadata.Title)
	filename := fmt.Sprintf("%02d - %s%s", metadata.Track, title, extension)
	sanitized := strings.TrimSpace(filename)
	return sanitized
}

func sanitize(string string) string {
	string = strings.ReplaceAll(string, "/", "-")
	return strings.ReplaceAll(string, ":", " -")
}
