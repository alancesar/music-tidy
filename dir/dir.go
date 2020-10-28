package dir

import (
	"fmt"
	"github.com/alancesar/music-tidy/metadata"
	"github.com/alancesar/music-tidy/sanitize"
	"strings"
)

func BuildPath(metadata metadata.Metadata) string {
	artist := sanitize.Sanitize(metadata.Artist)
	album := sanitize.Sanitize(metadata.Album)
	path := fmt.Sprintf("%s/[%d] %s", artist, metadata.Year, album)
	sanitized := strings.TrimSpace(path)
	return sanitized
}
