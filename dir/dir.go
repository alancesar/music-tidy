package dir

import (
	"fmt"
	"github.com/alancesar/tidy-music/metadata"
	"github.com/alancesar/tidy-music/sanitize"
	"strings"
)

func BuildPath(metadata metadata.Metadata) string {
	artist := sanitize.Sanitize(metadata.Artist)
	album := sanitize.Sanitize(metadata.Album)
	path := fmt.Sprintf("%s/[%d] %s", artist, metadata.Year, album)
	sanitized := strings.TrimSpace(path)
	return sanitized
}
